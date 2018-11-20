package main

import (
	"bytes"
	"fmt"
	gp "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/tcncloud/protoc-gen-state/state"
	"strings"
	"text/template"
)

func createActionImports() string {
	return `import { createAction } from 'typesafe-actions';
import * as ProtocTypes from './protoc_types_pb';

`
}

// const requestTemplate = `export const {{$entity.Crud}}{{$entity.JsonName}}Request = createAction('{{$entity.TypeName}}_REQUEST', (resolve) => {
// 	return ({{$entity.JsonName}}: {{$entity.InputType}}) => resolve({{$entity.JsonName}})
// });`
// const requestPromiseTemplate = `export const {{$entity.Crud}}{{$entity.JsonName}}RequestPromise = createAction('{{$entity.TypeName}}_REQUEST_PROMISE', (resolve) => {
// 	return (
// 		{{$entity.JsonName}}: {{$entity.InputType}},
// 		resolve: (payload: {{$entity.OutputName}}{{if $entity.Repeated}}[]{{else}}{{end}}) => void,
// 		reject: (error: NodeJS.ErrnoException) => void,
// 	) => resolve({ {{$entity.JsonName}}, resolve, reject })
// });`
// const successTemplate = `export const {{$entity.Crud}}{{$entity.JsonName}}Success = createAction('{{$entity.TypeName}}_SUCCESS', (resolve) => {
// 	return ({{$entity.JsonName}}: {{$entity.OutputType}}) => resolve({{$entity.JsonName}})
// });`
// const failureTemplate = `export const {{$entity.Crud}}{{$entity.JsonName}}Failure = createAction('{{$entity.TypeName}}_FAILURE', (resolve) => {
// 	return (error: NodeJS.ErrnoException) => resolve(error)
// });`
// const cancelTemplate = `export const {{$entity.Crud}}{{$entity.JsonName}}Cancel = createAction('{{$entity.TypeName}}_Cancel');`

const updateTemplate = `{ {{range $i, $entity := .}}
export const update{{$entity.JsonName}}Request = createAction('PROTOC_UPDATE_{{$entity.JsonName}}_REQUEST', (resolve) => {
	return (prev: {{$entity.InputType}}, updated: {{$entity.InputType}}) => resolve({prev, updated})
})

export const update{{$entity.JsonName}}RequestPromise = createAction('PROTOC_UPDATE_{{$entity.JsonName}}_REQUEST_PROMISE', (res) => {
	return (
		prev: {{$entity.InputType}},
		updated: {{$entity.InputType}},
		resolve: (prev: {{$entity.InputType}}, updated: {{$entity.InputType}}j) => void,
		reject: (error: NodeJS.ErrnoException) => void,
	) => res({ prev, updated, resolve, reject })
});

export const update{{$entity.JsonName}}Success = createAction('PROTOC_UPDATE_{{$entity.JsonName}}_SUCCESS', (resolve) => {
	return ({{$entity.JsonName}}: { prev: {{$entity.InputType}}, updated: {{$entity.InputType}} }) => resolve({{$entity.JsonName}})
})

export const update{{$entity.JsonName}}Failure = createAction('PROTOC_UPDATE_{{$entity.JsonName}}_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const update{{$entity.JsonName}}Cancel = createAction('PROTOC_UPDATE_{{$entity.JsonName}}_CANCEL');
}
`

const actionTemplate = `/* THIS FILE IS GENERATED FROM THE TOOL PROTOC-GEN-STATE */
/* ANYTHING YOU EDIT WILL BE OVERWRITTEN IN FUTURE BUILDS */

import { createAction } from 'typesafe-actions';
import * as ProtocTypes from './protoc_types_pb';

// TODO
`

type ActionEntity struct {
	JsonName   string
	InputType  string
	OutputType string
}

func CreateActionFile(stateFields []*gp.FieldDescriptorProto, customFields []*gp.FieldDescriptorProto, serviceFiles []*gp.FileDescriptorProto) (*File, error) {
	actionEntities := []*ActionEntity{}

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
			// verify the annotation exists
			if GetAnnotation(methods, c, repeated) != "" {
				meth, err = FindMethodDescriptor(serviceFiles, GetAnnotation(methods, c, repeated))
			}

			// if method is still empty, the method could not be found so do nothing
			if meth != nil {
				// make an action block for each side effect
				actionEntities = append(actionEntities, &ActionEntity{
					JsonName:   *field.JsonName,
					InputType:  meth.GetInputType(),
					OutputType: meth.GetOutputType(),
				})
			}
		}

		// create a reset action too
		// resetBlock, err := resetTemplate(field.GetJsonName())
		// if err != nil {
		// 	return nil, err
		// }
		// output += resetBlock
	}

	// generate
	tmpl := template.Must(template.New("update").Parse(updateTemplate))
	var output bytes.Buffer
	tmpl.Execute(&output, actionEntities)

	return &File{
		Name:    "actions_pb.ts",
		Content: output.String(),
	}, nil
}

// func CreateActionBlock(meth *gp.MethodDescriptorProto, c Crud, payloadName string, s SideEffect, repeated bool) (string, error) {
// 	switch c {
// 	case CREATE:
// 		switch s {
// 		case REQUEST:
// 			return requestTemplate(c, s, payloadName, meth.GetInputType(), repeated)
// 		case SUCCESS:
// 			return "", nil
// 		case FAILURE:
// 			return "", nil
// 		default:
// 			return "", nil
// 		}
// 	case UPDATE:
// 		switch s {
// 		case REQUEST:
// 			return requestTemplate(c, s, payloadName, meth.GetInputType(), repeated)
// 		case SUCCESS:
// 			return "", nil
// 		case FAILURE:
// 			return "", nil
// 		default:
// 			return "", nil
// 		}
// 	case GET:
// 		switch s {
// 		case REQUEST:
// 			return requestTemplate(c, s, payloadName, meth.GetInputType(), repeated)
// 		case SUCCESS:
// 			return "", nil
// 		case FAILURE:
// 			return "", nil
// 		default:
// 			return "", nil
// 		}
// 	case DELETE:
// 		switch s {
// 		case REQUEST:
// 			return requestTemplate(c, s, payloadName, meth.GetInputType(), repeated)
// 		case SUCCESS:
// 			return "", nil
// 		case FAILURE:
// 			return "", nil
// 		default:
// 			return "", nil
// 		}
// 	}

// 	// make a request action with promise callbacks
// 	if s == REQUEST {
// 		// TODO
// 	}

// 	return "", nil
// }

type Reset struct {
	Name string
	Type string
}

func resetTemplate(payloadName string) (string, error) {
	RESET := "export const {{.Name}} = createAction('{{.Type}}');\n"

	r := Reset{
		"reset" + payloadName,
		"PROTOC_RESET_" + strings.ToUpper(payloadName),
	}
	t, err := template.New("reset").Parse(RESET)

	var result bytes.Buffer
	if err = t.Execute(&result, r); err != nil {
		return "", fmt.Errorf("Failed executing reset template: %v", err)
	}

	return result.String(), nil
}
