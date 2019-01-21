// Copyright 2017, TCN Inc.
// All rights reserved.

// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:

//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of TCN Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package generator

import (
	"bytes"
	"fmt"

	// descriptor "github.com/golang/protobuf/descriptor"
	"strings"
	"text/template"

	gp "github.com/golang/protobuf/protoc-gen-go/descriptor"
	strcase "github.com/iancoleman/strcase"
	"github.com/tcncloud/protoc-gen-state/state"
)

/// note: adapted from a 400 line c++ file. sorry in advance

func createTypeMapAndImportSlice(servFiles []*gp.FileDescriptorProto, allFiles []*gp.FileDescriptorProto) (map[*gp.DescriptorProto]map[*gp.FieldDescriptorProto]*gp.DescriptorProto, map[*gp.DescriptorProto]*gp.FileDescriptorProto, []*gp.FileDescriptorProto, error) {
	typeMap := make(map[*gp.DescriptorProto]map[*gp.FieldDescriptorProto]*gp.DescriptorProto)
	descMap := make(map[*gp.DescriptorProto]*gp.FileDescriptorProto)
	fileSlice := []*gp.FileDescriptorProto{}

	for _, servFile := range servFiles {
		for _, service := range servFile.GetService() {
			for _, method := range service.GetMethod() {
				// get the input to the rpc
				inputType := method.GetInputType()
				// get the descriptor for that input type
				desc, file, _, err := FindDescriptor(allFiles, inputType)
				if err != nil {
					return typeMap, descMap, fileSlice, err
				}
				// add to the descriptor : file map
				descMap[desc] = file
				// add to fileSlice
				if !containsFile(fileSlice, file) {
					fileSlice = append(fileSlice, file)
				}
				// set the nested time in the big map
				typeMap, descMap, fileSlice, err = setNestedTypes(desc, servFile, allFiles, typeMap, descMap, fileSlice)
				if err != nil {
					return typeMap, descMap, fileSlice, err
				}
			}
		}
	}
	return typeMap, descMap, fileSlice, nil
}

// creates the map of descriptor: map[field]descriptor
func setNestedTypes(message *gp.DescriptorProto, msgFile *gp.FileDescriptorProto, files []*gp.FileDescriptorProto, typeMap map[*gp.DescriptorProto]map[*gp.FieldDescriptorProto]*gp.DescriptorProto, descMap map[*gp.DescriptorProto]*gp.FileDescriptorProto, fileSlice []*gp.FileDescriptorProto) (map[*gp.DescriptorProto]map[*gp.FieldDescriptorProto]*gp.DescriptorProto, map[*gp.DescriptorProto]*gp.FileDescriptorProto, []*gp.FileDescriptorProto, error) {

	// check if it's already done
	if _, ok := typeMap[message]; ok {
		return typeMap, descMap, fileSlice, nil
	}

	if !containsFile(fileSlice, msgFile) {
		fileSlice = append(fileSlice, msgFile)
	}

	map_exists := false
	for _, field := range message.GetField() {
		field_type := field.GetType()

		if field_type == gp.FieldDescriptorProto_TYPE_GROUP || field_type == gp.FieldDescriptorProto_TYPE_MESSAGE {
			if !map_exists {
				field_to_type := make(map[*gp.FieldDescriptorProto]*gp.DescriptorProto, 0)
				typeMap[message] = field_to_type
				map_exists = true
			}

			// get descriptor from the type name
			desc, file, _, err := FindDescriptor(files, field.GetTypeName())
			if err != nil {
				return nil, nil, nil, err
			}
			typeMap[message][field] = desc
			descMap[desc] = file
			typeMap, descMap, fileSlice, err := setNestedTypes(desc, file, files, typeMap, descMap, fileSlice)
			if err != nil {
				return typeMap, descMap, fileSlice, err
			}
		}
	}

	return typeMap, descMap, fileSlice, nil
}

type ImportEntity struct {
	ModuleName string
	FileName   string
}

func generateImportEntities(fileSlice []*gp.FileDescriptorProto, protocTsPath string) []*ImportEntity {
	importEntities := []*ImportEntity{}
	for _, f := range fileSlice {
		filepath := GetFilePath(f.GetName())
		index := strings.LastIndex(filepath, "/") + 1
		fname := strings.Replace(f.GetName(), "/", "_", -1)
		packageSlashes := strings.Replace(f.GetPackage(), ".", "/", -1)
		if f.GetName()[:15] == "google/protobuf" {
			importEntities = append(importEntities, &ImportEntity{
				FileName:   fmt.Sprintf("google-protobuf/%s_pb", f.GetName()[:len(f.GetName())-6]),
				ModuleName: fname[:len(fname)-6],
			})
		} else {
			importEntities = append(importEntities, &ImportEntity{
				FileName: protocTsPath + packageSlashes + "/" + filepath[index:],
				// FileName:   protocTsPath + "/" + filepath[index:],
				ModuleName: fname[:len(fname)-6],
			})
		}
	}
	return importEntities
}

type MappingEntity struct {
	MapName   string
	FileName  string
	TypeName  string
	TypeLines []*TypeLine
	Debug     bool
}

type TypeLine struct {
	Name     string
	FileName string
	TypeName string
}

