import { combineEpics, ActionsObservable } from 'redux-observable';
import { isActionOf, ActionType } from 'typesafe-actions';
import { of } from 'rxjs';
import { filter, map, flatMap, debounceTime, tap, catchError } from 'rxjs/operators';

import * as BookActions from './actions';
import { Library } from './types';

type bookActionsType = ActionType<typeof BookActions>

const createLibraryEpic = (action$: ActionsObservable<bookActionsType>) => action$.pipe(
  filter(isActionOf(BookActions.requestCreateLibrary))
  ,tap((action) => { console.log('right here: ', action) })
  ,debounceTime(400)
  ,flatMap((action:any) => (
    of(action).pipe(
      map((action) => (action.payload as Library))
      ,map(lib => BookActions.successCreateLibrary(lib))
      // .takeUntil(isActionOf(BookActions.cancelCreateLibrary))
      ,catchError((error:string) => of(BookActions.failureCreateLibrary(error)))
    )
  ))
)

const updateLibraryEpic = (action$: ActionsObservable<bookActionsType>) => action$.pipe(
  filter(isActionOf(BookActions.requestUpdateLibrary))
  ,tap(() => { console.log('update Library epic') })
  ,flatMap((action:any) => (
    of(action).pipe(
      // .map(action => ({ ...action.payload.updated } as Library))
      map(action => ({ ...action.payload } as Library))
      ,map(lib => BookActions.successUpdateLibrary(lib))
      ,catchError((error:string) => of(BookActions.failureUpdateLibrary(error)))
    )
  ))
)

const deleteLibraryEpic = (action$: ActionsObservable<bookActionsType>) => action$.pipe(
  filter(isActionOf(BookActions.requestDeleteLibrary))
  ,tap(() => { console.log('delete Library epic'); })
  ,debounceTime(400)
  ,flatMap((action:any) => (
    of(action).pipe(
      map((action) => BookActions.successDeleteLibrary({ ...action.payload } as Library))
      ,catchError((error:string) => of(BookActions.failureDeleteLibrary(error)))
    )
  ))
)

export const LibraryEpics = combineEpics(
  createLibraryEpic,
  updateLibraryEpic,
  deleteLibraryEpic,
)
