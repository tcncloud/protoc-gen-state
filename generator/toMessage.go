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
)

/// note: adapted from a 400 line c++ file. sorry in advance

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
		packageSlashes := strings.Replace(f.GetPackage(), ".", "/", -1)
		if f.GetName()[:15] == "google/protobuf" {
			importEntities = append(importEntities, &ImportEntity{
				FileName:   fmt.Sprintf("google-protobuf/%s_pb", f.GetName()[:len(f.GetName())-6]),
				ModuleName: fname[:len(fname)-6],
			})
		} else {
			importEntities = append(importEntities, &ImportEntity{
				FileName:   protocTsPath + packageSlashes + "/" + filepath[index:],
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
{{$e.MapName}}.set('{{$t.NestedName}}', {{$t.NestedFileName}}{{$t.NestedTypeName}});{{end}}
messageMap.set({{$e.FileName}}.{{$e.TypeName}}, {{$e.MapName}});
{{end}}`

type MappingEntity struct {
	MapName   string
	FileName  string
	TypeName  string
	TypeLines []*TypeLine
	Debug     bool
}

type TypeLine struct {
	NestedName     string
	NestedFileName string
	NestedTypeName string
}

func generateMappingEntities(protos []*gp.FileDescriptorProto, improvedDescriptors []*ImprovedMessageDescriptor, debug bool) ([]*MappingEntity, []*gp.FileDescriptorProto, error) {

	isComplex := func(fields []*gp.FieldDescriptorProto) bool {
		for _, field := range fields {
			if field.GetType() == gp.FieldDescriptorProto_TYPE_GROUP || field.GetType() == gp.FieldDescriptorProto_TYPE_MESSAGE {
				return true
			}
		}
		return false
	}

	mappingEntities := []*MappingEntity{}
	importFiles := []*gp.FileDescriptorProto{}

	for _, desc := range improvedDescriptors {
		improvedPathName := FindImprovedPathName(desc)
		improvedPathUnderscore := strings.Replace(improvedPathName, ".", "_", -1)
		packageNameUnderscore := strings.Replace(desc.packageName, ".", "_", -1)
		fullNameUnderscore := fmt.Sprintf("%s_%s%s", packageNameUnderscore, improvedPathUnderscore, desc.message.GetName())
		mapName := fullNameUnderscore + "_map"

		// skip map entries I guess
		if desc.message.GetOptions().GetMapEntry() {
			continue
		}

		typeLines := []*TypeLine{}
		for field, _ := range desc.message {
			repeated := field.GetLabel() == 3
			fileDesc := desc.file
			// get file that defines the message type

			field_type := field.GetTypeName()
			foundDesc, _, _, err := FindDescriptor(protos, field_type)
			if err != nil {
				return nil, nil, err
			}

			// add to import slice
			if !containsFile(importFiles, fileDesc) {
				importFiles = append(importFiles, fileDesc)
			}
			filename := strings.Replace(fileDesc.GetName(), "/", "_", -1)
			filename = filename[:len(filename)-6] // remove .proto

			typename := strings.TrimPrefix(field_type, "."+desc.packageName)
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
				// golang hacks to ignore the suffix Map or Entry
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

			typeLines = append(typeLines, &TypeLine{
				NestedName:     name,
				NestedFileName: filename,
				NestedTypeName: typename,
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
	return mappingEntities, importFiles, nil
}

const toMessageTemplate = `

function getNestedMessageConstructor(messageType, fieldName) {
  return messageMap.has(messageType) && messageMap.get(messageType).get(fieldName);
}

export function toMessage(obj: any, messageClass: any) {
  if (!obj) {
    return new messageClass();
  }
  {{if .}}console.groupCollapsed('toMessage');{{end}}
  const message = new messageClass();

  Object.keys(obj).forEach(key => {
    {{if .}}console.groupCollapsed('field:', key);{{end}}
    let ele = obj[key];
    const upperCaseKey = key.charAt(0).toUpperCase() + key.substr(1);
    const setterName = "set" + upperCaseKey;
    const getterName = "get" + upperCaseKey;
    {{if .}}console.log('ele:', ele);
    console.log('typeof ele:', typeof ele);
    console.log('setterName:', setterName);
    console.log('getterName:', getterName);{{end}}

    if (message[setterName]) {
      var nestedMessageContructor = getNestedMessageConstructor(messageClass, key);
      if (nestedMessageContructor) {
        if (key.length > 4 && key.slice(key.length - 4) === 'List' && Array.isArray(ele)) { // check if field is repeated
          {{if .}}console.log('REPEATED field');{{end}}
          ele = ele.map(subEle => toMessage(subEle, nestedMessageContructor));
        } else {
          {{if .}}console.log('regular field');{{end}}
          ele = toMessage(ele, nestedMessageContructor);
        }
      }

      message[setterName](ele);
    } else if (message[getterName] && key.slice(key.length - 3) === 'Map') { // check if field is a map
      {{if .}}console.log('MAP field');{{end}}
      // if the map field is missing, nothing needs to be done.
      if (ele !== undefined && ele !== null) {
        if (Array.isArray(ele)) {
          if (ele.length) {
            var mapObj = message[getterName]();
            var mappedFieldValueConstructor = getNestedMessageConstructor(messageClass, key);
            if (mappedFieldValueConstructor) {
              {{if .}}console.groupCollapsed('keys & values, unserialized');{{end}}
              ele = ele.map(([key, value]) => {
                {{if .}}console.log('key:', key);
                console.log('value:', value);{{end}}
                return [key, mappedFieldValueConstructor(value)];
              });
              {{if .}}console.groupEnd();{{end}}
            }
            {{if .}}console.groupCollapsed('keys & values, serialized');{{end}}
            ele.forEach(([key, value]) => {
              {{if .}}console.log('key:', key);
              console.log('value:', value);{{end}}
              mapObj.set(key, value);
            });
            {{if .}}console.groupEnd();{{end}}
          }
        } else {
          {{if .}}console.groupEnd();
          console.groupEnd();{{end}}
          throw new Error("Protoc-gen-state: Expected field " + key + " to be an array of tuples.");
        }
      }
    } else {
      {{if .}}console.groupEnd();
      console.groupEnd();{{end}}
      throw new Error("No corresponding gRPC setter method for given key: " + key);
    }
    {{if .}}console.groupEnd();{{end}}
  });
  {{if .}}console.groupEnd();{{end}}

  return message;
}
`

func CreateToMessageFile(servFiles []*gp.FileDescriptorProto, protos []*gp.FileDescriptorProto, protocTsPath string, debug bool) (*File, error) {
	improvedDescriptors := CreateImprovedDescriptors(protos)

	// use the improvedDescriptors
	// use the typeMap to create entities for the main template
	mappingEntities, importFiles, err := generateMappingEntities(protos, improvedDescriptors, debug)
	if err != nil {
		return nil, err
	}

	// map importFiles slice to import entities
	importEntities := generateImportEntities(importFiles, protocTsPath)

	var output bytes.Buffer

	// generate import statements for javascript
	imports := template.Must(template.New("imports").Parse(importTemplate))
	imports.Execute(&output, importEntities)

	// generate map structures for javascript
	body := template.Must(template.New("body").Parse(mappingTemplate))
	body.Execute(&output, mappingEntities)

	// javascript helper functions
	toMessageOutput := template.Must(template.New("toMessageOutput").Parse(toMessageTemplate))
	toMessageOutput.Execute(&output, debug)

	return &File{
		Name:    "to_message_pb.ts",
		Content: output.String(),
	}, nil
}
