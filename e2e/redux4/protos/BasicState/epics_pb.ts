/* THIS FILE IS GENERATED FROM THE TOOL PROTOC-GEN-STATE  */
/* ANYTHING YOU EDIT WILL BE OVERWRITTEN IN FUTURE BUILDS */

import { combineEpics } from 'redux-observable';
import { isActionOf } from 'typesafe-actions';
import { Observable } from 'rxjs';
import _ from 'lodash';
import { grpc } from 'grpc-web-client';
import { UnaryOutput } from 'grpc-web-client/dist/unary';
import 'rxjs/add/observable/dom/ajax';
import { toMessage } from './to_message_pb';
import * as protocActions from './actions_pb';
import * as ProtocTypes from './protoc_types_pb';
import * as ProtocServices from './protoc_services_pb';


function noop() {
	return;
}

function createErrorObject(code: number|string|undefined, message: string): NodeJS.ErrnoException {
	var err: NodeJS.ErrnoException = new Error();
	err.message = message;
	if(code && typeof code == 'number') { err.code = code.toString(); }
	if(code && typeof code == 'string') { err.code = code; }
	return err;
}


export const createLibraryEpic = (action$, store) => action$
	.filter(isActionOf(protocActions.createLibraryRequest))
	.debounceTime(510)
	.map(({ payload, meta: { resolve = noop, reject = noop } }) => ({
		message: toMessage(payload, ProtocTypes.readinglist.Book),
		resolve,
		reject,
	}))
	.flatMap((request) => {
    return Observable
		.defer(() => new Promise((resolve, reject) => {
      
			var host = store.getState().config.host.slice(0, -1) + ':9090';
			
			grpc.unary(ProtocServices.readinglist.ReadingList.CreateBook, {
				request: request.message,
				host: host,
				
				onEnd: (res: UnaryOutput<ProtocTypes.readinglist.Book>) => {
          
					if(res.status != grpc.Code.OK){
            
						const err: NodeJS.ErrnoException = createErrorObject(res.status, res.statusMessage);
						reject(err);
					}
					if(res.message){
						resolve(res.message.toObject());
					}
				}
			});
		})) 
		.retry(0)
		.timeout(3000)
		.map((resObj: ProtocTypes.readinglist.Book.AsObject) => {
			request.resolve(resObj);
			return protocActions.createLibrarySuccess(resObj);
		})
		.catch(error => {
			const err: NodeJS.ErrnoException = createErrorObject(error.code, error.message);
			if(request.reject){ request.reject(err); }
			return Observable.of(protocActions.createLibraryFailure(err));
		})
	})
	.takeUntil(action$.filter(isActionOf(protocActions.createLibraryCancel)))
	.repeat();

export const listLibraryEpic = (action$, store) => action$
	.filter(isActionOf(protocActions.listLibraryRequest))
	.debounceTime(510)
	.map(({ payload, meta: { resolve = noop, reject = noop } }) => ({
		message: toMessage(payload, ProtocTypes.readinglist.Empty),
		resolve,
		reject,
	}))
	.flatMap((request) => {
    var host = store.getState().config.host.slice(0, -1) + ':9090';
		return Observable
			.defer(() => new Promise((resolve, reject) => {
        
				var arr: ProtocTypes.readinglist.Book.AsObject[] = [];
				const client = grpc.client(ProtocServices.readinglist.ReadingList.ReadAllBooks, {
					host: host,
				});
				client.onMessage((message: ProtocTypes.readinglist.Book) => {
          
					arr.push(message.toObject());
				});
        client.onEnd((code: grpc.Code, msg: string) => { 
					if (code != grpc.Code.OK) {
            
						reject(createErrorObject(code, msg));
					}
					resolve(arr);
				});
				client.start();
				client.send(request.message);
			})) 
		.retry(0)
		.timeout(3000)
		.map((resObj: ProtocTypes.readinglist.Book.AsObject[]) => {
			request.resolve(resObj);
			return protocActions.listLibrarySuccess(resObj);
		})
		.catch(error => {
			const err: NodeJS.ErrnoException = createErrorObject(error.code, error.message);
			if(request.reject){ request.reject(err); }
			return Observable.of(protocActions.listLibraryFailure(err));
		})
	})
	.takeUntil(action$.filter(isActionOf(protocActions.listLibraryCancel)))
	.repeat();

