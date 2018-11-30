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

/// Sacrificed some code duplication for readability/maintainability sake

func createActionImports() string {
	return `/* THIS FILE IS GENERATED FROM THE TOOL PROTOC-GEN-STATE */
/* ANYTHING YOU EDIT WILL BE OVERWRITTEN IN FUTURE BUILDS */

import { createAction } from 'typesafe-actions';
import * as ProtocTypes from './protoc_types_pb';

`
}

const createTemplate = `{{range $i, $e := .}}
export const create{{$e.JsonName | title}}Request = createAction('PROTOC_CREATE_{{$e.JsonName | caps}}_REQUEST', (resolve) => {
	return ({{$e.JsonName}}: {{$e.InputType}}) => resolve({{$e.JsonName}})
});

export const create{{$e.JsonName | title}}RequestPromise = createAction('PROTOC_CREATE_{{$e.JsonName | caps}}_REQUEST_PROMISE', (res) => {
	return (
		{{$e.JsonName}}: {{$e.InputType}},
		resolve: (payload: {{$e.OutputType}}{{if $e.Repeat}}[]{{end}}) => void,
		reject: (error: NodeJS.ErrnoException) => void,
	) => res({ {{$e.JsonName}}, resolve, reject });
});

export const create{{$e.JsonName | title}}Success = createAction('PROTOC_CREATE_{{$e.JsonName | caps}}_SUCCESS', (resolve) => {
	return ({{$e.JsonName}}: {{$e.OutputType}}) => resolve({{$e.JsonName}})
});

export const create{{$e.JsonName | title}}Failure = createAction('PROTOC_CREATE_{{$e.JsonName | caps}}_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const create{{$e.JsonName | title}}Cancel = createAction('PROTOC_CREATE_{{$e.JsonName | caps}}_CANCEL');{{end}}
`

const updateTemplate = `{{range $i, $e := .}}
{{if $e.Repeat}}
export const update{{$e.JsonName | title}}Request = createAction('PROTOC_UPDATE_{{$e.JsonName | caps}}_REQUEST', (resolve) => {
	return (prev: {{$e.InputType}}, updated: {{$e.InputType}}) => resolve({prev, updated})
}){{else}}
export const update{{$e.JsonName | title}}Request = createAction('PROTOC_UPDATE_{{$e.JsonName | caps}}_REQUEST', (resolve) => {
	return ({{$e.JsonName}}: {{$e.InputType}}) => resolve({{$e.JsonName}})
}){{end}}

export const update{{$e.JsonName | title}}RequestPromise = createAction('PROTOC_UPDATE_{{$e.JsonName | caps}}_REQUEST_PROMISE', (res) => {
	return ({{if $e.Repeat}}
		prev: {{$e.InputType}},
		updated: {{$e.InputType}},
		resolve: (prev: {{$e.InputType}}, updated: {{$e.InputType}}) => void,
		reject: (error: NodeJS.ErrnoException) => void,
	) => res({ prev, updated, resolve, reject }){{else}}
		{{$e.JsonName}}: {{$e.InputType}},
		resolve: (payload: {{$e.InputType}}) => void,
		reject: (error: NodeJS.ErrnoException) => void,
	) => res({ {{$e.JsonName}}, resolve, reject }){{end}}
});

{{if $e.Repeat}}
export const update{{$e.JsonName | title}}Success = createAction('PROTOC_UPDATE_{{$e.JsonName | caps}}_SUCCESS', (resolve) => {
	return ({{$e.JsonName}}: { prev: {{$e.InputType}}, updated: {{$e.InputType}} }) => resolve({{$e.JsonName}})
}){{else}}
export const update{{$e.JsonName | title}}Success = createAction('PROTOC_UPDATE_{{$e.JsonName | caps}}_SUCCESS', (resolve) => {
	return ({{$e.JsonName}}: {{$e.InputType}}) => resolve({{$e.JsonName}})
}){{end}}

export const update{{$e.JsonName | title}}Failure = createAction('PROTOC_UPDATE_{{$e.JsonName | caps}}_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const update{{$e.JsonName | title}}Cancel = createAction('PROTOC_UPDATE_{{$e.JsonName | caps}}_CANCEL');{{end}}
`

