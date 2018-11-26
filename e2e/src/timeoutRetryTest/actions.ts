import { createAction } from 'typesafe-actions';
import * as ProtocTypes from 'proto/BasicState/protoc_types_pb';

export const timeoutRequestPromise = createAction('TIMEOUT_REQUEST_PROMISE', (res) => {
  return (
    book: ProtocTypes.readinglist.Book.AsObject,
    resolve: (payload: ProtocTypes.readinglist.Book.AsObject) => void,
    reject: (error: string) => void,
  ) => res({ book, resolve, reject });
})

export const timeoutSuccess = createAction('TIMEOUT_SUCCESS', (resolve) => {
  return (payload: ProtocTypes.readinglist.Book.AsObject) => resolve(payload);
})

export const timeoutFailure = createAction('TIMEOUT_FAILURE', (resolve) => {
  return (error: string) => resolve(error);
})

export const retryRequestPromise = createAction('RETRY_REQUEST_PROMISE', (res) => {
  return (
    book: ProtocTypes.readinglist.Book.AsObject,
    resolve: (payload: ProtocTypes.readinglist.Book.AsObject) => void,
    reject: (error: string) => void,
  ) => res({ book, resolve, reject });
})

export const retrySuccess = createAction('RETRY_SUCCESS', (resolve) => {
  return (payload: ProtocTypes.readinglist.Book.AsObject) => resolve(payload);
})

export const retryFailure = createAction('RETRY_FAILURE', (resolve) => {
  return (error: string) => resolve(error);
})

export const codeRequestPromise = createAction('CODE_REQUEST_PROMISE', (res) => {
  return (
    book: ProtocTypes.readinglist.Book.AsObject,
    resolve: (payload: ProtocTypes.readinglist.Book.AsObject) => void,
    reject: (error: string) => void,
  ) => res({ book, resolve, reject });
})

export const codeSuccess = createAction('CODE_SUCCESS', (resolve) => {
  return (payload: ProtocTypes.readinglist.Book.AsObject) => resolve(payload);
})

export const codeFailure = createAction('CODE_FAILURE', (resolve) => {
  return (error: string) => resolve(error);
})
