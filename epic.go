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

function createErrorObject(code: number|string|undefined, message: string): NodeJS.ErrnoException {
	var err: NodeJS.ErrnoException = new Error();
	err.message = message;
	if(code && typeof code == 'number') { err.code = code.toString(); }
	if(code && typeof code == 'string') { err.code = code; }
	return err;
}

{{range $i, $e := .}}
export const {{$e.Name}}Epic = (action$, store) => action$
	.filter(isActionOf([
		protocActions.{{$e.Name}}Request,
		protocActions.{{$e.Name}}RequestPromise,
	]))
	.debounceTime({{$e.Debounce}})
	.map((action) => {
		if(action.payload && action.payload.resolve && action.payload.reject){
			return {
				...action.payload,
				request: toMessage(action.payload.library, ProtocTypes.readinglist.Book)
			}
		} else {
			return { request: toMessage(action.payload, ProtocTypes.readinglist.Book) }
		}
	})
	.flatMap((action) => {
		{{if $e.Repeat}}
			{{template "grpcStream" .}}
		{{ else }}
			{{template "grpcUnary" .}}
		{{end}}
			.retry({{$e.Retry}})
			.timeout({{$e.Timeout}})
			.map(resObj => {
				if(action.resolve){
					action.resolve(resObj as {{$e.InputType}}{{if $e.Repeat}}[]{{end}});
				}
				reutrn protocActions.{{$e.Name}}Success(resObj as {{$e.InputType}}{{if $e.Repeat}}[]{{end}});
			})
			.catch(error => {
				const err: NodeJS.ErrnoException = createErrorObject(error.code, error.message);
				if(action.reject){ action.reject(err); }
				return Observable.of(protocActions.{{$e.Name}}Failure(err));
			})
	})
	.takeUntil(action$.filter(isActionOf(protocActions.{{$e.Name}}Cancel)))
	.repeat();
{{end}
`

const grpcUnary = `return Observable
	.defer(() => new Promise((resolve, reject) => {
		{{.Host}}
		{{.Auth}}
		grpc.unary({{$e.FullMethodName}}, {
			request: action.request,
			host: host,
			onEnd: (res: UnaryOutput<{{$e.OutputType}}>) => {
				if(res.status != grpc.Code.OK){
					const err: NodeJS.ErrnoException = createErrorObject(res.status, res.statusMessage);
					reject(err);
				}
				if(res.message){
					resolve(res.message.toObject());
				}
			}
		});
	}))`

const grpcStream = `return Observable
	.defer(() => new Promise((resolve, reject) => {
		var arr: {{$e.OutputType}}[] = [];
		const client = grpc.client({{$e.FullMethodName}}, {
			host: host,
		});
		client.onMessage((message: {{$e.OutputType}}) => {
			arr.push(message.toObject());
		})
		client.onEnd((code: grpc.Code, msg: string) => {
			if (code != grpc.Code.OK) {
				reject(createErrorObject(code, msg));
			}
			resolve(arr);
		});
		client.start({{if .Auth == ""}}new grpc.Metadata({ "Authorization": Bearer ${store.getState().{{.Auth}}}{{end}} }));
		client.send(action.request)
	}))`

type EpicEntity struct {
	Name           string
	InputType      string
	OutputType     string
	FullMethodName string
	Debounce       int64
	Timeout        int64
	Retries        int64
	Repeat         bool
}

func CreateEpicFile(stateFields []*gp.FieldDescriptorProto) (*File, error) {
	epicEntities := []*EpicEntity{}

	// transform stateFields into our EpicEntity implementation so template can read values
	for _, _ = range stateFields {
		epicEntities = append(epicEntities, &EpicEntity{
			//TODO
		})
	}

	tmpl := template.Must(template.New("state").Parse(stateTemplate))

	var output bytes.Buffer
	tmpl.Execute(&output, epicEntities)

	return &File{
		Name:    "epics_pb.ts",
		Content: output.String(),
	}, nil
}
