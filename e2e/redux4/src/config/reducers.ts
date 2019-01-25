import { getType } from 'typesafe-actions';
import { SetHost } from './actions';

import { RootAction } from '../rootAction';


export interface ConfigState {
  host: string;
}

export const InitialConfigState: ConfigState = {
  host: 'http://localhost/',
  token: 'thisisasecrettoken',
};

export function ConfigReducer(state: ConfigState = InitialConfigState, action: RootAction) {
  switch(action.type) {
    case getType(SetHost):
      return {
        ...state,
        host: action.payload,
      }

    default: return state;
  }
}
