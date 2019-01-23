import { combineEpics } from 'redux-observable';

// import { DogEpics } from './dog/epics';
// import { RootGraphEpics } from './graph/epics';
// import { LibraryEpics } from './book/epics';
import { TimeoutRetryEpics } from './timeoutRetryTest/epics';
// import { protocEpics } from 'protos/BasicState/epics_pb';


export const RootEpic = combineEpics(
  // DogEpics,
  // RootGraphEpics,
  // LibraryEpics,
  // protocEpics,
  TimeoutRetryEpics,
)
