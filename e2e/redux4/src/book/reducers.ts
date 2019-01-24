import { getType } from 'typesafe-actions';

import { RootAction } from '../rootAction';
import { Library } from './types';
import * as BookActions from './actions';


export interface BookState {
  library: Library;
  isLoading: boolean;
  error: string | null;

}

export const InitialBookState: BookState = {
  library: { name: 'compendium', city: 'boston' },
  isLoading: false,
  error: null,
};

export function BookReducer(state: BookState = InitialBookState, action: RootAction) {
  switch(action.type) {
    case getType(BookActions.cancelCreateLibrary):
      return {
        ...state,
        isLoading: false,
      }
    case getType(BookActions.requestCreateLibrary):
      return {
        ...state,
        isLoading: true,
      }
    case getType(BookActions.successCreateLibrary):
      return {
        ...state,
        library: action.payload,
        isLoading: false,
        error: InitialBookState.error
      }
    case getType(BookActions.failureCreateLibrary):
      return {
        ...state,
        isLoading: false,
        error: action.payload
      }


    case getType(BookActions.requestDeleteLibrary):
      return {
        ...state,
        isLoading: true,
      }
    case getType(BookActions.successDeleteLibrary):
      return {
        ...state,
        isLoading: false,
        library: action.payload,
        error: InitialBookState.error
      }
    case getType(BookActions.failureDeleteLibrary):
      return {
        ...state,
        isLoading: false,
        error: action.payload
      }


    case getType(BookActions.requestUpdateLibrary):
      return {
        ...state,
        isLoading: true,
      }
    case getType(BookActions.successUpdateLibrary):
      return {
        ...state,
        isLoading: false,
        library: action.payload,
        error: InitialBookState.error
      }
    case getType(BookActions.failureUpdateLibrary):
      return {
        ...state,
        isLoading: false,
        error: action.payload
      }

    default: return state;
  }
}
