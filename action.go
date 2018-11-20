package main

import (
	"bytes"
	"fmt"
	gp "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/tcncloud/protoc-gen-state/state"
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
export const create{{$e.JsonName}}Request = createAction('PROTOC_CREATE_{{$e.JsonName}}_REQUEST', (resolve) => {
	return ({{$e.JsonName}}: {{$e.InputType}}) => resolve({{$e.JsonName}})
});

export const create{{$e.JsonName}}RequestPromise = createAction('PROTOC_CREATE_{{$e.JsonName}}_REQUEST_PROMISE', (resolve) => {
	return (
		{{$e.JsonName}}: {{$e.InputType}},
		resolve: (payload: {{$e.OutputType}}) => void,
		reject: (error: NodeJS.ErrnoException) => void,
	) => res({ {{$e.JsonName}}, resolve, reject });
});

export const create{{$e.JsonName}}Success = createAction('PROTOC_CREATE_{{$e.JsonName}}_SUCCESS', (resolve) => {
	return ({{$e.JsonName}}: {{$e.OutputType}}) => resolve({{$e.JsonName}})
});

export const create{{$e.JsonName}}Failure = createAction('PROTOC_CREATE_{{$e.JsonName}}_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const create{{$e.JsonName}}Cancel = createAction('PROTOC_CREATE_{{$e.JsonName}}_CANCEL');{{end}}
`

const updateTemplate = `{{range $i, $e := .}}
export const update{{$e.JsonName}}Request = createAction('PROTOC_UPDATE_{{$e.JsonName}}_REQUEST', (resolve) => {
	return (prev: {{$e.InputType}}, updated: {{$e.InputType}}) => resolve({prev, updated})
})

export const update{{$e.JsonName}}RequestPromise = createAction('PROTOC_UPDATE_{{$e.JsonName}}_REQUEST_PROMISE', (res) => {
	return (
		prev: {{$e.InputType}},
		updated: {{$e.InputType}},
		resolve: (prev: {{$e.InputType}}, updated: {{$e.InputType}}j) => void,
		reject: (error: NodeJS.ErrnoException) => void,
	) => res({ prev, updated, resolve, reject })
});

export const update{{$e.JsonName}}Success = createAction('PROTOC_UPDATE_{{$e.JsonName}}_SUCCESS', (resolve) => {
	return ({{$e.JsonName}}: { prev: {{$e.InputType}}, updated: {{$e.InputType}} }) => resolve({{$e.JsonName}})
})

export const update{{$e.JsonName}}Failure = createAction('PROTOC_UPDATE_{{$e.JsonName}}_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const update{{$e.JsonName}}Cancel = createAction('PROTOC_UPDATE_{{$e.JsonName}}_CANCEL');{{end}}
`

const getTemplate = `{{range $i, $e := .}}
export const get{{$e.JsonName}}Request = createAction('PROTOC_GET_{{$e.JsonName}}_REQUEST', (resolve) => {
	return ({{$e.JsonName}}: {{$e.InputType}}) => resolve({{$e.JsonName}})
});

export const get{{$e.JsonName}}RequestPromise = createAction('PROTOC_GET_{{$e.JsonName}}_REQUEST_PROMISE', (resolve) => {
	return (
		{{$e.JsonName}}: {{$e.InputType}},
		resolve: (payload: {{$e.OutputType}}) => void,
		reject: (error: NodeJS.ErrnoException) => void,
	) => res({ {{$e.JsonName}}, resolve, reject });
});

export const get{{$e.JsonName}}Success = createAction('PROTOC_GET_{{$e.JsonName}}_SUCCESS', (resolve) => {
	return ({{$e.JsonName}}: {{$e.OutputType}}) => resolve({{$e.JsonName}})
});

export const get{{$e.JsonName}}Failure = createAction('PROTOC_GET_{{$e.JsonName}}_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const get{{$e.JsonName}}Cancel = createAction('PROTOC_GET_{{$e.JsonName}}_CANCEL');{{end}}
`

const listTemplate = `{{range $i, $e := .}}
export const list{{$e.JsonName}}Request = createAction('PROTOC_LIST_{{$e.JsonName}}_REQUEST', (resolve) => {
	return ({{$e.JsonName}}: {{$e.InputType}}) => resolve({{$e.JsonName}})
});

export const list{{$e.JsonName}}RequestPromise = createAction('PROTOC_LIST_{{$e.JsonName}}_REQUEST_PROMISE', (resolve) => {
	return (
		{{$e.JsonName}}: {{$e.InputType}},
		resolve: (payload: {{$e.OutputType}}) => void,
		reject: (error: NodeJS.ErrnoException) => void,
	) => res({ {{$e.JsonName}}, resolve, reject });
});

export const list{{$e.JsonName}}Success = createAction('PROTOC_LIST_{{$e.JsonName}}_SUCCESS', (resolve) => {
	return ({{$e.JsonName}}: {{$e.OutputType}}) => resolve({{$e.JsonName}})
});

export const list{{$e.JsonName}}Failure = createAction('PROTOC_LIST_{{$e.JsonName}}_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const list{{$e.JsonName}}Cancel = createAction('PROTOC_LIST_{{$e.JsonName}}_CANCEL');{{end}}
`

const deleteTemplate = `{{range $i, $e := .}}
export const delete{{$e.JsonName}}Request = createAction('PROTOC_DELETE_{{$e.JsonName}}_REQUEST', (resolve) => {
	return ({{$e.JsonName}}: {{$e.InputType}}) => resolve({{$e.JsonName}})
});

export const delete{{$e.JsonName}}RequestPromise = createAction('PROTOC_DELETE_{{$e.JsonName}}_REQUEST_PROMISE', (resolve) => {
	return (
		{{$e.JsonName}}: {{$e.InputType}},
		resolve: (payload: {{$e.OutputType}}) => void,
		reject: (error: NodeJS.ErrnoException) => void,
	) => res({ {{$e.JsonName}}, resolve, reject });
});

export const delete{{$e.JsonName}}Success = createAction('PROTOC_DELETE_{{$e.JsonName}}_SUCCESS', (resolve) => {
	return ({{$e.JsonName}}: {{$e.OutputType}}) => resolve({{$e.JsonName}})
});

export const delete{{$e.JsonName}}Failure = createAction('PROTOC_DELETE_{{$e.JsonName}}_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const delete{{$e.JsonName}}Cancel = createAction('PROTOC_DELETE_{{$e.JsonName}}_CANCEL');{{end}}
`

const resetTemplate = `{{range $i, $e := .}}
export const reset{{$e.JsonName}} = createAction('PROTOC_RESET_{{$e.JsonName}}');{{end}}
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
}

func CreateActionFile(stateFields []*gp.FieldDescriptorProto, customFields []*gp.FieldDescriptorProto, serviceFiles []*gp.FileDescriptorProto) (*File, error) {
	createEntities := []*ActionEntity{}
	updateEntities := []*ActionEntity{}
	deleteEntities := []*ActionEntity{}
	getEntities := []*ActionEntity{}
	listEntities := []*ActionEntity{}
	resetEntities := []*ActionEntity{}

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
					InputType:  meth.GetInputType(),
					OutputType: meth.GetOutputType(),
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

		// create a reset action too
		resetEntities = append(resetEntities, &ActionEntity{
			JsonName: *field.JsonName,
		})
	}

	// generate the templates
	createT := template.Must(template.New("create").Parse(createTemplate))
	deleteT := template.Must(template.New("delete").Parse(deleteTemplate))
	updateT := template.Must(template.New("update").Parse(updateTemplate))
	getT := template.Must(template.New("get").Parse(getTemplate))
	listT := template.Must(template.New("list").Parse(listTemplate))
	resetT := template.Must(template.New("reset").Parse(resetTemplate))

	// append to output
	var output bytes.Buffer
	createT.Execute(&output, createEntities)
	deleteT.Execute(&output, deleteEntities)
	updateT.Execute(&output, updateEntities)
	getT.Execute(&output, getEntities)
	listT.Execute(&output, listEntities)
	resetT.Execute(&output, resetEntities)

	// return the completed file
	return &File{
		Name:    "actions_pb.ts",
		Content: createActionImports() + output.String(),
	}, nil
}
