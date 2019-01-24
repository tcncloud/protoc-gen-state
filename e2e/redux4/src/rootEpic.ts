import { combineEpics } from 'redux-observable';

import { LibraryEpics } from './book/epics';
import { TimeoutRetryEpics } from './timeoutRetryTest/epics';
import { protocEpics } from 'protos/BasicState/epics_pb';


export const RootEpic = combineEpics(
  LibraryEpics,
  protocEpics,
  TimeoutRetryEpics,
)
