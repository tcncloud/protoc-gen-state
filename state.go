package main

import (
  "bytes"
  // "fmt"
  "text/template"

	gp "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

const stateTemplate = `/* THIS FILE IS GENERATED FROM THE TOOL PROTOC-GEN-STATE  */
/* ANYTHING YOU EDIT WILL BE OVERWRITTEN IN FUTURE BUILDS */

import * as ProtocTypes from './protoc_types_pb';

export interface ProtocState { {{range $i, $entity := .}}
  {{$entity.FieldName}}: {
    isLoading: boolean;
    error: { code: string; message: string; },
    {{if $entity.Repeated}}value: {{$entity.FullTypeName}};
    {{else}}value: {{$entity.FullTypeName}} | null;{{end}}
  },
  {{end}}
}

export const initialProtocState : ProtocState = { {{range $i, $entity := .}}
  {{$entity.FieldName}}: {
    isLoading: false,
    error: null,
    {{if $entity.Repeated}}value: [],
    {{else}}value: null,{{end}}
  },
  {{end}}
}
`
// const stateTemplate = `/* THIS FILE IS GENERATED FROM THE TOOL PROTOC-GEN-STATE  */
// /* ANYTHING YOU EDIT WILL BE OVERWRITTEN IN FUTURE BUILDS */

// import * as ProtocTypes from './protoc_types_pb';

// export interface ProtocState {
//   {{range $i, $StateEntity := .}}
//   {{$StateEntity.FieldName}}
//   {{end}}
// }
// `

type StateEntity struct {
  FieldName string
  FullTypeName string
  Repeated bool
}


func CreateStateFile(stateFields []*gp.FieldDescriptorProto) (*File, error) {
  stateEntities := []*StateEntity{}

  for _, entity := range stateFields {
    newEntity := &StateEntity{
      FieldName: *entity.Name,
      FullTypeName: *entity.TypeName,
      Repeated: entity.GetLabel() == 3,
    }
    stateEntities = append(stateEntities, newEntity)
  }

  // fmt.Println("output: ", stateEntities[0].fieldName)
  tmpl := template.Must(template.New("state").Parse(stateTemplate))

  var output bytes.Buffer
  tmpl.Execute(&output, stateEntities)

	return &File{
		Name:    "state_pb.ts",
		Content: output.String(),
	}, nil
}