export const updateLibraryEpic = (action$, store) => action$
	.filter(isActionOf(protocActions.updateLibraryRequest))
	.debounceTime(510)
	.map(({ payload, meta: { resolve = noop, reject = noop } }) => ({
		message: toMessage(payload, ProtocTypes.readinglist.Book),
		resolve,
		reject,
	}))
	.flatMap((request) => {
    return Observable
		.defer(() => new Promise((resolve, reject) => {
      
			var host = store.getState().config.host.slice(0, -1) + ':9090';
			
			grpc.unary(ProtocServices.readinglist.ReadingList.UpdateBook, {
				request: request.message,
				host: host,
				
				onEnd: (res: UnaryOutput<ProtocTypes.readinglist.Book>) => {
          
					if(res.status != grpc.Code.OK){
            
						const err: NodeJS.ErrnoException = createErrorObject(res.status, res.statusMessage);
						reject(err);
					}
					if(res.message){
						resolve(res.message.toObject());
					}
				}
			});
		})) 
		.retry(0)
		.timeout(3000)
		.map(obj => ({ ...obj } as { prev: ProtocTypes.readinglist.Book.AsObject, updated: ProtocTypes.readinglist.Book.AsObject } ))
		.map(lib => {
			request.resolve(lib.prev, lib.updated);
			return protocActions.updateLibrarySuccess(lib);
		})
		.catch(error => {
			const err: NodeJS.ErrnoException = createErrorObject(error.code, error.message);
			if(request.reject){ request.reject(err); }
			return Observable.of(protocActions.updateLibraryFailure(err));
		})
	})
	.takeUntil(action$.filter(isActionOf(protocActions.updateLibraryCancel)))
	.repeat();

export const deleteLibraryEpic = (action$, store) => action$
	.filter(isActionOf(protocActions.deleteLibraryRequest))
	.debounceTime(510)
	.map(({ payload, meta: { resolve = noop, reject = noop } }) => ({
		message: toMessage(payload, ProtocTypes.readinglist.Book),
		resolve,
		reject,
	}))
	.flatMap((request) => {
    return Observable
		.defer(() => new Promise((resolve, reject) => {
      
			var host = store.getState().config.host.slice(0, -1) + ':9090';
			
			grpc.unary(ProtocServices.readinglist.ReadingList.DeleteBook, {
				request: request.message,
				host: host,
				
				onEnd: (res: UnaryOutput<ProtocTypes.readinglist.Book>) => {
          
					if(res.status != grpc.Code.OK){
            
						const err: NodeJS.ErrnoException = createErrorObject(res.status, res.statusMessage);
						reject(err);
					}
					if(res.message){
						resolve(res.message.toObject());
					}
				}
			});
		})) 
		.retry(0)
		.timeout(3000)
		.map((resObj: ProtocTypes.readinglist.Book.AsObject) => {
			request.resolve(resObj);
			return protocActions.deleteLibrarySuccess(resObj);
		})
		.catch(error => {
			const err: NodeJS.ErrnoException = createErrorObject(error.code, error.message);
			if(request.reject){ request.reject(err); }
			return Observable.of(protocActions.deleteLibraryFailure(err));
		})
	})
	.takeUntil(action$.filter(isActionOf(protocActions.deleteLibraryCancel)))
	.repeat();

export const createBookOfTheMonthEpic = (action$, store) => action$
	.filter(isActionOf(protocActions.createBookOfTheMonthRequest))
	.debounceTime(510)
	.map(({ payload, meta: { resolve = noop, reject = noop } }) => ({
		message: toMessage(payload, ProtocTypes.readinglist.Book),
		resolve,
		reject,
	}))
	.flatMap((request) => {
    return Observable
		.defer(() => new Promise((resolve, reject) => {
      
			var host = store.getState().config.host.slice(0, -1) + ':9090';
			
			grpc.unary(ProtocServices.readinglist.ReadingList.CreateBook, {
				request: request.message,
				host: host,
				
				onEnd: (res: UnaryOutput<ProtocTypes.readinglist.Book>) => {
          
					if(res.status != grpc.Code.OK){
            
						const err: NodeJS.ErrnoException = createErrorObject(res.status, res.statusMessage);
						reject(err);
					}
					if(res.message){
						resolve(res.message.toObject());
					}
				}
			});
		})) 
		.retry(0)
		.timeout(3000)
		.map((resObj: ProtocTypes.readinglist.Book.AsObject) => {
			request.resolve(resObj);
			return protocActions.createBookOfTheMonthSuccess(resObj);
		})
		.catch(error => {
			const err: NodeJS.ErrnoException = createErrorObject(error.code, error.message);
			if(request.reject){ request.reject(err); }
			return Observable.of(protocActions.createBookOfTheMonthFailure(err));
		})
	})
	.takeUntil(action$.filter(isActionOf(protocActions.createBookOfTheMonthCancel)))
	.repeat();

