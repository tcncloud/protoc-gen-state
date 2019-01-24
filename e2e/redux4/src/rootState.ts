import { BookState, InitialBookState } from './book/reducers';
import { ProtocState, initialProtocState } from 'protos/BasicState/state_pb';
import { ConfigState, InitialConfigState } from './config/reducers';

export const InitialState : RootState = {
  book: InitialBookState,
  protoc: initialProtocState,
  config: InitialConfigState,
};

export type RootState = {
  book: BookState;
  protoc: ProtocState;
  config: ConfigState;
};
