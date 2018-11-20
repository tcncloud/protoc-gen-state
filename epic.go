package main

import (
	"bytes"
	"text/template"

	gp "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

const epicTemplate = `/* THIS FILE IS GENERATED FROM THE TOOL PROTOC-GEN-STATE  */
/* ANYTHING YOU EDIT WILL BE OVERWRITTEN IN FUTURE BUILDS */

import { combineEpics } from 'redux-observable';
import { isActionOf } from 'typesafe-actions';
import { Observable } from 'rxjs';
import _ from 'lodash';
import { grpc } from 'grpc-web-client';
import { UnaryOutput } from 'grpc-web-client/dist/unary';
import 'rxjs/add/observable/dom/ajax';
import { toMessage } from './to_message_pb';
import * as protocActions from './actions_pb';
import * as ProtocTypes from './protoc_types_pb';
import * as ProtocServices from './protoc_services_pb';

{{range $i, $e := .}}
export const createLibraryEpic = (action$, store) => action$
	.filter(isActionOf([
		protocActions.{{$
{{end}}
export interface ProtocState { {{range $i, $entity := .}}
{{$entity.FieldName}}: {
	isLoading: boolean;
	error: { code: string; message: string; },
	{{if $entity.Repeated}}value: ProtocTypes{{$entity.FullTypeName}}.AsObject;
	{{else}}value: ProtocTypes{{$entity.FullTypeName}}.AsObject | null;{{end}}
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

type EpicEntity struct {
	FieldName    string
	FullTypeName string
	Repeated     bool
}

func CreateEpicFile(stateFields []*gp.FieldDescriptorProto) (*File, error) {
	stateEntities := []*StateEntity{}

	// transform stateFields into our StateEntity implementation so template can read values
	for _, entity := range stateFields {
		stateEntities = append(stateEntities, &StateEntity{
			FieldName:    entity.GetJsonName(),
			FullTypeName: entity.GetTypeName(),
			Repeated:     entity.GetLabel() == 3,
		})
	}

	tmpl := template.Must(template.New("state").Parse(stateTemplate))

	var output bytes.Buffer
	tmpl.Execute(&output, stateEntities)

	return &File{
		Name:    "epics_pb.ts",
		Content: output.String(),
	}, nil
}