export const getBookOfTheMonthEpic = (action$, store) => action$
	.filter(isActionOf(protocActions.getBookOfTheMonthRequest))
	.debounceTime(510)
	.map(({ payload, meta: { resolve = noop, reject = noop } }) => ({
		message: toMessage(payload, ProtocTypes.readinglist.Book),
		resolve,
		reject,
	}))
	.flatMap((request) => {
    return Observable
		.defer(() => new Promise((resolve, reject) => {
      
			var host = store.getState().config.host.slice(0, -1) + ':9090';
			
			grpc.unary(ProtocServices.readinglist.ReadingList.ReadBook, {
				request: request.message,
				host: host,
				
				onEnd: (res: UnaryOutput<ProtocTypes.readinglist.Book>) => {
          
					if(res.status != grpc.Code.OK){
            
						const err: NodeJS.ErrnoException = createErrorObject(res.status, res.statusMessage);
						reject(err);
					}
					if(res.message){
						resolve(res.message.toObject());
					}
				}
			});
		})) 
		.retry(0)
		.timeout(3000)
		.map((resObj: ProtocTypes.readinglist.Book.AsObject) => {
			request.resolve(resObj);
			return protocActions.getBookOfTheMonthSuccess(resObj);
		})
		.catch(error => {
			const err: NodeJS.ErrnoException = createErrorObject(error.code, error.message);
			if(request.reject){ request.reject(err); }
			return Observable.of(protocActions.getBookOfTheMonthFailure(err));
		})
	})
	.takeUntil(action$.filter(isActionOf(protocActions.getBookOfTheMonthCancel)))
	.repeat();

export const updateBookOfTheMonthEpic = (action$, store) => action$
	.filter(isActionOf(protocActions.updateBookOfTheMonthRequest))
	.debounceTime(510)
	.map(({ payload, meta: { resolve = noop, reject = noop } }) => ({
		message: toMessage(payload, ProtocTypes.readinglist.Book),
		resolve,
		reject,
	}))
	.flatMap((request) => {
    return Observable
		.defer(() => new Promise((resolve, reject) => {
      
			var host = store.getState().config.host.slice(0, -1) + ':9090';
			
			grpc.unary(ProtocServices.readinglist.ReadingList.UpdateBook, {
				request: request.message,
				host: host,
				
				onEnd: (res: UnaryOutput<ProtocTypes.readinglist.Book>) => {
          
					if(res.status != grpc.Code.OK){
            
						const err: NodeJS.ErrnoException = createErrorObject(res.status, res.statusMessage);
						reject(err);
					}
					if(res.message){
						resolve(res.message.toObject());
					}
				}
			});
		})) 
		.retry(0)
		.timeout(3000)
		.map((resObj: ProtocTypes.readinglist.Book.AsObject) => {
			request.resolve(resObj);
			return protocActions.updateBookOfTheMonthSuccess(resObj);
		})
		.catch(error => {
			const err: NodeJS.ErrnoException = createErrorObject(error.code, error.message);
			if(request.reject){ request.reject(err); }
			return Observable.of(protocActions.updateBookOfTheMonthFailure(err));
		})
	})
	.takeUntil(action$.filter(isActionOf(protocActions.updateBookOfTheMonthCancel)))
	.repeat();

