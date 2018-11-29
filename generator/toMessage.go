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
	gp "github.com/golang/protobuf/protoc-gen-go/descriptor"
	strcase "github.com/iancoleman/strcase"
	"strings"
	"text/template"
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
				desc, file, err := FindDescriptor(allFiles, inputType)
				if err != nil {
					return typeMap, descMap, fileSlice, err
				}
				// add to the descriptor : file map
				descMap[desc] = file
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
			desc, file, err := FindDescriptor(files, field.GetTypeName())
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

const importTemplate = `{{range $i, $e := .}}
import {{$e.ModuleName}} from '{{$e.FileName}}';{{end}}`

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
		if f.GetName()[:15] == "google/protobuf" {
			importEntities = append(importEntities, &ImportEntity{
				FileName:   fmt.Sprintf("google-protobuf/%s_pb", f.GetName()[:len(f.GetName())-6]),
				ModuleName: fname[:len(fname)-6],
			})
		} else {
			importEntities = append(importEntities, &ImportEntity{
				FileName:   protocTsPath + filepath[index:],
				ModuleName: fname[:len(fname)-6],
			})
		}
	}
	return importEntities
}

const mappingTemplate = `
const messageMap = new Map();
{{range $i, $e := .}}
const {{$e.MapName}} = new Map();{{range $x, $t := $e.TypeLines }}
{{$e.MapName}}.set('{{$t.Name}}', {{$t.FileName}}{{$t.TypeName}});{{end}}
messageMap.set({{$e.FileName}}.{{$e.TypeName}}, {{$e.MapName}});
{{end}}`

type MappingEntity struct {
	MapName   string
	FileName  string
	TypeName  string
	TypeLines []*TypeLine
}

type TypeLine struct {
	Name     string
	FileName string
	TypeName string
}

func generateMappingEntities(typeMap map[*gp.DescriptorProto]map[*gp.FieldDescriptorProto]*gp.DescriptorProto, descMap map[*gp.DescriptorProto]*gp.FileDescriptorProto, protos []*gp.FileDescriptorProto) ([]*MappingEntity, error) {

	isMap := func(field *gp.FieldDescriptorProto, desc *gp.DescriptorProto) bool {
		opts := desc.GetOptions()
		if field.GetType() == gp.FieldDescriptorProto_TYPE_MESSAGE && opts != nil && *opts.MapEntry {
			return true
		}
		return false
	}

	mappingEntities := []*MappingEntity{}
	// TODO handle case with empty map in template
	for desc, sub := range typeMap {
		packageNameUnderscore := strings.Replace(descMap[desc].GetPackage(), ".", "_", -1)
		fullNameUnderscore := fmt.Sprintf("%s_%s", packageNameUnderscore, desc.GetName())
		mapName := fullNameUnderscore + "_map"

		typeLines := []*TypeLine{}
		for field, descriptor := range sub {
			repeated := field.GetLabel() == 3
			// get file that defines the message type

			field_type := field.GetTypeName()
			_, fileDesc, err := FindDescriptor(protos, field_type)
			if err != nil {
				return nil, err
			}
			filename := strings.Replace(fileDesc.GetName(), "/", "_", -1)
			filename = filename[:len(filename)-6] // remove .proto

			var name string
			if isMap(field, descriptor) {
				name = strcase.ToLowerCamel(field.GetName()) + "Map"
			} else if repeated {
				name = strcase.ToLowerCamel(field.GetName()) + "List"
			} else {
				name = strcase.ToLowerCamel(field.GetName())
			}

			typename := strings.TrimPrefix(field_type, "."+descMap[desc].GetPackage())
			if strings.HasPrefix(typename, ".google.protobuf") {
				typename = typename[16:] // hacky
			} else if strings.HasPrefix(typename, ".commons") {
				typename = typename[8:]
			}

			// number of dots in package
			typeLines = append(typeLines, &TypeLine{
				Name:     name,
				FileName: filename,
				TypeName: typename,
				// TypeName: field_type[strings.LastIndex(field_type, ".")+1:],
			})
		}
		file := descMap[desc]
		fileName := strings.Replace(file.GetName(), "/", "_", -1)
		fileName = fileName[:len(fileName)-6] // remove .proto

		mappingEntities = append(mappingEntities, &MappingEntity{
			MapName:   mapName,
			FileName:  fileName,
			TypeName:  desc.GetName(),
			TypeLines: typeLines,
		})
	}
	return mappingEntities, nil
}

func CreateToMessageFile(servFiles []*gp.FileDescriptorProto, protos []*gp.FileDescriptorProto, protocTsPath string) (*File, error) {

	// TODO: first part finished, should have a type map set up now
	typeMap, descMap, fileSlice, err := createTypeMapAndImportSlice(servFiles, protos)
	if err != nil {
		return nil, fmt.Errorf("Error when building typeMap and fileSlice: %v\n%s\n%s", err, typeMap, fileSlice) // TODO debugging
	}

	// map fileSlice to import entities
	importEntities := generateImportEntities(fileSlice, protocTsPath) /*[]*ImportEntity{}*/

	// use the typeMap to create entities for the main template
	mappingEntities, err := generateMappingEntities(typeMap, descMap, protos) /*[]*MappingEntity{}*/
	if err != nil {
		return nil, err
	}

	var output bytes.Buffer

	// generate import statements
	imports := template.Must(template.New("imports").Parse(importTemplate))
	imports.Execute(&output, importEntities)

	body := template.Must(template.New("body").Parse(mappingTemplate))
	body.Execute(&output, mappingEntities)

	return &File{
		Name:    "to_message_pb.ts",
		Content: output.String() + toMessageFunction,
	}, nil
}

const toMessageFunction = `

function getNestedMessageConstructor(messageType, fieldName) {
	return messageMap.has(messageType) && messageMap.get(messageType).get(fieldName);
}

export function toMessage(obj: any, messageClass: any) {
	if (!obj) {
		return new messageClass();
	}

	const message = new messageClass();

	Object.keys(obj).forEach(key => {
		let ele = obj[key];
		const upperCaseKey = key.charAt(0).toUpperCase() + key.substr(1);
		const setterName = "set" + upperCaseKey;
		const getterName = "get" + upperCaseKey;

		if (message[setterName]) {
			var nestedMessageContructor = getNestedMessageConstructor(messageClass, key);
			if (nestedMessageContructor) {
				if (key.length > 4 && key.slice(key.length - 4) === 'List' && Array.isArray(ele)) { // check if field is repeated
					ele = ele.map(subEle => toMessage(subEle, nestedMessageContructor));
				} else {
					ele = toMessage(ele, nestedMessageContructor);
				}
			}

			message[setterName](ele);
		} else if (message[getterName] && key.slice(key.length - 3) === 'Map') { // check if field is a map
			// if the map field is missing, nothing needs to be done.
			if (ele !== undefined && ele !== null) {
				if (Array.isArray(ele)) {
					if (ele.length) {
						var mapObj = message[getterName]();
						var mappedFieldValueConstructor = getNestedMessageConstructor(messageClass, key);
						if (mappedFieldValueConstructor) {
							ele = ele.map(([key, value]) => {
								return [key, mappedFieldValueConstructor(value)];
							});
						}
						ele.forEach(([key, value]) => {
							mapObj.set(key, value);
						});
					}
				} else {
					throw new Error("Protoc-gen-state: Expected field " + key + " to be an array of tuples.");
				}
			}
		} else {
			throw new Error("No corresponding gRPC setter method for given key: " + key);
		}
	});

	return message;
}`
