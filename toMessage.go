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

const mappingTemplate = `{{range $i, $e := .}}
const messageMap = new Map();
{{end}}`

type ImportEntity struct {
	ModuleName string
	FileName   string
}

type MappingEntity struct {
}

func CreateToMessageFile(servFiles []*gp.FileDescriptorProto, protos []*gp.FileDescriptorProto, protocTsPath string) (*File, error) {

	// TODO: first part finished, should have a type map set up now
	_, fileSlice, err := createTypeMapAndImportSlice(servFiles, protos)
	if err != nil {
		return nil, fmt.Errorf("Error when building typeMap and fileSlice: %v", err)
	}

	// map fileSlice to import entities
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

	// generate the mappings
	// mappingEntities, err := generateTypeMapToString(typeMap)

	var output bytes.Buffer
	// generate import statements
	imports := template.Must(template.New("imports").Parse(importTemplate))
	imports.Execute(&output, importEntities)

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
