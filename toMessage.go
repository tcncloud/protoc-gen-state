package main

import (
	"bytes"
	"strings"
	"text/template"

	"fmt"
	gp "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

/// note: adapted from a 400 line c++ file. sorry in advance

const toMessageTemplate = `/* THIS FILE IS GENERATED FROM THE TOOL PROTOC-GEN-STATE  */
/* ANYTHING YOU EDIT WILL BE OVERWRITTEN IN FUTURE BUILDS */

import * as ProtocTypes from './protoc_types_pb';

`

type MessageEntity struct {
	FieldName    string
	FullTypeName string
	Repeated     bool
}

func createTypeMapAndImportSlice(servFiles []*gp.FileDescriptorProto, allFiles []*gp.FileDescriptorProto) (map[*gp.DescriptorProto]map[*gp.FieldDescriptorProto]*gp.DescriptorProto, []*gp.FileDescriptorProto, error) {
	typeMap := make(map[*gp.DescriptorProto]map[*gp.FieldDescriptorProto]*gp.DescriptorProto, 0)
	fileSlice := []*gp.FileDescriptorProto{}

	for _, servFile := range servFiles {
		for _, service := range servFile.GetService() {
			for _, method := range service.GetMethod() {
				// get the input to the rpc
				inputType := method.GetInputType()
				// get the descriptor for that input type
				desc, _, err := FindDescriptor(allFiles, inputType)
				if err != nil {
					return typeMap, fileSlice, err
				}
				typeMap, fileSlice, err = setNestedTypes(desc, servFile, allFiles, typeMap, fileSlice)
				if err != nil {
					return typeMap, fileSlice, err
				}
			}
		}
	}
	return typeMap, fileSlice, nil
}

func setNestedTypes(message *gp.DescriptorProto, msgFile *gp.FileDescriptorProto, files []*gp.FileDescriptorProto, typeMap map[*gp.DescriptorProto]map[*gp.FieldDescriptorProto]*gp.DescriptorProto, fileSlice []*gp.FileDescriptorProto) (map[*gp.DescriptorProto]map[*gp.FieldDescriptorProto]*gp.DescriptorProto, []*gp.FileDescriptorProto, error) {
	// check if it's already done
	if _, ok := typeMap[message]; ok {
		return typeMap, fileSlice, nil
	}

	if !containsFile(fileSlice, msgFile) {
		fileSlice = append(fileSlice, msgFile)
	}

	map_exists := false
	for _, field := range message.GetField() {
		if field.GetType() == gp.FieldDescriptorProto_TYPE_GROUP || field.GetType() == gp.FieldDescriptorProto_TYPE_MESSAGE {
			if !map_exists { // instantiate the field
				field_to_type := make(map[*gp.FieldDescriptorProto]*gp.DescriptorProto, 0)
				typeMap[message] = field_to_type
				map_exists = true
			}

			field_type := field.GetTypeName()
			desc, file, err := FindDescriptor(files, field_type)
			if err != nil {
				return typeMap, fileSlice, err
			}

			typeMap[message][field] = desc
			typeMap, fileSlice, err = setNestedTypes(desc, file, files, typeMap, fileSlice)
			if err != nil {
				return typeMap, fileSlice, err
			}
		}
	}
	return typeMap, fileSlice, nil
}

const importTemplate = `{{range $i, $e := .}}
import {{$e.ModuleName}} from '{{$e.FileName}}_pb';{{end}}`

type ImportEntity struct {
	ModuleName string
	FileName   string
}

func generateImportEntities(fileSlice []*gp.FileDescriptorProto, protocTsPath string) []*ImportEntity {
	importEntities := []*ImportEntity{}
	for _, f := range fileSlice {
		fname := strings.Replace(f.GetName(), "/", "_", -1)
		if f.GetName()[:15] == "google/protobuf" {
			importEntities = append(importEntities, &ImportEntity{
				FileName:   fmt.Sprintf("google-protobuf/%s", f.GetName()[:len(f.GetName())-6]),
				ModuleName: fname[:len(fname)-6],
			})
		} else {
			importEntities = append(importEntities, &ImportEntity{
				FileName:   protocTsPath + GetFilePath(f.GetName()),
				ModuleName: fname[:len(fname)-6],
			})
		}
	}
	return importEntities
}

const mappingTemplate = `const messageMap = new Map();
const mapFieldMap = new Map();
{{range $i, $e := .}}
const {{$e.MapName}} = new Map();
{{range _, $t := $e.TypeLines }}
{{$e.MapName}}.set('{{$t.Name}}', {{$t.FileName}}.{{$t.TypeName}});
{{end}}
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

func generateMappingEntities(typeMap map[*gp.DescriptorProto]map[*gp.FieldDescriptorProto]*gp.DescriptorProto, protos []*gp.FileDescriptorProto) ([]*MappingEntity, error) {

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
		fullNameUnderscore := strings.Replace(desc.GetName(), ".", "_", -1)
		mapName := fullNameUnderscore + "_map"

		typeLines := []*TypeLine{}
		for field, descriptor := range sub {
			repeated := field.GetLabel() == 3
			// get file that defines the message type
			_, fileDesc, err := FindDescriptor(protos, field.GetTypeName())
			if err != nil {
				return nil, err
			}
			filename := strings.Replace(fileDesc.GetName(), "/", "_", -1)
			filename = filename[:len(filename)-6] // remove .proto

			var name string
			if isMap(field, descriptor) {
				// TODO
				// mapFieldValueType := descriptor.GetField()[1].GetType() // 1?
				// if mapFieldValueType == FieldDescriptorProto_TYPE_MESSAGE { } // ?
				name = field.GetName() + "Map"
			} else if repeated {
				name = field.GetName() + "List"
			} else {
				name = field.GetName()
			}

			typeLines = append(typeLines, &TypeLine{
				Name:     name,
				FileName: filename,
				TypeName: "tmp",
			})
		}
		// get filename
		_, file, err := FindDescriptor(protos, desc.GetName())
		if err != nil {
			return nil, err
		}
		fileName := file.GetName()
		fileName = fileName[:len(fileName)-6] // remove .proto

		mappingEntities = append(mappingEntities, &MappingEntity{
			MapName:   mapName,
			FileName:  fileName,
			TypeName:  "tmp",
			TypeLines: typeLines,
		})
	}
	return mappingEntities, nil
}

func CreateToMessageFile(servFiles []*gp.FileDescriptorProto, protos []*gp.FileDescriptorProto, protocTsPath string) (*File, error) {

	// TODO: first part finished, should have a type map set up now
	typeMap, fileSlice, err := createTypeMapAndImportSlice(servFiles, protos)
	if err != nil {
		return nil, fmt.Errorf("Error when building typeMap and fileSlice: %v", err)
	}

	// map fileSlice to import entities
	importEntities := generateImportEntities(fileSlice, protocTsPath) /*[]*ImportEntity{}*/

	// use the typeMap to create entities for the main template
	mappingEntities, err := generateMappingEntities(typeMap, protos) /*[]*MappingEntity{}*/
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
			throw new Error("No corresponding gRPC setter method for given key: " + key + ");
		}
	});

	return message;
}`
