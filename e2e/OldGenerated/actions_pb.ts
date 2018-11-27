/* THIS FILE IS GENERATED FROM THE TOOL PROTOC-GEN-STATE  */
/* ANYTHING YOU EDIT WILL BE OVERWRITTEN IN FUTURE BUILDS */
import { createAction } from 'typesafe-actions';
import * as ProtocTypes from './protoc_types_pb';


export const createLibraryRequest = createAction('PROTOC_CREATE_LIBRARY_REQUEST', (resolve) => {
	return (library: ProtocTypes.readinglist.Book.AsObject) => resolve(library)
});
export const createLibraryRequestPromise = createAction('PROTOC_CREATE_LIBRARY_REQUEST_PROMISE', (res) => {
	return (
		library: ProtocTypes.readinglist.Book.AsObject,
		resolve: (payload: ProtocTypes.readinglist.Book.AsObject[]) => void,
		reject: (error: NodeJS.ErrnoException) => void,
	) => res({ library, resolve, reject });
});
export const createLibrarySuccess = createAction('PROTOC_CREATE_LIBRARY_SUCCESS', (resolve) => {
	return (library: ProtocTypes.readinglist.Book.AsObject) => resolve(library)
});
export const createLibraryFailure = createAction('PROTOC_CREATE_LIBRARY_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});
export const createLibraryCancel = createAction('PROTOC_CREATE_LIBRARY_CANCEL');
export const listLibraryRequest = createAction('PROTOC_LIST_LIBRARY_REQUEST', (resolve) => {
	return (library: ProtocTypes.readinglist.Empty.AsObject) => resolve(library)
});
export const listLibraryRequestPromise = createAction('PROTOC_LIST_LIBRARY_REQUEST_PROMISE', (res) => {
	return (
		library: ProtocTypes.readinglist.Empty.AsObject,
		resolve: (payload: ProtocTypes.readinglist.Book.AsObject[]) => void,
		reject: (error: NodeJS.ErrnoException) => void,
	) => res({ library, resolve, reject });
});
export const listLibrarySuccess = createAction('PROTOC_LIST_LIBRARY_SUCCESS', (resolve) => {
	return (library: ProtocTypes.readinglist.Book.AsObject[]) => resolve(library)
});
export const listLibraryFailure = createAction('PROTOC_LIST_LIBRARY_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});
export const listLibraryCancel = createAction('PROTOC_LIST_LIBRARY_CANCEL');
export const updateLibraryRequest = createAction('PROTOC_UPDATE_LIBRARY_REQUEST', (resolve) => {
	return (prev: ProtocTypes.readinglist.Book.AsObject, updated: ProtocTypes.readinglist.Book.AsObject ) => resolve({prev, updated})
});
export const updateLibraryRequestPromise = createAction('PROTOC_UPDATE_LIBRARY_REQUEST_PROMISE', (res) => {
	return (
		prev: ProtocTypes.readinglist.Book.AsObject,
		updated: ProtocTypes.readinglist.Book.AsObject,
		resolve: (prev: ProtocTypes.readinglist.Book.AsObject, updated: ProtocTypes.readinglist.Book.AsObject) => void,
		reject: (error: NodeJS.ErrnoException) => void,
	) => res({ prev, updated, resolve, reject })
});
export const updateLibrarySuccess = createAction('PROTOC_UPDATE_LIBRARY_SUCCESS', (resolve) => {
	return (library: {prev: ProtocTypes.readinglist.Book.AsObject, updated: ProtocTypes.readinglist.Book.AsObject}) => resolve(library)
});
export const updateLibraryFailure = createAction('PROTOC_UPDATE_LIBRARY_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});
export const updateLibraryCancel = createAction('PROTOC_UPDATE_LIBRARY_CANCEL');
export const deleteLibraryRequest = createAction('PROTOC_DELETE_LIBRARY_REQUEST', (resolve) => {
	return (library: ProtocTypes.readinglist.Book.AsObject) => resolve(library)
});
export const deleteLibraryRequestPromise = createAction('PROTOC_DELETE_LIBRARY_REQUEST_PROMISE', (res) => {
	return (
		library: ProtocTypes.readinglist.Book.AsObject,
		resolve: (payload: ProtocTypes.readinglist.Book.AsObject[]) => void,
		reject: (error: NodeJS.ErrnoException) => void,
	) => res({ library, resolve, reject });
});
export const deleteLibrarySuccess = createAction('PROTOC_DELETE_LIBRARY_SUCCESS', (resolve) => {
	return (library: ProtocTypes.readinglist.Book.AsObject) => resolve(library)
});
export const deleteLibraryFailure = createAction('PROTOC_DELETE_LIBRARY_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});
export const deleteLibraryCancel = createAction('PROTOC_DELETE_LIBRARY_CANCEL');
export const resetLibrary = createAction('PROTOC_RESET_LIBRARY');
export const createBookOfTheMonthRequest = createAction('PROTOC_CREATE_BOOKOFTHEMONTH_REQUEST', (resolve) => {
	return (bookOfTheMonth: ProtocTypes.readinglist.Book.AsObject) => resolve(bookOfTheMonth)
});
export const createBookOfTheMonthRequestPromise = createAction('PROTOC_CREATE_BOOKOFTHEMONTH_REQUEST_PROMISE', (res) => {
	return (
		bookOfTheMonth: ProtocTypes.readinglist.Book.AsObject,
		resolve: (payload: ProtocTypes.readinglist.Book.AsObject) => void,
		reject: (error: NodeJS.ErrnoException) => void,
	) => res({ bookOfTheMonth, resolve, reject });
});
export const createBookOfTheMonthSuccess = createAction('PROTOC_CREATE_BOOKOFTHEMONTH_SUCCESS', (resolve) => {
	return (bookOfTheMonth: ProtocTypes.readinglist.Book.AsObject) => resolve(bookOfTheMonth)
});
export const createBookOfTheMonthFailure = createAction('PROTOC_CREATE_BOOKOFTHEMONTH_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});
export const createBookOfTheMonthCancel = createAction('PROTOC_CREATE_BOOKOFTHEMONTH_CANCEL');
export const getBookOfTheMonthRequest = createAction('PROTOC_GET_BOOKOFTHEMONTH_REQUEST', (resolve) => {
	return (bookOfTheMonth: ProtocTypes.readinglist.Book.AsObject) => resolve(bookOfTheMonth)
});
export const getBookOfTheMonthRequestPromise = createAction('PROTOC_GET_BOOKOFTHEMONTH_REQUEST_PROMISE', (res) => {
	return (
		bookOfTheMonth: ProtocTypes.readinglist.Book.AsObject,
		resolve: (payload: ProtocTypes.readinglist.Book.AsObject) => void,
		reject: (error: NodeJS.ErrnoException) => void,
	) => res({ bookOfTheMonth, resolve, reject });
});
export const getBookOfTheMonthSuccess = createAction('PROTOC_GET_BOOKOFTHEMONTH_SUCCESS', (resolve) => {
	return (bookOfTheMonth: ProtocTypes.readinglist.Book.AsObject) => resolve(bookOfTheMonth)
});
export const getBookOfTheMonthFailure = createAction('PROTOC_GET_BOOKOFTHEMONTH_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});
export const getBookOfTheMonthCancel = createAction('PROTOC_GET_BOOKOFTHEMONTH_CANCEL');
export const updateBookOfTheMonthRequest = createAction('PROTOC_UPDATE_BOOKOFTHEMONTH_REQUEST', (resolve) => {
	return (bookOfTheMonth: ProtocTypes.readinglist.Book.AsObject) => resolve(bookOfTheMonth)
});
export const updateBookOfTheMonthRequestPromise = createAction('PROTOC_UPDATE_BOOKOFTHEMONTH_REQUEST_PROMISE', (res) => {
	return (
		bookOfTheMonth: ProtocTypes.readinglist.Book.AsObject,
		resolve: (payload: ProtocTypes.readinglist.Book.AsObject) => void,
		reject: (error: NodeJS.ErrnoException) => void,
	) => res({ bookOfTheMonth, resolve, reject });
});
export const updateBookOfTheMonthSuccess = createAction('PROTOC_UPDATE_BOOKOFTHEMONTH_SUCCESS', (resolve) => {
	return (bookOfTheMonth: ProtocTypes.readinglist.Book.AsObject) => resolve(bookOfTheMonth)
});
export const updateBookOfTheMonthFailure = createAction('PROTOC_UPDATE_BOOKOFTHEMONTH_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});
export const updateBookOfTheMonthCancel = createAction('PROTOC_UPDATE_BOOKOFTHEMONTH_CANCEL');
export const deleteBookOfTheMonthRequest = createAction('PROTOC_DELETE_BOOKOFTHEMONTH_REQUEST', (resolve) => {
	return (bookOfTheMonth: ProtocTypes.readinglist.Book.AsObject) => resolve(bookOfTheMonth)
});
export const deleteBookOfTheMonthRequestPromise = createAction('PROTOC_DELETE_BOOKOFTHEMONTH_REQUEST_PROMISE', (res) => {
	return (
		bookOfTheMonth: ProtocTypes.readinglist.Book.AsObject,
		resolve: (payload: ProtocTypes.readinglist.Book.AsObject) => void,
		reject: (error: NodeJS.ErrnoException) => void,
	) => res({ bookOfTheMonth, resolve, reject });
});
export const deleteBookOfTheMonthSuccess = createAction('PROTOC_DELETE_BOOKOFTHEMONTH_SUCCESS', (resolve) => {
	return (bookOfTheMonth: ProtocTypes.readinglist.Book.AsObject) => resolve(bookOfTheMonth)
});
export const deleteBookOfTheMonthFailure = createAction('PROTOC_DELETE_BOOKOFTHEMONTH_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});
export const deleteBookOfTheMonthCancel = createAction('PROTOC_DELETE_BOOKOFTHEMONTH_CANCEL');
export const resetBookOfTheMonth = createAction('PROTOC_RESET_BOOKOFTHEMONTH');
export const getTimeoutBookRequest = createAction('PROTOC_GET_TIMEOUTBOOK_REQUEST', (resolve) => {
	return (timeoutBook: ProtocTypes.readinglist.Book.AsObject) => resolve(timeoutBook)
});
export const getTimeoutBookRequestPromise = createAction('PROTOC_GET_TIMEOUTBOOK_REQUEST_PROMISE', (res) => {
	return (
		timeoutBook: ProtocTypes.readinglist.Book.AsObject,
		resolve: (payload: ProtocTypes.readinglist.Book.AsObject) => void,
		reject: (error: NodeJS.ErrnoException) => void,
	) => res({ timeoutBook, resolve, reject });
});
export const getTimeoutBookSuccess = createAction('PROTOC_GET_TIMEOUTBOOK_SUCCESS', (resolve) => {
	return (timeoutBook: ProtocTypes.readinglist.Book.AsObject) => resolve(timeoutBook)
});
export const getTimeoutBookFailure = createAction('PROTOC_GET_TIMEOUTBOOK_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});
export const getTimeoutBookCancel = createAction('PROTOC_GET_TIMEOUTBOOK_CANCEL');
export const resetTimeoutBook = createAction('PROTOC_RESET_TIMEOUTBOOK');
export const customErrorBookRequest = createAction('PROTOC_CUSTOM_ERRORBOOK_REQUEST', (resolve) => {
	return (errorBook: ProtocTypes.readinglist.Book.AsObject) => resolve(errorBook)
});
export const customErrorBookRequestPromise = createAction('PROTOC_CUSTOM_ERRORBOOK_REQUEST_PROMISE', (res) => {
	return (
		errorBook: ProtocTypes.readinglist.Book.AsObject,
		resolve: (payload: ProtocTypes.readinglist.Book.AsObject) => void,
		reject: (error: NodeJS.ErrnoException) => void,
	) => res({ errorBook, resolve, reject });
});
export const customErrorBookSuccess = createAction('PROTOC_CUSTOM_ERRORBOOK_SUCCESS', (resolve) => {
	return (errorBook: ProtocTypes.readinglist.Book.AsObject) => resolve(errorBook)
});
export const customErrorBookFailure = createAction('PROTOC_CUSTOM_ERRORBOOK_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});
export const customErrorBookCancel = createAction('PROTOC_CUSTOM_ERRORBOOK_CANCEL');
