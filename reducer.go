package main

import (
  "bytes"
  "fmt"
  "strings"
  "text/template"

	gp "github.com/golang/protobuf/protoc-gen-go/descriptor"
  "github.com/tcncloud/protoc-gen-state/state"
)
// cludg also has reset
// Should try out subtemplates
// TODO make sure maps are supported


const reducerTemplate = `{{define "getRequest"}}case getType(protocActions['get{{.Name}}Request']):
      return {
        ...state,
        {{.Name}}: {
          ...state.{{.Name}},
          isLoading: true,
        }
      }{{end}}

{{define "getSuccess"}}case getType(protocActions['get{{.Name}}Success']):
      var {{.Name}}Value: ProtocState["{{.Name}}"]["value"] = action.payload;
      return {
        ...state,
        {{.Name}}: {
          ...state.{{.Name}},
          isLoading: false,
          value: {{.Name}}Value,
          error: initialProtocState.{{.Name}}.error,
        }
      }{{end}}

/* THIS FILE IS GENERATED FROM THE TOOL PROTOC-GEN-STATE  */
/* ANYTHING YOU EDIT WILL BE OVERWRITTEN IN FUTURE BUILDS */

import { getType, ActionType } from 'typesafe-actions';
import _ from 'lodash';
import * as protocActions from './actions_pb';
import * as ProtocTypes from './protoc_types_pb';
import { ProtocState, initialProtocState } from './state_pb';

type RootAction = ActionType<typeof protocActions>;

export function protocReducer(state: ProtocState = initialProtocState, action: RootAction) {
  switch(action.type) { {{range $i, $entity := .}}
    {{if eq $entity.SwitchCase 0}}{{template "getRequest" $entity}}{{end}}
    {{if eq $entity.SwitchCase 1}}{{template "getSuccess" $entity}}{{end}}
    {{end}}
    default: return state;
  }
};
}
`

type SwitchCaseEnum int

const (
  getRequest    SwitchCaseEnum = 0
  getSuccess    SwitchCaseEnum = 1
  getFailure    SwitchCaseEnum = 2
  getCancel     SwitchCaseEnum = 3
  listRequest   SwitchCaseEnum = 4
  listSuccess   SwitchCaseEnum = 5
  listFailure   SwitchCaseEnum = 6
  listCancel    SwitchCaseEnum = 7
  reset         SwitchCaseEnum = 8
)


type ReducerEntity struct {
  SwitchCase SwitchCaseEnum
  Name string
  CludgName string
}

func CreateReducerFile(stateFields []*gp.FieldDescriptorProto) (*File, error) {
  reducerEntities := []*ReducerEntity{}

  // TODO
  // loop through stateFields
      // find the annotations defined for each field
  methods, err := GetFieldOptionsString(stateFields[0], state.E_Method)
  if err != nil {
    fmt.Println("woops")
  }
  ann := GetAnnotation(methods, 0, stateFields[0].GetLabel() == 3)
  if methods.GetList() == "" {
    fmt.Println("fuuuuck", "\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
  } else {
    fmt.Println(ann, "\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
  }
  
  for _, entity := range stateFields {
    reducerEntities = append(reducerEntities, &ReducerEntity{
      SwitchCase: 1,
      CludgName: "get" + strings.Title(entity.GetJsonName()) + "Request",
      Name: entity.GetJsonName(),
    })
  }

  tmpl := template.Must(template.New("reducer").Parse(reducerTemplate))

  var output bytes.Buffer
  tmpl.Execute(&output, reducerEntities)

	return &File{
		Name:    "reducer_pb.ts",
		Content: output.String(),
	}, nil
}