export const deleteBookOfTheMonthEpic = (action$, store) => action$
	.filter(isActionOf(protocActions.deleteBookOfTheMonthRequest))
	.debounceTime(510)
	.map(({ payload, meta: { resolve = noop, reject = noop } }) => ({
		message: toMessage(payload, ProtocTypes.readinglist.Book),
		resolve,
		reject,
	}))
	.flatMap((request) => {
    return Observable
		.defer(() => new Promise((resolve, reject) => {
      
			var host = store.getState().config.host.slice(0, -1) + ':9090';
			
			grpc.unary(ProtocServices.readinglist.ReadingList.DeleteBook, {
				request: request.message,
				host: host,
				
				onEnd: (res: UnaryOutput<ProtocTypes.readinglist.Book>) => {
          
					if(res.status != grpc.Code.OK){
            
						const err: NodeJS.ErrnoException = createErrorObject(res.status, res.statusMessage);
						reject(err);
					}
					if(res.message){
						resolve(res.message.toObject());
					}
				}
			});
		})) 
		.retry(0)
		.timeout(3000)
		.map((resObj: ProtocTypes.readinglist.Book.AsObject) => {
			request.resolve(resObj);
			return protocActions.deleteBookOfTheMonthSuccess(resObj);
		})
		.catch(error => {
			const err: NodeJS.ErrnoException = createErrorObject(error.code, error.message);
			if(request.reject){ request.reject(err); }
			return Observable.of(protocActions.deleteBookOfTheMonthFailure(err));
		})
	})
	.takeUntil(action$.filter(isActionOf(protocActions.deleteBookOfTheMonthCancel)))
	.repeat();

export const getTimeoutBookEpic = (action$, store) => action$
	.filter(isActionOf(protocActions.getTimeoutBookRequest))
	.debounceTime(510)
	.map(({ payload, meta: { resolve = noop, reject = noop } }) => ({
		message: toMessage(payload, ProtocTypes.readinglist.Book),
		resolve,
		reject,
	}))
	.flatMap((request) => {
    return Observable
		.defer(() => new Promise((resolve, reject) => {
      
			var host = store.getState().config.host.slice(0, -1) + ':9090';
			
			grpc.unary(ProtocServices.readinglist.ReadingList.ReadBook, {
				request: request.message,
				host: host,
				
				onEnd: (res: UnaryOutput<ProtocTypes.readinglist.Book>) => {
          
					if(res.status != grpc.Code.OK){
            
						const err: NodeJS.ErrnoException = createErrorObject(res.status, res.statusMessage);
						reject(err);
					}
					if(res.message){
						resolve(res.message.toObject());
					}
				}
			});
		})) 
		.retry(0)
		.timeout(3000)
		.map((resObj: ProtocTypes.readinglist.Book.AsObject) => {
			request.resolve(resObj);
			return protocActions.getTimeoutBookSuccess(resObj);
		})
		.catch(error => {
			const err: NodeJS.ErrnoException = createErrorObject(error.code, error.message);
			if(request.reject){ request.reject(err); }
			return Observable.of(protocActions.getTimeoutBookFailure(err));
		})
	})
	.takeUntil(action$.filter(isActionOf(protocActions.getTimeoutBookCancel)))
	.repeat();

export const customErrorBookEpic = (action$, store) => action$
	.filter(isActionOf(protocActions.customErrorBookRequest))
	.debounceTime(510)
	.map(({ payload, meta: { resolve = noop, reject = noop } }) => ({
		message: toMessage(payload, ProtocTypes.readinglist.Book),
		resolve,
		reject,
	}))
	.flatMap((request) => {
    return Observable
		.defer(() => new Promise((resolve, reject) => {
      
			var host = store.getState().config.host.slice(0, -1) + ':9090';
			
			grpc.unary(ProtocServices.readinglist.ReadingList.ErrorOut, {
				request: request.message,
				host: host,
				
				onEnd: (res: UnaryOutput<ProtocTypes.readinglist.Book>) => {
          
					if(res.status != grpc.Code.OK){
            
						const err: NodeJS.ErrnoException = createErrorObject(res.status, res.statusMessage);
						reject(err);
					}
					if(res.message){
						resolve(res.message.toObject());
					}
				}
			});
		})) 
		.retry(0)
		.timeout(3000)
		.map((resObj: ProtocTypes.readinglist.Book.AsObject) => {
			request.resolve(resObj);
			return protocActions.customErrorBookSuccess(resObj);
		})
		.catch(error => {
			const err: NodeJS.ErrnoException = createErrorObject(error.code, error.message);
			if(request.reject){ request.reject(err); }
			return Observable.of(protocActions.customErrorBookFailure(err));
		})
	})
	.takeUntil(action$.filter(isActionOf(protocActions.customErrorBookCancel)))
	.repeat();


export const protocEpics = combineEpics(
	createLibraryEpic,
	listLibraryEpic,
	updateLibraryEpic,
	deleteLibraryEpic,
	createBookOfTheMonthEpic,
	getBookOfTheMonthEpic,
	updateBookOfTheMonthEpic,
	deleteBookOfTheMonthEpic,
	getTimeoutBookEpic,
	customErrorBookEpic,
)