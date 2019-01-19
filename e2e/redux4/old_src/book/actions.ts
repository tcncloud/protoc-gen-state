import { createAction } from 'typesafe-actions';
import { Library } from './types';

// export const cancelCreateLibrary = createAction('BOOK_CANCEL_CREATE_LIBRARY', (resolve) => { return () => resolve()});
export const cancelCreateLibrary = createAction('BOOK_CANCEL_CREATE_LIBRARY');
export const requestCreateLibrary = createAction('BOOK_REQUEST_CREATE_LIBRARY', (resolve) => {
  return (library: Library) => resolve(library);
});
export const successCreateLibrary = createAction('BOOK_SUCCESS_CREATE_LIBRARY', (resolve) => {
  return (library: Library) => resolve(library);
});
export const failureCreateLibrary = createAction('BOOK_FAILURE_CREATE_LIBRARY', (resolve) => {
  return (error: string) => resolve(error);
});


export const requestDeleteLibrary = createAction('BOOK_REQUEST_DELETE_LIBRARY', (resolve) => {
  return (library: Library) => resolve(library);
});
export const successDeleteLibrary = createAction('BOOK_SUCCESS_DELETE_LIBRARY', (resolve) => {
  return (library: Library) => resolve(library);
});
export const failureDeleteLibrary = createAction('BOOK_FAILURE_DELETE_LIBRARY', (resolve) => {
  return (error: string) => resolve(error);
});


export const requestUpdateLibrary = createAction('BOOK_REQUEST_UPDATE_LIBRARY', (resolve) => {
  return (prev: Library, updated: Library) => resolve(prev, updated);
});
export const successUpdateLibrary = createAction('BOOK_SUCCESS_UPDATE_LIBRARY', (resolve) => {
  return (library: Library) => resolve(library);
});
export const failureUpdateLibrary = createAction('BOOK_FAILURE_UPDATE_LIBRARY', (resolve) => {
  return (error: string) => resolve(error);
});