const getTemplate = `{{range $i, $e := .}}
export const get{{$e.JsonName | title}}Request = createAction('PROTOC_GET_{{$e.JsonName | caps}}_REQUEST', (resolve) => {
	return ({{$e.JsonName}}: {{$e.InputType}}) => resolve({{$e.JsonName}})
});

export const get{{$e.JsonName | title}}RequestPromise = createAction('PROTOC_GET_{{$e.JsonName | caps}}_REQUEST_PROMISE', (res) => {
	return (
		{{$e.JsonName}}: {{$e.InputType}},
		resolve: (payload: {{$e.OutputType}}) => void,
		reject: (error: NodeJS.ErrnoException) => void,
	) => res({ {{$e.JsonName}}, resolve, reject });
});

export const get{{$e.JsonName | title}}Success = createAction('PROTOC_GET_{{$e.JsonName | caps}}_SUCCESS', (resolve) => {
	return ({{$e.JsonName}}: {{$e.OutputType}}) => resolve({{$e.JsonName}})
});

export const get{{$e.JsonName | title}}Failure = createAction('PROTOC_GET_{{$e.JsonName | caps}}_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const get{{$e.JsonName | title}}Cancel = createAction('PROTOC_GET_{{$e.JsonName | caps}}_CANCEL');{{end}}
`

const listTemplate = `{{range $i, $e := .}}
export const list{{$e.JsonName | title}}Request = createAction('PROTOC_LIST_{{$e.JsonName | caps}}_REQUEST', (resolve) => {
	return ({{$e.JsonName}}: {{$e.InputType}}) => resolve({{$e.JsonName}})
});

export const list{{$e.JsonName | title}}RequestPromise = createAction('PROTOC_LIST_{{$e.JsonName | caps}}_REQUEST_PROMISE', (res) => {
	return (
		{{$e.JsonName}}: {{$e.InputType}},
		resolve: (payload: {{$e.OutputType}}[]) => void,
		reject: (error: NodeJS.ErrnoException) => void,
	) => res({ {{$e.JsonName}}, resolve, reject });
});

export const list{{$e.JsonName | title}}Success = createAction('PROTOC_LIST_{{$e.JsonName | caps}}_SUCCESS', (resolve) => {
	return ({{$e.JsonName}}: {{$e.OutputType}}[]) => resolve({{$e.JsonName}})
});

export const list{{$e.JsonName | title}}Failure = createAction('PROTOC_LIST_{{$e.JsonName | caps}}_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const list{{$e.JsonName | title}}Cancel = createAction('PROTOC_LIST_{{$e.JsonName | caps}}_CANCEL');{{end}}
`

const deleteTemplate = `{{range $i, $e := .}}
export const delete{{$e.JsonName | title}}Request = createAction('PROTOC_DELETE_{{$e.JsonName | caps}}_REQUEST', (resolve) => {
	return ({{$e.JsonName}}: {{$e.InputType}}) => resolve({{$e.JsonName}})
});

export const delete{{$e.JsonName | title}}RequestPromise = createAction('PROTOC_DELETE_{{$e.JsonName | caps}}_REQUEST_PROMISE', (res) => {
	return (
		{{$e.JsonName}}: {{$e.InputType}},
		resolve: (payload: {{$e.OutputType}}{{if $e.Repeat}}[]{{end}}) => void,
		reject: (error: NodeJS.ErrnoException) => void,
	) => res({ {{$e.JsonName}}, resolve, reject });
});

export const delete{{$e.JsonName | title}}Success = createAction('PROTOC_DELETE_{{$e.JsonName | caps}}_SUCCESS', (resolve) => {
	return ({{$e.JsonName}}: {{$e.OutputType}}) => resolve({{$e.JsonName}})
});

export const delete{{$e.JsonName | title}}Failure = createAction('PROTOC_DELETE_{{$e.JsonName | caps}}_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const delete{{$e.JsonName | title}}Cancel = createAction('PROTOC_DELETE_{{$e.JsonName | caps}}_CANCEL');{{end}}
`

