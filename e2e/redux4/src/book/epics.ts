import { Epic, combineEpics } from 'redux-observable';
import { isActionOf } from 'typesafe-actions';
import { Observable } from 'rxjs';
import 'rxjs/add/observable/dom/ajax';

import { RootState } from '../rootState';
import { RootAction } from "../rootAction";
import * as BookActions from './actions';
import { Library } from './types';


const createLibraryEpic: Epic<RootAction, RootState> = (action$) => action$
  .filter(isActionOf(BookActions.requestCreateLibrary))
  .do((action) => { console.log('right here: ', action) })
  .debounceTime(400)
  .flatMap((action) => {
    return Observable
      .of(action)
      .map((action) => (action.payload as Library))
      .map(lib => BookActions.successCreateLibrary(lib))
      // .takeUntil(isActionOf(BookActions.cancelCreateLibrary))
      .catch((error:string) => Observable.of(BookActions.failureCreateLibrary(error)))
  })

const updateLibraryEpic: Epic<RootAction, RootState> = (action$) => action$
  .filter(isActionOf(BookActions.requestUpdateLibrary))
  .do(() => { console.log('update Library epic') })
  .flatMap((action) => {
    return Observable
      .of(action)
      // .map(action => ({ ...action.payload.updated } as Library))
      .map(action => ({ ...action.payload } as Library))
      .map(lib => BookActions.successUpdateLibrary(lib))
      .catch((error:string) => Observable.of(BookActions.failureUpdateLibrary(error)))
  })

const deleteLibraryEpic: Epic<RootAction, RootState> = (action$) => action$
  .filter(isActionOf(BookActions.requestDeleteLibrary))
  .do(() => { console.log('delete Library epic'); })
  .debounceTime(400)
  .flatMap((action) => {
    return Observable
      .of(action)
      .map((action) => BookActions.successDeleteLibrary({ ...action.payload } as Library))
      .catch((error:string) => Observable.of(BookActions.failureDeleteLibrary(error)))
  })

export const LibraryEpics = combineEpics(
  createLibraryEpic,
  updateLibraryEpic,
  deleteLibraryEpic,
)
