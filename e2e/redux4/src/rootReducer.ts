import { combineReducers, Reducer, } from 'redux'

import { RootState } from './rootState';
import { BookReducer } from './book/reducers';
import { ConfigReducer } from './config/reducers';
import { protocReducer } from 'protos/BasicState/reducer_pb';


export const RootReducer: Reducer<RootState> = combineReducers({
  book: BookReducer,
  protoc: protocReducer,
  config: ConfigReducer,
})
