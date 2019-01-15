package redux3

const StateTemplate = `/* THIS FILE IS GENERATED FROM THE TOOL PROTOC-GEN-STATE  */
/* ANYTHING YOU EDIT WILL BE OVERWRITTEN IN FUTURE BUILDS */

import * as ProtocTypes from './protoc_types_pb';

export interface ProtocState { {{range $i, $entity := .}}
{{$entity.FieldName}}: {
  isLoading: boolean;
  error: { code: string; message: string; };
  {{if $entity.Repeated}}value: ProtocTypes.{{$entity.FullTypeName}}.AsObject[];
  {{else}}value: ProtocTypes.{{$entity.FullTypeName}}.AsObject | null;{{end}}
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