func generateMappingEntities(typeMap map[*gp.DescriptorProto]map[*gp.FieldDescriptorProto]*gp.DescriptorProto, descMap map[*gp.DescriptorProto]*gp.FileDescriptorProto, improvedDescriptors []*ImprovedMessageDescriptor, protos []*gp.FileDescriptorProto, fileSlice []*gp.FileDescriptorProto, debug bool) ([]*MappingEntity, []*gp.FileDescriptorProto, error) {

	isComplex := func(fields []*gp.FieldDescriptorProto) bool {
		for _, field := range fields {
			if field.GetType() == gp.FieldDescriptorProto_TYPE_GROUP || field.GetType() == gp.FieldDescriptorProto_TYPE_MESSAGE {
				return true
			}
		}
		return false
	}

	mappingEntities := []*MappingEntity{}
	// TODO handle case with empty map in template
	for desc, sub := range typeMap {
		improvedDescriptor := FindImprovedFromDescriptor(improvedDescriptors, desc)
		improvedPathName := FindImprovedPathName(improvedDescriptor)
		improvedPathUnderscore := strings.Replace(improvedPathName, ".", "_", -1)
		packageNameUnderscore := strings.Replace(improvedDescriptor.packageName, ".", "_", -1)
		fullNameUnderscore := fmt.Sprintf("%s_%s%s", packageNameUnderscore, improvedPathUnderscore, desc.GetName())

		// disgusting
		if desc.GetOptions().GetMapEntry() {
			continue
			// fullNameUnderscore = fullNameUnderscore[:len(fullNameUnderscore)-5]
		}

		mapName := fullNameUnderscore + "_map"

		typeLines := []*TypeLine{}
		for field, _ := range sub {
			repeated := field.GetLabel() == 3
			// get file that defines the message type

			field_type := field.GetTypeName()
			foundDesc, fileDesc, _, err := FindDescriptor(protos, field_type)
			if err != nil {
				return nil, nil, err
			}

			// add to import slice
			if !containsFile(fileSlice, fileDesc) {
				fileSlice = append(fileSlice, fileDesc)
			}
			filename := strings.Replace(fileDesc.GetName(), "/", "_", -1)
			filename = filename[:len(filename)-6] // remove .proto

			typename := strings.TrimPrefix(field_type, "."+improvedDescriptor.packageName)
			if strings.HasPrefix(typename, ".google.protobuf") {
				typename = typename[16:] // hacky
			} else if strings.HasPrefix(typename, ".commons") {
				typename = typename[8:] // even hackier
			} else if strings.HasPrefix(typename, ".matrix.lms") {
				typename = typename[11:]
			} else if strings.HasPrefix(typename, ".matrix.eps") {
				typename = typename[11:]
			}

			var name string
			if foundDesc.GetOptions().GetMapEntry() && isComplex(foundDesc.GetField()) {
				name = strcase.ToLowerCamel(field.GetName()) + "Map"
				// begin the hacky sack
				if strings.Contains(typename, "Map.") && strings.HasSuffix(typename, "Entry") {
					ind := strings.LastIndex(typename, "Map.")
					typename = typename[:ind]
				} else if strings.HasSuffix(typename, "Entry") {
					ind := strings.LastIndex(typename, ".")
					typename = typename[:ind]
				}
			} else if foundDesc.GetOptions().GetMapEntry() {
				continue // basic map, basic types
			} else if repeated {
				name = strcase.ToLowerCamel(field.GetName()) + "List"
			} else {
				name = strcase.ToLowerCamel(field.GetName())
			}

			// number of dots in package
			typeLines = append(typeLines, &TypeLine{
				Name:     name,
				FileName: filename,
				TypeName: typename,
				// TypeName: field_type[strings.LastIndex(field_type, ".")+1:],
			})
		}

		file := improvedDescriptor.file
		fileName := strings.Replace(file.GetName(), "/", "_", -1)
		fileName = fileName[:len(fileName)-6] // remove .proto

		tName := improvedPathName + desc.GetName() // nice hax
		if desc.GetOptions().GetMapEntry() {
			tName = tName[:len(tName)-5]
		}
		mappingEntities = append(mappingEntities, &MappingEntity{
			MapName:   mapName,
			FileName:  fileName,
			TypeName:  tName,
			TypeLines: typeLines,
			Debug:     debug,
		})
	}
	return mappingEntities, fileSlice, nil
}

func (this *GenericOutputter) CreateToMessageFile(servFiles []*gp.FileDescriptorProto, outputType state.OutputTypes, protos []*gp.FileDescriptorProto, protocTsPath string, debug bool) (*File, error) {
	improvedDescriptors := CreateImprovedDescriptors(protos)

	// TODO: first part finished, should have a type map set up now
	typeMap, descMap, fileSlice, err := createTypeMapAndImportSlice(servFiles, protos)
	if err != nil {
		return nil, fmt.Errorf("Error when building typeMap and fileSlice: %v\n%s\n%s", err, typeMap, fileSlice) // TODO debugging
	}

	// use the typeMap to create entities for the main template
	mappingEntities, fileSlice, err := generateMappingEntities(typeMap, descMap, improvedDescriptors, protos, fileSlice, debug) /*[]*MappingEntity{}*/
	if err != nil {
		return nil, err
	}

	// map fileSlice to import entities
	importEntities := generateImportEntities(fileSlice, protocTsPath) /*[]*ImportEntity{}*/

	var output bytes.Buffer

	// generate import statements
	imports := template.Must(template.New("imports").Parse(this.ToMessageFile.ImportTemplate))
	imports.Execute(&output, importEntities)

	body := template.Must(template.New("body").Parse(this.ToMessageFile.MappingTemplate))
	body.Execute(&output, mappingEntities)

	toMessageOutput := template.Must(template.New("toMessageOutput").Parse(this.ToMessageFile.Template))
	toMessageOutput.Execute(&output, debug)

	return &File{
		Name:    "to_message_pb.ts",
		Content: output.String(),
	}, nil
}
