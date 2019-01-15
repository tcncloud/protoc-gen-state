package redux4

const ActionImportTemplate = `/* THIS FILE IS GENERATED FROM THE TOOL PROTOC-GEN-STATE */
/* ANYTHING YOU EDIT WILL BE OVERWRITTEN IN FUTURE BUILDS */

import { createAction } from 'typesafe-actions';
import * as ProtocTypes from './protoc_types_pb';

`

const ActionCreateTemplate = `{{range $i, $e := .}}
export const create{{$e.JsonName | title}}Request = createAction('PROTOC_CREATE_{{$e.JsonName | caps}}_REQUEST', (res) => {
	return (
		{{$e.JsonName}}: {{$e.InputType}},
		resolve?: (payload: {{$e.OutputType}}{{if $e.Repeat}}[]{{end}}) => void,
		reject?: (error: NodeJS.ErrnoException) => void,
	) => res({{$e.JsonName}}, { resolve, reject });
});

// deprecated
export const create{{$e.JsonName | title}}RequestPromise = create{{$e.JsonName | title}}Request;

export const create{{$e.JsonName | title}}Success = createAction('PROTOC_CREATE_{{$e.JsonName | caps}}_SUCCESS', (resolve) => {
	return ({{$e.JsonName}}: {{$e.OutputType}}) => resolve({{$e.JsonName}})
});

export const create{{$e.JsonName | title}}Failure = createAction('PROTOC_CREATE_{{$e.JsonName | caps}}_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const create{{$e.JsonName | title}}Cancel = createAction('PROTOC_CREATE_{{$e.JsonName | caps}}_CANCEL');{{end}}
`

const ActionUpdateTemplate = `{{range $i, $e := .}}
export const update{{$e.JsonName | title}}Request = createAction('PROTOC_UPDATE_{{$e.JsonName | caps}}_REQUEST', (res) => {
	return ({{if $e.Repeat}}
		prev: {{$e.InputType}},
		updated: {{$e.InputType}},
		resolve?: (prev: {{$e.InputType}}, updated: {{$e.InputType}}) => void,
		reject?: (error: NodeJS.ErrnoException) => void,
	) => res({ prev, updated }, { resolve, reject }){{else}}
		{{$e.JsonName}}: {{$e.InputType}},
		resolve?: (payload: {{$e.OutputType}}) => void,
		reject?: (error: NodeJS.ErrnoException) => void,
	) => res({{$e.JsonName}}, { resolve, reject }){{end}}
});

export const update{{$e.JsonName | title}}RequestPromise = update{{$e.JsonName | title}}Request

{{if $e.Repeat}}
export const update{{$e.JsonName | title}}Success = createAction('PROTOC_UPDATE_{{$e.JsonName | caps}}_SUCCESS', (resolve) => {
	return ({{$e.JsonName}}: { prev: {{$e.InputType}}, updated: {{$e.InputType}} }) => resolve({{$e.JsonName}})
}){{else}}
export const update{{$e.JsonName | title}}Success = createAction('PROTOC_UPDATE_{{$e.JsonName | caps}}_SUCCESS', (resolve) => {
	return ({{$e.JsonName}}: {{$e.OutputType}}) => resolve({{$e.JsonName}})
}){{end}}

export const update{{$e.JsonName | title}}Failure = createAction('PROTOC_UPDATE_{{$e.JsonName | caps}}_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const update{{$e.JsonName | title}}Cancel = createAction('PROTOC_UPDATE_{{$e.JsonName | caps}}_CANCEL');{{end}}
`

const ActionGetTemplate = `{{range $i, $e := .}}
export const get{{$e.JsonName | title}}Request = createAction('PROTOC_GET_{{$e.JsonName | caps}}_REQUEST', (res) => {
	return (
		{{$e.JsonName}}: {{$e.InputType}},
		resolve?: (payload: {{$e.OutputType}}) => void,
		reject?: (error: NodeJS.ErrnoException) => void,
	) => res({{$e.JsonName}}, { resolve, reject });
});

// deprecated
export const get{{$e.JsonName | title}}RequestPromise = get{{$e.JsonName | title}}Request;

export const get{{$e.JsonName | title}}Success = createAction('PROTOC_GET_{{$e.JsonName | caps}}_SUCCESS', (resolve) => {
	return ({{$e.JsonName}}: {{$e.OutputType}}) => resolve({{$e.JsonName}})
});

export const get{{$e.JsonName | title}}Failure = createAction('PROTOC_GET_{{$e.JsonName | caps}}_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const get{{$e.JsonName | title}}Cancel = createAction('PROTOC_GET_{{$e.JsonName | caps}}_CANCEL');{{end}}
`

