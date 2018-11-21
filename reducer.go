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


const reducerTemplate = `/* THIS FILE IS GENERATED FROM THE TOOL PROTOC-GEN-STATE  */
/* ANYTHING YOU EDIT WILL BE OVERWRITTEN IN FUTURE BUILDS */

import { getType, ActionType } from 'typesafe-actions';
import _ from 'lodash';
import * as protocActions from './actions_pb';
import * as ProtocTypes from './protoc_types_pb';
import { ProtocState, initialProtocState } from './state_pb';

type RootAction = ActionType<typeof protocActions>;

export function protocReducer(state: ProtocState = initialProtocState, action: RootAction) {
  switch(action.type) { {{range $i, $entity := .}}
    case getType(protocActions['{{$entity.CludgEffectName}}']):
      {{$entity.SwitchCase}}{{end}}
    default: return state;
  }
};
`

type ReducerEntity struct {
  SwitchCase string
  Name string
  CludgEffectName string
}

func CreateReducerFile(stateFields []*gp.FieldDescriptorProto) (*File, error) {
  reducerEntities := []*ReducerEntity{}

  for _, entity := range stateFields {
    methods, err := GetFieldOptionsString(entity, state.E_Method)
    if err != nil {
      return nil, err
    }

    for c := CREATE; c < CRUD_MAX; c++ {
      switchCase := ""
      repeated := entity.GetLabel() == 3
      annotation := GetAnnotation(methods, c, repeated)

      if annotation != "" {
        for s := REQUEST; s < SIDE_EFFECT_MAX; s++ {
          switch(s){
          case REQUEST:
            switchCase = fmt.Sprintf(`return {
        ...state,
        %s: {
          ...state.%s,
          isLoading: true,
        }
      }`, entity.GetJsonName(), entity.GetJsonName())
          case SUCCESS:
            varName := "new" + strings.Title(entity.GetJsonName()) + "Value"
            if repeated {
              varName += "Array"
            }
            switchCase += CrudNewValue(c, entity, repeated, varName) 
            switchCase = fmt.Sprintf(`return {
        ...state,
        %s: {
          ...state.%s,
          isLoading: true,
        }
      }`, entity.GetJsonName(), entity.GetJsonName())
          case FAILURE:
            switchCase = fmt.Sprintf(`return {
        ...state,
        %s: {
          ...state.%s,
          isLoading: false,
          error: { code: action.payload.code, message: action.payload.message },
        }
      }`, entity.GetJsonName(), entity.GetJsonName())
          case CANCEL:
            switchCase = fmt.Sprintf(`return {
        ...state,
        %s: {
          ...state.%s,
          isLoading: false,
        }
      }`, entity.GetJsonName(), entity.GetJsonName())
          default:
            return nil, fmt.Errorf("Unsupported Side Effect: %v", s)
          }

          reducerEntities = append(reducerEntities, &ReducerEntity{
            SwitchCase: switchCase,
            CludgEffectName: CrudName(c, repeated) + strings.Title(entity.GetJsonName()) + strings.Title(SideEffectName(s)),
            Name: entity.GetJsonName(),
          })
        }
      }
    }

    // create reset reducer block
    reducerEntities = append(reducerEntities, &ReducerEntity{
      SwitchCase: fmt.Sprintf(`return {
        ...state,
        %s: initialProtocState.%s
      }`, entity.GetJsonName(), entity.GetJsonName()),
      CludgEffectName: "reset" + strings.Title(entity.GetJsonName()),
      Name: entity.GetJsonName(),
    })
  }

  funcMap := template.FuncMap{
    "title": strings.Title,
  }

  tmpl := template.Must(template.New("reducer").Funcs(funcMap).Parse(reducerTemplate))

  var output bytes.Buffer
  tmpl.Execute(&output, reducerEntities)

	return &File{
		Name:    "reducer_pb.ts",
		Content: output.String(),
	}, nil
}

func CrudNewValue(c Crud, entity *gp.FieldDescriptorProto, repeated bool, varName string) string {
  payloadName := entity.GetJsonName()
  tsType := entity.MessageType().GetJsonName()
  fmt.Println("tsType: ", tsType, "\n\n\nn\n\n\n\n\nn\n")
  tsTypePackage := strings.Replace("ts.type.package", ".", "_", -1)
  // TODO find these message type stuff below
  // tsType := entity.message_type().name() + ".AsObject"
  // tsTypePackage := replacePeriodsWithUnderscore(entity.message_type().file().package()
  return tsTypePackage
}
