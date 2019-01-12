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
	"strconv"
	"strings"
	"text/template"
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
export const {{$e.Name}}Epic = (action$, store) => action$
	.filter(isActionOf(protocActions.{{$e.Name}}Request))
	.debounceTime({{$e.Debounce}})
	.map(({ payload, meta: { resolve = noop, reject = noop } }) => ({
		message: toMessage(payload, {{$e.InputType}}),
		resolve,
		reject,
	}))
	.flatMap((request) => {
{{if $e.Repeat}} {{template "grpcStream" $e}} {{ else }} {{template "grpcUnary" $e}} {{end}}
		.retry({{$e.Retries}})
		.timeout({{$e.Timeout}}){{if $e.Updater}}
		.map(obj => ({ ...obj } as { prev: {{$e.OutputType}}.AsObject, updated: {{$e.OutputType}}.AsObject } ))
		.map(lib => {
			request.resolve(lib.prev, lib.updated);
			return protocActions.{{$e.Name}}Success(lib);
		}){{else}}
		.map((resObj: {{$e.OutputType}}.AsObject{{if $e.Repeat}}[]{{end}}) => {
			request.resolve(resObj);
			return protocActions.{{$e.Name}}Success(resObj);
		}){{end}}
		.catch(error => {
			const err: NodeJS.ErrnoException = createErrorObject(error.code, error.message);
			if(request.reject){ request.reject(err); }
			return Observable.of(protocActions.{{$e.Name}}Failure(err));
		})
	})
	.takeUntil(action$.filter(isActionOf(protocActions.{{$e.Name}}Cancel)))
	.repeat();
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
			})){{end}}`

const epicExportTemplate = `export const protocEpics = combineEpics({{range $i, $e := .}}
	{{$e.Name}}Epic,{{end}}
)`

type EpicEntity struct {
	Name           string
	InputType      string
	OutputType     string
	FullMethodName string
	JsonName       string
	Debounce       int64
	Timeout        int64
	Retries        int64
	Repeat         bool
	Auth           string
	AuthFollowup   string
	Host           string
	Updater        bool
	Debug          bool
}

func CreateEpicFile(stateFields []*gp.FieldDescriptorProto, outputType state.OutputTypes, customFields []*gp.FieldDescriptorProto, serviceFiles []*gp.FileDescriptorProto, defaultTimeout int64, defaultRetries int64, authTokenLocation string, hostnameLocation string, hostname string, portin int64, debounce int64, debug bool) (*File, error) {
	epicEntities := []*EpicEntity{}

	// set up port string
	var port string
	if portin != -1 {
		port = ":" + strconv.FormatInt(portin, 10)
	}

	//set up host string
	var host string
	if hostname != "" {
		host = fmt.Sprintf("var host = '%s%s';", hostname, port)
	} else if hostnameLocation != "" {
		host = fmt.Sprintf("var host = store.getState().%s.slice(0, -1) + '%s';", hostnameLocation, port)
	} else {
		return nil, fmt.Errorf("No hostname or hostnameLocation provided. Provide either the hostname or the hostname location in redux so the plugin knows where to send api calls.")
	}

	// transform stateFields into our EpicEntity implementation so template can read values
	for _, field := range stateFields {
		repeated := field.GetLabel() == 3

		// verify the method annotations
		methods, err := GetFieldOptionsString(field, state.E_Method)
		if err != nil {
			return nil, fmt.Errorf("Error getting field level annotations: %v", err)
		}

		// field level overrides for timeout/retry
		timeout, err := GetFieldAnnotationInt(field, state.E_Timeout)
		if err != nil {
			return nil, fmt.Errorf("Error getting field level timeout annotation: %v", err)
		}
		if timeout == -1 { // if it wasn't overriden
			timeout = defaultTimeout
		}
		retries, err := GetFieldAnnotationInt(field, state.E_Retries)
		if err != nil {
			return nil, fmt.Errorf("Error getting field level retries annotation: %v", err)
		}
		if retries == -1 { // if it wasn't overriden
			retries = defaultRetries
		}

		var meth *gp.MethodDescriptorProto
		// get the method for each crud
		for c := CREATE; c < CRUD_MAX; c++ {
			// clear for the loop
			meth = nil

			crudAnnotation := GetAnnotation(methods, c, repeated)
			if crudAnnotation != "" {
				meth, err = FindMethodDescriptor(serviceFiles, crudAnnotation)
				if err != nil {
					return nil, err
				}
			}

			if meth != nil {
				// set up auth string for repeated values (streaming epics)
				var idToken string
				authFollowup := ""
				if authTokenLocation != "" {
					if repeated {
						idToken = fmt.Sprintf("new grpc.Metadata({ 'Authorization': `Bearer ${store.getState().%s}` })", authTokenLocation)
					} else {
						idToken = fmt.Sprintf("var idToken = store.getState().%s;", authTokenLocation)
						authFollowup = "metadata: new grpc.Metadata({ 'Authorization': `Bearer ${idToken}` }),"
					}
				}
				// only returns arrays on these
				var repeatEntity bool
				if CrudName(c, repeated) == "list" {
					repeatEntity = true
				} else {
					repeatEntity = false
				}

				var updater bool
				if CrudName(c, repeated) == "update" && repeated {
					updater = true
				} else {
					updater = false
				}

				epicEntities = append(epicEntities, &EpicEntity{
					Name:           CrudName(c, repeated) + strings.Title(*field.JsonName),
					InputType:      fmt.Sprintf("ProtocTypes.%s", CreatePackageAndTypeString(meth.GetInputType())),
					OutputType:     fmt.Sprintf("ProtocTypes.%s", CreatePackageAndTypeString(meth.GetOutputType())),
					FullMethodName: fmt.Sprintf("ProtocServices.%s", FullMethodNameFormat(crudAnnotation)),
					JsonName:       *field.JsonName,
					Debounce:       debounce,
					Timeout:        timeout,
					Retries:        retries,
					Repeat:         repeatEntity,
					Auth:           idToken,
					AuthFollowup:   authFollowup,
					Host:           host,
					Updater:        updater,
					Debug:          debug,
				})
			}
		}
	}

	// do the same for customActions
	// TODO combine the logic
	for _, field := range customFields {
		repeated := field.GetLabel() == 3

		// verify the method annotations
		methods, err := GetFieldOptionsString(field, state.E_Method)
		if err != nil {
			return nil, fmt.Errorf("Error getting field level annotations: %v", err)
		}

		// field level overrides for timeout/retry
		timeout, err := GetFieldAnnotationInt(field, state.E_Timeout)
		if err != nil {
			return nil, fmt.Errorf("Error getting field level timeout annotation: %v", err)
		}
		if timeout == -1 { // if it wasn't overriden
			timeout = defaultTimeout
		}
		retries, err := GetFieldAnnotationInt(field, state.E_Retries)
		if err != nil {
			return nil, fmt.Errorf("Error getting field level retries annotation: %v", err)
		}
		if retries == -1 { // if it wasn't overriden
			retries = defaultRetries
		}

		var meth *gp.MethodDescriptorProto

		crudAnnotation := methods.GetCustom()
		if crudAnnotation != "" {
			meth, err = FindMethodDescriptor(serviceFiles, crudAnnotation)
			if err != nil {
				return nil, err
			}
		}

		if meth != nil {
			// set up auth string for repeated values (streaming epics)
			var idToken string
			authFollowup := ""
			if authTokenLocation != "" {
				if repeated {
					idToken = fmt.Sprintf("new grpc.Metadata({ 'Authorization': `Bearer ${store.getState().%s` })", authTokenLocation)
				} else {
					idToken = fmt.Sprintf("var idToken = store.getState().%s;", authTokenLocation)
					authFollowup = "metadata: new grpc.Metadata({ 'Authorization': `Bearer ${idToken}` }),"
				}
			}

			// TODO uses repeated from the field name, should use the output type
			epicEntities = append(epicEntities, &EpicEntity{
				Name:           "custom" + strings.Title(*field.JsonName),
				InputType:      fmt.Sprintf("ProtocTypes.%s", CreatePackageAndTypeString(meth.GetInputType())),
				OutputType:     fmt.Sprintf("ProtocTypes.%s", CreatePackageAndTypeString(meth.GetOutputType())),
				FullMethodName: fmt.Sprintf("ProtocServices.%s", FullMethodNameFormat(crudAnnotation)),
				JsonName:       *field.JsonName,
				Debounce:       debounce,
				Timeout:        timeout,
				Retries:        retries,
				Repeat:         repeated,
				Auth:           idToken,
				AuthFollowup:   authFollowup,
				Host:           host,
				Debug:          debug,
			})
		}
	}

	tmpl := template.Must(template.New("epic").Parse(epicTemplate))
	exTmpl := template.Must(template.New("epic-exports").Parse(epicExportTemplate))

	var output bytes.Buffer
	tmpl.Execute(&output, epicEntities)
	exTmpl.Execute(&output, epicEntities)

	return &File{
		Name:    "epics_pb.ts",
		Content: output.String(),
	}, nil
}

func FullMethodNameFormat(name string) string {
	index := strings.LastIndex(name, ".")        // first
	index = strings.LastIndex(name[:index], ".") // second
	return strings.Replace(name, name[:index], strings.Replace(name[:index], ".", "_", -1), 1)
}