const ActionListTemplate = `{{range $i, $e := .}}
export const list{{$e.JsonName | title}}Request = createAction('PROTOC_LIST_{{$e.JsonName | caps}}_REQUEST', (res) => {
	return (
		{{$e.JsonName}}: {{$e.InputType}},
		resolve?: (payload: {{$e.OutputType}}[]) => void,
		reject?: (error: NodeJS.ErrnoException) => void,
	) => res({{$e.JsonName}}, { resolve, reject });
});

// deprecated
export const list{{$e.JsonName | title}}RequestPromise = list{{$e.JsonName | title}}Request;

export const list{{$e.JsonName | title}}Success = createAction('PROTOC_LIST_{{$e.JsonName | caps}}_SUCCESS', (resolve) => {
	return ({{$e.JsonName}}: {{$e.OutputType}}[]) => resolve({{$e.JsonName}})
});

export const list{{$e.JsonName | title}}Failure = createAction('PROTOC_LIST_{{$e.JsonName | caps}}_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const list{{$e.JsonName | title}}Cancel = createAction('PROTOC_LIST_{{$e.JsonName | caps}}_CANCEL');{{end}}
`

const ActionDeleteTemplate = `{{range $i, $e := .}}
export const delete{{$e.JsonName | title}}Request = createAction('PROTOC_DELETE_{{$e.JsonName | caps}}_REQUEST', (res) => {
	return (
		{{$e.JsonName}}: {{$e.InputType}},
		resolve?: (payload: {{$e.OutputType}}{{if $e.Repeat}}[]{{end}}) => void,
		reject?: (error: NodeJS.ErrnoException) => void,
	) => res({{$e.JsonName}}, { resolve, reject });
});

// deprecated
export const delete{{$e.JsonName | title}}RequestPromise = delete{{$e.JsonName | title}}Request;

export const delete{{$e.JsonName | title}}Success = createAction('PROTOC_DELETE_{{$e.JsonName | caps}}_SUCCESS', (resolve) => {
	return ({{$e.JsonName}}: {{$e.OutputType}}) => resolve({{$e.JsonName}})
});

export const delete{{$e.JsonName | title}}Failure = createAction('PROTOC_DELETE_{{$e.JsonName | caps}}_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const delete{{$e.JsonName | title}}Cancel = createAction('PROTOC_DELETE_{{$e.JsonName | caps}}_CANCEL');{{end}}
`

const ActionCustomTemplate = `{{range $i, $e := .}}
export const custom{{$e.JsonName | title}}Request = createAction('PROTOC_CUSTOM_{{$e.JsonName | caps}}_REQUEST', (res) => {
	return (
		{{$e.JsonName}}: {{$e.InputType}},
		resolve?: (payload: {{$e.OutputType}}{{if $e.Repeat}}[]{{end}}) => void,
		reject?: (error: NodeJS.ErrnoException) => void,
	) => res({{$e.JsonName}}, { resolve, reject });
});

//deprecated
export const custom{{$e.JsonName | title}}RequestPromise = custom{{$e.JsonName | title}}Request;

export const custom{{$e.JsonName | title}}Success = createAction('PROTOC_CUSTOM_{{$e.JsonName | caps}}_SUCCESS', (resolve) => {
	return ({{$e.JsonName}}: {{$e.OutputType}}{{if $e.Repeat}}[]{{end}}) => resolve({{$e.JsonName}})
});

export const custom{{$e.JsonName | title}}Failure = createAction('PROTOC_CUSTOM_{{$e.JsonName | caps}}_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const custom{{$e.JsonName | title}}Cancel = createAction('PROTOC_CUSTOM_{{$e.JsonName | caps}}_CANCEL');{{end}}
`

const ActionResetTemplate = `{{range $i, $e := .}}
export const reset{{$e.JsonName | title}} = createAction('PROTOC_RESET_{{$e.JsonName | caps}}');{{end}}
`
