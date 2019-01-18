package redux4

const EpicTemplate = `/* THIS FILE IS GENERATED FROM THE TOOL PROTOC-GEN-STATE  */
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


function noop() {
	return;
}

function createErrorObject(code: number|string|undefined, message: string): NodeJS.ErrnoException {
	var err: NodeJS.ErrnoException = new Error();
	err.message = message;
	if(code && typeof code == 'number') { err.code = code.toString(); }
	if(code && typeof code == 'string') { err.code = code; }
	return err;
}

{{range $i, $e := .}}
export const {{$e.Name}}Epic = (action$, store) => action$.pipe(
	filter(isActionOf(protocActions.{{$e.Name}}Request)),
	debounceTime({{$e.Debounce}}),
	map(({ payload, meta: { resolve = noop, reject = noop } }) => ({
		message: toMessage(payload, {{$e.InputType}}),
		resolve,
		reject,
	})),
	flatMap((request) => {
{{if $e.Repeat}} {{template "grpcStream" $e}} {{ else }} {{template "grpcUnary" $e}} {{end}}
		retry({{$e.Retries}}),
		timeout({{$e.Timeout}}),{{if $e.Updater}}
		map(obj => ({ ...obj } as { prev: {{$e.OutputType}}.AsObject, updated: {{$e.OutputType}}.AsObject } )),
		map(lib => {
			request.resolve(lib.prev, lib.updated);
			return protocActions.{{$e.Name}}Success(lib);
		}),{{else}}
		map((resObj: {{$e.OutputType}}.AsObject{{if $e.Repeat}}[]{{end}}) => {
			request.resolve(resObj);
			return protocActions.{{$e.Name}}Success(resObj);
		}){{end}}
		.catch(error => {
			const err: NodeJS.ErrnoException = createErrorObject(error.code, error.message);
			if(request.reject){ request.reject(err); }
			return Observable.of(protocActions.{{$e.Name}}Failure(err));
		})
	}),
	takeUntil(action$.filter(isActionOf(protocActions.{{$e.Name}}Cancel))),
	repeat();
{{end}}
{{define "grpcUnary"}}   return Observable
		.defer(() => new Promise((resolve, reject) => {
      {{if .Debug}}console.log('calling {{.FullMethodName}} with payload: ', request.message);{{end}}
			{{.Host}}
			{{.Auth}}
			grpc.unary({{.FullMethodName}}, {
				request: request.message,
				host: host,
				{{.AuthFollowup}}
				onEnd: (res: UnaryOutput<{{.OutputType}}>) => {
          {{if .Debug}}console.log('onEnd {{.FullMethodName}}: ', res.message);{{end}}
					if(res.status != grpc.Code.OK){
            {{if .Debug}}console.log('Error in epic -- status: ', res.status, ' message: ', res.statusMessage);{{end}}
						const err: NodeJS.ErrnoException = createErrorObject(res.status, res.statusMessage);
						reject(err);
					}
					if(res.message){
						resolve(res.message.toObject());
					}
				}
			});
		})){{end}}
{{define "grpcStream"}}   {{.Host}}
		return Observable
			.defer(() => new Promise((resolve, reject) => {
        {{if .Debug}}console.log('calling {{.FullMethodName}} with payload: ', request.message);{{end}}
				var arr: {{.OutputType}}.AsObject[] = [];
				const client = grpc.client({{.FullMethodName}}, {
					host: host,
				});
				client.onMessage((message: {{.OutputType}}) => {
          {{if .Debug}}console.log('in {{.FullMethodName}} streaming message: ', message.toObject());{{end}}
					arr.push(message.toObject());
				});
        {{if .Debug}}client.onEnd((code: grpc.Code, msg: string, trailers: grpc.Metadata) => {
          console.log('in {{.FullMethodName}} streaming onEnd: ', code, msg, trailers, request.message);{{else}}client.onEnd((code: grpc.Code, msg: string) => { {{end}}
					if (code != grpc.Code.OK) {
            {{if .Debug}}console.log('Error in streaming epic -- code: ', code, ' message: ', msg);{{end}}
						reject(createErrorObject(code, msg));
					}
					resolve(arr);
				});
				client.start({{.Auth}});
				client.send(request.message);
			})){{end}}

export const protocEpics = combineEpics({{range $i, $e := .}}
	{{$e.Name}}Epic,{{end}}
)`

