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
	return `import { createAction from 'typesafe-actions';
import * as ProtocTypes from './protoc_types_pb';

`
}

type Reset struct {
	Name string
	Type string
}

func GetAnnotation(meth state.StringFieldOptions, crud Crud, repeated bool) string {
	switch crud {
	case CREATE:
		return meth.GetCreate()
	case GET:
		{
			if repeated {
				return meth.GetList()
			} else {
				return meth.GetGet()
			}
		}
	case UPDATE:
		return meth.GetUpdate()
	case DELETE:
		return meth.GetDelete()
	default:
		return ""
	}
}

func CreateActionFile(stateFields []*gp.FieldDescriptorProto, customFields []*gp.FieldDescriptorProto, serviceFiles []*gp.FileDescriptorProto) (*File, error) {
	output := createActionImports()
	for _, field := range stateFields {
		repeated := field.GetLabel() == 3
		// get the method annotations
		methods, err := GetFieldOptionsString(field, state.E_Method)
		// // get the method timeout and retry field level override
		// methodTimeout, err := GetFieldOptionsInt(field, state.E_MethodTimeout)
		// methodRetries, err := GetFieldOptionsInt(field, state.E_MethodRetries)
		// // get timeout and retries field level override
		// timeout, err := GetFieldAnnotationInt(field, state.E_Timeout)
		// retries, err := GetFieldAnnotationInt(field, state.E_Retries)

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
		}

		// if method is still empty, the method could not be found so do nothing
		if meth != nil {
			// // make an action block for each side effect
			// for s := REQUEST; s < SIDE_EFFECT_MAX; s++ {
			//   actionBlock, err := CreateActionBlock(meth, c, field.GetJsonName(), s, repeated)
			//   if err != nil {
			// 	   return nil, err
			//   }
			//   output += actionBlock
			// }
		}

		// create a reset action too
		resetBlock, err := resetTemplate(field.GetJsonName())
		if err != nil {
			return nil, err
		}
		output += resetBlock
	}

	return &File{
		Name:    "actions_pb.ts",
		Content: output,
	}, nil
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
