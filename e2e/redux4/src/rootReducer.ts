import { combineReducers, Reducer, } from 'redux'

import { RootState } from './rootState';
// import { DogReducer } from './dog/reducer';
// import { GraphReducer } from './graph/reducers';
// import { BookReducer } from './book/reducers';
import { ConfigReducer } from './config/reducers';
import { protocReducer } from 'protos/BasicState/reducer_pb';


export const RootReducer: Reducer<RootState> = combineReducers({
  // dog: DogReducer,
  // graph: GraphReducer,
  // book: BookReducer,
  protoc: protocReducer,
  config: ConfigReducer,
})