const customTemplate = `{{range $i, $e := .}}
export const custom{{$e.JsonName | title}}Request = createAction('PROTOC_CUSTOM_{{$e.JsonName | caps}}_REQUEST', (resolve) => {
	return ({{$e.JsonName}}: {{$e.InputType}}) => resolve({{$e.JsonName}})
});

export const custom{{$e.JsonName | title}}RequestPromise = createAction('PROTOC_CUSTOM_{{$e.JsonName | caps}}_REQUEST_PROMISE', (res) => {
	return (
		{{$e.JsonName}}: {{$e.InputType}},
		resolve: (payload: {{$e.OutputType}}{{if $e.Repeat}}[]{{end}}) => void,
		reject: (error: NodeJS.ErrnoException) => void,
	) => res({ {{$e.JsonName}}, resolve, reject });
});

export const custom{{$e.JsonName | title}}Success = createAction('PROTOC_CUSTOM_{{$e.JsonName | caps}}_SUCCESS', (resolve) => {
	return ({{$e.JsonName}}: {{$e.OutputType}}{{if $e.Repeat}}[]{{end}}) => resolve({{$e.JsonName}})
});

export const custom{{$e.JsonName | title}}Failure = createAction('PROTOC_CUSTOM_{{$e.JsonName | caps}}_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const custom{{$e.JsonName | title}}Cancel = createAction('PROTOC_CUSTOM_{{$e.JsonName | caps}}_CANCEL');{{end}}
`

const resetTemplate = `{{range $i, $e := .}}
export const reset{{$e.JsonName | title}} = createAction('PROTOC_RESET_{{$e.JsonName | caps}}');{{end}}
`

const header = `/* THIS FILE IS GENERATED FROM THE TOOL PROTOC-GEN-STATE */
/* ANYTHING YOU EDIT WILL BE OVERWRITTEN IN FUTURE BUILDS */

import { createAction } from 'typesafe-actions';
import * as ProtocTypes from './protoc_types_pb';

`

type ActionEntity struct {
	JsonName   string
	InputType  string
	OutputType string
	Repeat     bool
}

func CreateActionFile(stateFields []*gp.FieldDescriptorProto, customFields []*gp.FieldDescriptorProto, serviceFiles []*gp.FileDescriptorProto) (*File, error) {
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
		methods, err := GetFieldOptionsString(field, state.E_Method)
		if err != nil {
			return nil, fmt.Errorf("Error getting field level annotations: %v", err)
		}

		// using custom method annotation anywhere but CustomActions is an error
		if methods.GetCustom() != "" {
			return nil, fmt.Errorf("Method annotation 'custom' provided outside of CustomActions message. Correct this by removing the '(method).custom' annotation on field: %s", field.GetName())
		}

		// loop through CLUDG verbs
		var meth *gp.MethodDescriptorProto
		for c := CREATE; c < CRUD_MAX; c++ {
			// verify the annotation exists in the methods struct
			crudAnnotation := GetAnnotation(methods, c, repeated)
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
					InputType:  fmt.Sprintf("ProtocTypes%s.AsObject", meth.GetInputType()),
					OutputType: fmt.Sprintf("ProtocTypes%s.AsObject", meth.GetOutputType()),
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
		methods, err := GetFieldOptionsString(field, state.E_Method)
		if err != nil {
			return nil, fmt.Errorf("Error getting field level annotations: %v", err)
		}

		// only custom method annotations allowed in CustomActions
		for c := CREATE; c < CRUD_MAX; c++ {
			if GetAnnotation(methods, c, repeated) != "" {
				return nil, fmt.Errorf("Invalid method annotation. Only method annotation '(method).custom' is allowed within CustomActions message. Correct the annotation on field: %s", *field.JsonName)
			}
		}

		// and it must have a custom annotation
		var meth *gp.MethodDescriptorProto
		crudAnnotation := methods.GetCustom()
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
	getT := template.Must(template.New("get").Funcs(funcMap).Parse(getTemplate))
	listT := template.Must(template.New("list").Funcs(funcMap).Parse(listTemplate))
	resetT := template.Must(template.New("reset").Funcs(funcMap).Parse(resetTemplate))
	createT := template.Must(template.New("create").Funcs(funcMap).Parse(createTemplate))
	deleteT := template.Must(template.New("delete").Funcs(funcMap).Parse(deleteTemplate))
	updateT := template.Must(template.New("update").Funcs(funcMap).Parse(updateTemplate))
	customT := template.Must(template.New("update").Funcs(funcMap).Parse(customTemplate))

	// append to output
	var output bytes.Buffer
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
		Content: createActionImports() + output.String(),
	}, nil
}
