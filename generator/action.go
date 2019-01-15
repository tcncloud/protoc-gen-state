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
	gp "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/tcncloud/protoc-gen-state/state"
	"strings"
	"text/template"
)


type ActionEntity struct {
	JsonName   string
	InputType  string
	OutputType string
	Repeat     bool
}

func (this *GenericOutputter) CreateActionFile(stateFields []*gp.FieldDescriptorProto, outputType state.OutputTypes, customFields []*gp.FieldDescriptorProto, serviceFiles []*gp.FileDescriptorProto, debug bool) (*File, error) {
	getEntities := []*ActionEntity{}
	listEntities := []*ActionEntity{}
	resetEntities := []*ActionEntity{}
	createEntities := []*ActionEntity{}
	updateEntities := []*ActionEntity{}
	deleteEntities := []*ActionEntity{}
	customEntities := []*ActionEntity{}

	// populate Entities slices with state fields
	for _, field := range stateFields {
		repeated := field.GetLabel() == 3

		// verify the method annotations
    fieldAnnotations, err := GetFieldOptions(field)
		if err != nil {
			return nil, fmt.Errorf("Error getting field level annotations: %v", err)
		}

		// using custom method annotation anywhere but CustomActions is an error
		if fieldAnnotations.GetMethod().GetCustom() != "" {
			return nil, fmt.Errorf("Method annotation 'custom' provided outside of CustomActions message. Correct this by removing the '(method).custom' annotation on field: %s", field.GetName())
		}

		// loop through CLUDG verbs
		var meth *gp.MethodDescriptorProto
		for c := CREATE; c < CRUD_MAX; c++ {
			// verify the annotation exists in the fieldAnnotations struct
			crudAnnotation := GetAnnotation(*fieldAnnotations.GetMethod(), c, repeated)
			if crudAnnotation != "" {
				meth, err = FindMethodDescriptor(serviceFiles, crudAnnotation)
				if err != nil {
					return nil, err
				}
			}

			// if method is still empty, the method could not be found so do nothing
			if meth != nil {
				// found it so add it to the action entity for this crud value (c)
				action := &ActionEntity{
					JsonName:   *field.JsonName,
					InputType:  fmt.Sprintf("ProtocTypes.%s.AsObject", CreatePackageAndTypeString(meth.GetInputType())),
					OutputType: fmt.Sprintf("ProtocTypes.%s.AsObject", CreatePackageAndTypeString(meth.GetOutputType())),
					Repeat:     repeated,
				}
				switch c {
				case CREATE:
					createEntities = append(createEntities, action)
				case DELETE:
					deleteEntities = append(deleteEntities, action)
				case UPDATE:
					updateEntities = append(updateEntities, action)
				case GET:
					if repeated {
						listEntities = append(listEntities, action)
					} else {
						getEntities = append(getEntities, action)
					}
				default:
					// nothing
				}
			}
		}

		// create a reset action for each field name
		resetEntities = append(resetEntities, &ActionEntity{
			JsonName: *field.JsonName,
		})
	}

	// do the same things for custom actions
	// TODO combine the logic since its basically the same
	for _, field := range customFields {
		repeated := field.GetLabel() == 3
		// get the method annoations
    fieldAnnotations, err := GetFieldOptions(field)
		if err != nil {
			return nil, fmt.Errorf("Error getting field level annotations: %v", err)
		}

		// only custom method annotations allowed in CustomActions
		for c := CREATE; c < CRUD_MAX; c++ {
			if GetAnnotation(*fieldAnnotations.GetMethod(), c, repeated) != "" {
				return nil, fmt.Errorf("Invalid method annotation. Only method annotation '(method).custom' is allowed within CustomActions message. Correct the annotation on field: %s", *field.JsonName)
			}
		}

		// and it must have a custom annotation
		var meth *gp.MethodDescriptorProto
		crudAnnotation := fieldAnnotations.GetMethod().GetCustom()
		if crudAnnotation != "" {
			meth, err = FindMethodDescriptor(serviceFiles, crudAnnotation)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("CustomAction field provided without an accompanying '(method).custom' annotation. Please provide one on field: %s", *field.JsonName)
		}

		// hopefully the method exists
		if meth != nil {
			// TODO this uses repeated from the field value but should use repeated from the output type
			customEntities = append(customEntities, &ActionEntity{
				JsonName:   *field.JsonName,
				InputType:  fmt.Sprintf("ProtocTypes.%s.AsObject", CreatePackageAndTypeString(meth.GetInputType())),
				OutputType: fmt.Sprintf("ProtocTypes.%s.AsObject", CreatePackageAndTypeString(meth.GetOutputType())),
				Repeat:     repeated,
			})
		}
	}

	// helper funcs for templates
	funcMap := template.FuncMap{
		"caps":  strings.ToUpper,
		"title": strings.Title,
	}

	// generate the templates
  importsT := template.Must(template.New("imports").Funcs(funcMap).Parse(this.ActionFile.ImportTemplate))
	getT := template.Must(template.New("get").Funcs(funcMap).Parse(this.ActionFile.GetTemplate))
	listT := template.Must(template.New("list").Funcs(funcMap).Parse(this.ActionFile.ListTemplate))
	resetT := template.Must(template.New("reset").Funcs(funcMap).Parse(this.ActionFile.ResetTemplate))
	createT := template.Must(template.New("create").Funcs(funcMap).Parse(this.ActionFile.CreateTemplate))
	deleteT := template.Must(template.New("delete").Funcs(funcMap).Parse(this.ActionFile.DeleteTemplate))
	updateT := template.Must(template.New("update").Funcs(funcMap).Parse(this.ActionFile.UpdateTemplate))
	customT := template.Must(template.New("update").Funcs(funcMap).Parse(this.ActionFile.CustomTemplate))

	// append to output
	var output bytes.Buffer
  importsT.Execute(&output, nil)
	getT.Execute(&output, getEntities)
	listT.Execute(&output, listEntities)
	resetT.Execute(&output, resetEntities)
	createT.Execute(&output, createEntities)
	deleteT.Execute(&output, deleteEntities)
	updateT.Execute(&output, updateEntities)
	customT.Execute(&output, customEntities)

	// return the completed file
	return &File{
		Name:    "actions_pb.ts",
		Content: output.String(),
	}, nil
}
