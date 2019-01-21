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
	"strings"
	"text/template"

	gp "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

// cludg also has reset
// Should try out subtemplates
// TODO make sure maps are supported

type ReducerEntity struct {
	SwitchCase      string
	Name            string
	CludgEffectName string
}

func (this *GenericOutputter) CreateReducerFile(stateFields []*gp.FieldDescriptorProto, debug bool) (*File, error) {
	reducerEntities := []*ReducerEntity{}

	for _, entity := range stateFields {
		fieldAnnotations, err := GetFieldOptions(entity)
		if err != nil {
			return nil, err
		}

		for c := CREATE; c < CRUD_MAX; c++ {
			switchCase := ""
			repeated := entity.GetLabel() == 3
			annotation := GetAnnotation(*fieldAnnotations.GetMethod(), c, repeated)

			if annotation != "" {
				for s := REQUEST; s < SIDE_EFFECT_MAX; s++ {
					switch s {
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
          isLoading: false,
          value: %s,
          error: initialProtocState.%s.error
        }
      }`, val, entity.GetJsonName(), entity.GetJsonName(), varName, entity.GetJsonName())
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
						SwitchCase:      switchCase,
						CludgEffectName: CrudName(c, repeated) + strings.Title(entity.GetJsonName()) + strings.Title(SideEffectName(s)),
						Name:            entity.GetJsonName(),
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
			Name:            entity.GetJsonName(),
		})
	}

	funcMap := template.FuncMap{
		"title": strings.Title,
	}

	tmpl := template.Must(template.New("reducer").Funcs(funcMap).Parse(this.ReducerFile.Template))

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
			output = fmt.Sprintf(`var %s: ProtocTypes.%s.AsObject[] = action.payload;`, varName, tsPackageAndType)
		} else {
			output = fmt.Sprintf(`var %s: %s = action.payload;`, varName, tsTypeFromState)
		}
	case UPDATE:
		if repeated {
			output = fmt.Sprintf(`var %s: ProtocTypes.%s.AsObject[] = [...state.%s.value] as ProtocTypes.%s.AsObject[];
      var index: number = _.findIndex(%s, action.payload.prev);
      if(index === -1){
        %s.push(action.payload.updated);
      } else {
        %s[index] = action.payload.updated as ProtocTypes.%s.AsObject;
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
