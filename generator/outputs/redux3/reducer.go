package redux3

const ReducerTemplate = `/* THIS FILE IS GENERATED FROM THE TOOL PROTOC-GEN-STATE  */
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
