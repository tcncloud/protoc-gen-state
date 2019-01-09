import { DogState, InitialDogState } from './dog/reducer';
import { GraphState, InitialGraphState } from './graph/reducers';
import { BookState, InitialBookState } from './book/reducers';
import { ProtocState, initialProtocState } from 'proto/BasicState/state_pb';
import { ConfigState, InitialConfigState } from './config/reducers';

export const InitialState : RootState = {
  dog: InitialDogState,
  graph: InitialGraphState,
  book: InitialBookState,
  protoc: initialProtocState,
  config: InitialConfigState,
};

export type RootState = {
  dog: DogState;
  graph: GraphState;
  book: BookState;
  protoc: ProtocState;
  config: ConfigState;
};
