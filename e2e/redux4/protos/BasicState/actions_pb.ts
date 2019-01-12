/* THIS FILE IS GENERATED FROM THE TOOL PROTOC-GEN-STATE */
/* ANYTHING YOU EDIT WILL BE OVERWRITTEN IN FUTURE BUILDS */

import { createAction } from 'typesafe-actions';
import * as ProtocTypes from './protoc_types_pb';


export const getBookOfTheMonthRequest = createAction('PROTOC_GET_BOOKOFTHEMONTH_REQUEST', (res) => {
	return (
		bookOfTheMonth: ProtocTypes.readinglist.Book.AsObject,
		resolve?: (payload: ProtocTypes.readinglist.Book.AsObject) => void,
		reject?: (error: NodeJS.ErrnoException) => void,
	) => res(bookOfTheMonth, { resolve, reject });
});

// deprecated
export const getBookOfTheMonthRequestPromise = getBookOfTheMonthRequest;

export const getBookOfTheMonthSuccess = createAction('PROTOC_GET_BOOKOFTHEMONTH_SUCCESS', (resolve) => {
	return (bookOfTheMonth: ProtocTypes.readinglist.Book.AsObject) => resolve(bookOfTheMonth)
});

export const getBookOfTheMonthFailure = createAction('PROTOC_GET_BOOKOFTHEMONTH_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const getBookOfTheMonthCancel = createAction('PROTOC_GET_BOOKOFTHEMONTH_CANCEL');
export const getTimeoutBookRequest = createAction('PROTOC_GET_TIMEOUTBOOK_REQUEST', (res) => {
	return (
		timeoutBook: ProtocTypes.readinglist.Book.AsObject,
		resolve?: (payload: ProtocTypes.readinglist.Book.AsObject) => void,
		reject?: (error: NodeJS.ErrnoException) => void,
	) => res(timeoutBook, { resolve, reject });
});

// deprecated
export const getTimeoutBookRequestPromise = getTimeoutBookRequest;

export const getTimeoutBookSuccess = createAction('PROTOC_GET_TIMEOUTBOOK_SUCCESS', (resolve) => {
	return (timeoutBook: ProtocTypes.readinglist.Book.AsObject) => resolve(timeoutBook)
});

export const getTimeoutBookFailure = createAction('PROTOC_GET_TIMEOUTBOOK_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const getTimeoutBookCancel = createAction('PROTOC_GET_TIMEOUTBOOK_CANCEL');

export const listLibraryRequest = createAction('PROTOC_LIST_LIBRARY_REQUEST', (res) => {
	return (
		library: ProtocTypes.readinglist.Empty.AsObject,
		resolve?: (payload: ProtocTypes.readinglist.Book.AsObject[]) => void,
		reject?: (error: NodeJS.ErrnoException) => void,
	) => res(library, { resolve, reject });
});

// deprecated
export const listLibraryRequestPromise = listLibraryRequest;

export const listLibrarySuccess = createAction('PROTOC_LIST_LIBRARY_SUCCESS', (resolve) => {
	return (library: ProtocTypes.readinglist.Book.AsObject[]) => resolve(library)
});

export const listLibraryFailure = createAction('PROTOC_LIST_LIBRARY_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const listLibraryCancel = createAction('PROTOC_LIST_LIBRARY_CANCEL');

export const resetLibrary = createAction('PROTOC_RESET_LIBRARY');
export const resetBookOfTheMonth = createAction('PROTOC_RESET_BOOKOFTHEMONTH');
export const resetTimeoutBook = createAction('PROTOC_RESET_TIMEOUTBOOK');

export const createLibraryRequest = createAction('PROTOC_CREATE_LIBRARY_REQUEST', (res) => {
	return (
		library: ProtocTypes.readinglist.Book.AsObject,
		resolve?: (payload: ProtocTypes.readinglist.Book.AsObject[]) => void,
		reject?: (error: NodeJS.ErrnoException) => void,
	) => res(library, { resolve, reject });
});

// deprecated
export const createLibraryRequestPromise = createLibraryRequest;

export const createLibrarySuccess = createAction('PROTOC_CREATE_LIBRARY_SUCCESS', (resolve) => {
	return (library: ProtocTypes.readinglist.Book.AsObject) => resolve(library)
});

export const createLibraryFailure = createAction('PROTOC_CREATE_LIBRARY_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const createLibraryCancel = createAction('PROTOC_CREATE_LIBRARY_CANCEL');
export const createBookOfTheMonthRequest = createAction('PROTOC_CREATE_BOOKOFTHEMONTH_REQUEST', (res) => {
	return (
		bookOfTheMonth: ProtocTypes.readinglist.Book.AsObject,
		resolve?: (payload: ProtocTypes.readinglist.Book.AsObject) => void,
		reject?: (error: NodeJS.ErrnoException) => void,
	) => res(bookOfTheMonth, { resolve, reject });
});

// deprecated
export const createBookOfTheMonthRequestPromise = createBookOfTheMonthRequest;

export const createBookOfTheMonthSuccess = createAction('PROTOC_CREATE_BOOKOFTHEMONTH_SUCCESS', (resolve) => {
	return (bookOfTheMonth: ProtocTypes.readinglist.Book.AsObject) => resolve(bookOfTheMonth)
});

export const createBookOfTheMonthFailure = createAction('PROTOC_CREATE_BOOKOFTHEMONTH_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const createBookOfTheMonthCancel = createAction('PROTOC_CREATE_BOOKOFTHEMONTH_CANCEL');

export const deleteLibraryRequest = createAction('PROTOC_DELETE_LIBRARY_REQUEST', (res) => {
	return (
		library: ProtocTypes.readinglist.Book.AsObject,
		resolve?: (payload: ProtocTypes.readinglist.Book.AsObject[]) => void,
		reject?: (error: NodeJS.ErrnoException) => void,
	) => res(library, { resolve, reject });
});

// deprecated
export const deleteLibraryRequestPromise = deleteLibraryRequest;

export const deleteLibrarySuccess = createAction('PROTOC_DELETE_LIBRARY_SUCCESS', (resolve) => {
	return (library: ProtocTypes.readinglist.Book.AsObject) => resolve(library)
});

export const deleteLibraryFailure = createAction('PROTOC_DELETE_LIBRARY_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const deleteLibraryCancel = createAction('PROTOC_DELETE_LIBRARY_CANCEL');
export const deleteBookOfTheMonthRequest = createAction('PROTOC_DELETE_BOOKOFTHEMONTH_REQUEST', (res) => {
	return (
		bookOfTheMonth: ProtocTypes.readinglist.Book.AsObject,
		resolve?: (payload: ProtocTypes.readinglist.Book.AsObject) => void,
		reject?: (error: NodeJS.ErrnoException) => void,
	) => res(bookOfTheMonth, { resolve, reject });
});

// deprecated
export const deleteBookOfTheMonthRequestPromise = deleteBookOfTheMonthRequest;

export const deleteBookOfTheMonthSuccess = createAction('PROTOC_DELETE_BOOKOFTHEMONTH_SUCCESS', (resolve) => {
	return (bookOfTheMonth: ProtocTypes.readinglist.Book.AsObject) => resolve(bookOfTheMonth)
});

export const deleteBookOfTheMonthFailure = createAction('PROTOC_DELETE_BOOKOFTHEMONTH_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const deleteBookOfTheMonthCancel = createAction('PROTOC_DELETE_BOOKOFTHEMONTH_CANCEL');
export const deleteTimeoutBookRequest = createAction('PROTOC_DELETE_TIMEOUTBOOK_REQUEST', (res) => {
	return (
		timeoutBook: ProtocTypes.readinglist.Book.AsObject,
		resolve?: (payload: ProtocTypes.readinglist.Book.AsObject) => void,
		reject?: (error: NodeJS.ErrnoException) => void,
	) => res(timeoutBook, { resolve, reject });
});

// deprecated
export const deleteTimeoutBookRequestPromise = deleteTimeoutBookRequest;

export const deleteTimeoutBookSuccess = createAction('PROTOC_DELETE_TIMEOUTBOOK_SUCCESS', (resolve) => {
	return (timeoutBook: ProtocTypes.readinglist.Book.AsObject) => resolve(timeoutBook)
});

export const deleteTimeoutBookFailure = createAction('PROTOC_DELETE_TIMEOUTBOOK_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const deleteTimeoutBookCancel = createAction('PROTOC_DELETE_TIMEOUTBOOK_CANCEL');

export const updateLibraryRequest = createAction('PROTOC_UPDATE_LIBRARY_REQUEST', (res) => {
	return (
		prev: ProtocTypes.readinglist.Book.AsObject,
		updated: ProtocTypes.readinglist.Book.AsObject,
		resolve?: (prev: ProtocTypes.readinglist.Book.AsObject, updated: ProtocTypes.readinglist.Book.AsObject) => void,
		reject?: (error: NodeJS.ErrnoException) => void,
	) => res({ prev, updated }, { resolve, reject })
});

export const updateLibraryRequestPromise = updateLibraryRequest


export const updateLibrarySuccess = createAction('PROTOC_UPDATE_LIBRARY_SUCCESS', (resolve) => {
	return (library: { prev: ProtocTypes.readinglist.Book.AsObject, updated: ProtocTypes.readinglist.Book.AsObject }) => resolve(library)
})

export const updateLibraryFailure = createAction('PROTOC_UPDATE_LIBRARY_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const updateLibraryCancel = createAction('PROTOC_UPDATE_LIBRARY_CANCEL');
export const updateBookOfTheMonthRequest = createAction('PROTOC_UPDATE_BOOKOFTHEMONTH_REQUEST', (res) => {
	return (
		bookOfTheMonth: ProtocTypes.readinglist.Book.AsObject,
		resolve?: (payload: ProtocTypes.readinglist.Book.AsObject) => void,
		reject?: (error: NodeJS.ErrnoException) => void,
	) => res(bookOfTheMonth, { resolve, reject })
});

export const updateBookOfTheMonthRequestPromise = updateBookOfTheMonthRequest


export const updateBookOfTheMonthSuccess = createAction('PROTOC_UPDATE_BOOKOFTHEMONTH_SUCCESS', (resolve) => {
	return (bookOfTheMonth: ProtocTypes.readinglist.Book.AsObject) => resolve(bookOfTheMonth)
})

export const updateBookOfTheMonthFailure = createAction('PROTOC_UPDATE_BOOKOFTHEMONTH_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const updateBookOfTheMonthCancel = createAction('PROTOC_UPDATE_BOOKOFTHEMONTH_CANCEL');
export const updateTimeoutBookRequest = createAction('PROTOC_UPDATE_TIMEOUTBOOK_REQUEST', (res) => {
	return (
		timeoutBook: ProtocTypes.readinglist.Book.AsObject,
		resolve?: (payload: ProtocTypes.readinglist.Book.AsObject) => void,
		reject?: (error: NodeJS.ErrnoException) => void,
	) => res(timeoutBook, { resolve, reject })
});

export const updateTimeoutBookRequestPromise = updateTimeoutBookRequest


export const updateTimeoutBookSuccess = createAction('PROTOC_UPDATE_TIMEOUTBOOK_SUCCESS', (resolve) => {
	return (timeoutBook: ProtocTypes.readinglist.Book.AsObject) => resolve(timeoutBook)
})

export const updateTimeoutBookFailure = createAction('PROTOC_UPDATE_TIMEOUTBOOK_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const updateTimeoutBookCancel = createAction('PROTOC_UPDATE_TIMEOUTBOOK_CANCEL');

export const customErrorBookRequest = createAction('PROTOC_CUSTOM_ERRORBOOK_REQUEST', (res) => {
	return (
		errorBook: ProtocTypes.readinglist.Book.AsObject,
		resolve?: (payload: ProtocTypes.readinglist.Book.AsObject) => void,
		reject?: (error: NodeJS.ErrnoException) => void,
	) => res(errorBook, { resolve, reject });
});

//deprecated
export const customErrorBookRequestPromise = customErrorBookRequest;

export const customErrorBookSuccess = createAction('PROTOC_CUSTOM_ERRORBOOK_SUCCESS', (resolve) => {
	return (errorBook: ProtocTypes.readinglist.Book.AsObject) => resolve(errorBook)
});

export const customErrorBookFailure = createAction('PROTOC_CUSTOM_ERRORBOOK_FAILURE', (resolve) => {
	return (error: NodeJS.ErrnoException) => resolve(error)
});

export const customErrorBookCancel = createAction('PROTOC_CUSTOM_ERRORBOOK_CANCEL');
