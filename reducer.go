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
            val, err := CrudNewValue(c, entity, repeated, varName) 
            if err != nil {
              return nil, err
            }
            switchCase = fmt.Sprintf(`%s
      return {
        ...state,
        %s: {
          ...state.%s,
          isLoading: true,
        }
      }`, val, entity.GetJsonName(), entity.GetJsonName())
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

func CrudNewValue(c Crud, entity *gp.FieldDescriptorProto, repeated bool, varName string) (string, error) {
  tsPackageAndType := CreatePackageAndTypeString(entity.GetTypeName())
  payloadName := entity.GetJsonName()
  tsTypeFromState := "ProtocState[\"" + payloadName + "\"][\"value\"]"
  output := ""
  var err error

  switch c {
  case CREATE:
    if repeated {
      output = fmt.Sprintf(`var %s: ProtocTypes.%s.AsObject[] = [...state.%s.value, action.payload] as ProtocTypes.%s.AsObject[];`, varName, tsPackageAndType, payloadName, tsPackageAndType)
    } else {
      output = fmt.Sprintf(`var %s: %s = action.payload as %s;`, varName, tsTypeFromState, tsTypeFromState)
    }
  case GET:
    if repeated {
      output = fmt.Sprintf(`var %s: Protoctypes.%s.AsObject[] = action.payload;`, varName, tsPackageAndType)
    } else {
      output = fmt.Sprintf(`var %s: %s = action.payload;`, varName, tsTypeFromState)
    }
  case UPDATE:
    if repeated {
      output = fmt.Sprintf(`var %s: Protoctypes.%s.AsObject[] = [...state.%s.value] as ProtocTypes.%s.AsObject[];
      var index: number = _.findIndex(%s, action.payload.prev);
      if(index === -1){
        %s.push(action.payload.updated);
      } else {
        %s[index] = action.payload.updated as Protoctypes.%s.AsObject[];
      }`, varName, tsPackageAndType, payloadName, tsPackageAndType, varName, varName, varName, tsPackageAndType)
    } else {
      output = fmt.Sprintf(`var %s: %s = { ...action.payload } as %s;`, varName, tsTypeFromState, tsTypeFromState)
    }
  case DELETE:
    if repeated {
      output = fmt.Sprintf(`var index: number = _.findIndex(state.%s.value, action.payload);
      var %s: ProtocTypes.%s.AsObject[] = [...state.%s.value.slice(0,index), ...state.%s.value.slice(index+1)];`, payloadName, varName, tsPackageAndType, payloadName, payloadName)
    } else {
      output = fmt.Sprintf(`var %s: %s = null;`, varName, tsTypeFromState)
    }
  default:
    err = fmt.Errorf("Invalid CRUD received in CrudNewValue: %v", int(c))
  }

  return output, err
}
