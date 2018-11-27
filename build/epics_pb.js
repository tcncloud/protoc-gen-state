/* THIS FILE IS GENERATED FROM THE TOOL PROTOC-GEN-STATE  */
/* ANYTHING YOU EDIT WILL BE OVERWRITTEN IN FUTURE BUILDS */
import { combineEpics } from 'redux-observable';
import { isActionOf } from 'typesafe-actions';
import { Observable } from 'rxjs';
import { grpc } from 'grpc-web-client';
import 'rxjs/add/observable/dom/ajax';
import { toMessage } from './to_message_pb';
import * as protocActions from './actions_pb';
import * as ProtocTypes from './protoc_types_pb';
import * as ProtocServices from './protoc_services_pb';
function createErrorObject(code, message) {
    var err = new Error();
    err.message = message;
    if (code && typeof code == 'number') {
        err.code = code.toString();
    }
    if (code && typeof code == 'string') {
        err.code = code;
    }
    return err;
}
export const createLibraryEpic = (action$, store) => action$
    .filter(isActionOf([
    protocActions.createLibraryRequest,
    protocActions.createLibraryRequestPromise,
]))
    .debounceTime(510)
    .map((action) => {
    if (action.payload && action.payload.resolve && action.payload.reject) {
        return {
            ...action.payload,
            request: toMessage(action.payload.library, ProtocTypes.readinglist.Book)
        };
    }
    else {
        return { request: toMessage(action.payload, ProtocTypes.readinglist.Book) };
    }
})
    .flatMap((action) => {
    var host = store.getState().config.host.slice(0, -1) + ':9090';
    return Observable
        .defer(() => new Promise((resolve, reject) => {
        var arr = [];
        const client = grpc.client(ProtocServices.readinglist.ReadingList.CreateBook, {
            host: host,
        });
        client.onMessage((message) => {
            arr.push(message.toObject());
        });
        client.onEnd((code, msg) => {
            if (code != grpc.Code.OK) {
                reject(createErrorObject(code, msg));
            }
            resolve(arr);
        });
        client.start();
        client.send(action.request);
    }))
        .retry(0)
        .timeout(3000)
        .map(resObj => {
        if (action.resolve) {
            action.resolve(resObj);
        }
        return protocActions.createLibrarySuccess(resObj);
    })
        .catch(error => {
        const err = createErrorObject(error.code, error.message);
        if (action.reject) {
            action.reject(err);
        }
        return Observable.of(protocActions.createLibraryFailure(err));
    });
})
    .takeUntil(action$.filter(isActionOf(protocActions.createLibraryCancel)))
    .repeat();
export const listLibraryEpic = (action$, store) => action$
    .filter(isActionOf([
    protocActions.listLibraryRequest,
    protocActions.listLibraryRequestPromise,
]))
    .debounceTime(510)
    .map((action) => {
    if (action.payload && action.payload.resolve && action.payload.reject) {
        return {
            ...action.payload,
            request: toMessage(action.payload.library, ProtocTypes.readinglist.Empty)
        };
    }
    else {
        return { request: toMessage(action.payload, ProtocTypes.readinglist.Empty) };
    }
})
    .flatMap((action) => {
    var host = store.getState().config.host.slice(0, -1) + ':9090';
    return Observable
        .defer(() => new Promise((resolve, reject) => {
        var arr = [];
        const client = grpc.client(ProtocServices.readinglist.ReadingList.ReadAllBooks, {
            host: host,
        });
        client.onMessage((message) => {
            arr.push(message.toObject());
        });
        client.onEnd((code, msg) => {
            if (code != grpc.Code.OK) {
                reject(createErrorObject(code, msg));
            }
            resolve(arr);
        });
        client.start();
        client.send(action.request);
    }))
        .retry(0)
        .timeout(3000)
        .map(resObj => {
        if (action.resolve) {
            action.resolve(resObj);
        }
        return protocActions.listLibrarySuccess(resObj);
    })
        .catch(error => {
        const err = createErrorObject(error.code, error.message);
        if (action.reject) {
            action.reject(err);
        }
        return Observable.of(protocActions.listLibraryFailure(err));
    });
})
    .takeUntil(action$.filter(isActionOf(protocActions.listLibraryCancel)))
    .repeat();
export const updateLibraryEpic = (action$, store) => action$
    .filter(isActionOf([
    protocActions.updateLibraryRequest,
    protocActions.updateLibraryRequestPromise,
]))
    .debounceTime(510)
    .map((action) => {
    if (action.payload && action.payload.resolve && action.payload.reject) {
        return {
            ...action.payload,
            request: toMessage(action.payload.library, ProtocTypes.readinglist.Book)
        };
    }
    else {
        return { request: toMessage(action.payload, ProtocTypes.readinglist.Book) };
    }
})
    .flatMap((action) => {
    var host = store.getState().config.host.slice(0, -1) + ':9090';
    return Observable
        .defer(() => new Promise((resolve, reject) => {
        var arr = [];
        const client = grpc.client(ProtocServices.readinglist.ReadingList.UpdateBook, {
            host: host,
        });
        client.onMessage((message) => {
            arr.push(message.toObject());
        });
        client.onEnd((code, msg) => {
            if (code != grpc.Code.OK) {
                reject(createErrorObject(code, msg));
            }
            resolve(arr);
        });
        client.start();
        client.send(action.request);
    }))
        .retry(0)
        .timeout(3000)
        .map(resObj => {
        if (action.resolve) {
            action.resolve(resObj);
        }
        return protocActions.updateLibrarySuccess(resObj);
    })
        .catch(error => {
        const err = createErrorObject(error.code, error.message);
        if (action.reject) {
            action.reject(err);
        }
        return Observable.of(protocActions.updateLibraryFailure(err));
    });
})
    .takeUntil(action$.filter(isActionOf(protocActions.updateLibraryCancel)))
    .repeat();
export const deleteLibraryEpic = (action$, store) => action$
    .filter(isActionOf([
    protocActions.deleteLibraryRequest,
    protocActions.deleteLibraryRequestPromise,
]))
    .debounceTime(510)
    .map((action) => {
    if (action.payload && action.payload.resolve && action.payload.reject) {
        return {
            ...action.payload,
            request: toMessage(action.payload.library, ProtocTypes.readinglist.Book)
        };
    }
    else {
        return { request: toMessage(action.payload, ProtocTypes.readinglist.Book) };
    }
})
    .flatMap((action) => {
    var host = store.getState().config.host.slice(0, -1) + ':9090';
    return Observable
        .defer(() => new Promise((resolve, reject) => {
        var arr = [];
        const client = grpc.client(ProtocServices.readinglist.ReadingList.DeleteBook, {
            host: host,
        });
        client.onMessage((message) => {
            arr.push(message.toObject());
        });
        client.onEnd((code, msg) => {
            if (code != grpc.Code.OK) {
                reject(createErrorObject(code, msg));
            }
            resolve(arr);
        });
        client.start();
        client.send(action.request);
    }))
        .retry(0)
        .timeout(3000)
        .map(resObj => {
        if (action.resolve) {
            action.resolve(resObj);
        }
        return protocActions.deleteLibrarySuccess(resObj);
    })
        .catch(error => {
        const err = createErrorObject(error.code, error.message);
        if (action.reject) {
            action.reject(err);
        }
        return Observable.of(protocActions.deleteLibraryFailure(err));
    });
})
    .takeUntil(action$.filter(isActionOf(protocActions.deleteLibraryCancel)))
    .repeat();
export const createBookOfTheMonthEpic = (action$, store) => action$
    .filter(isActionOf([
    protocActions.createBookOfTheMonthRequest,
    protocActions.createBookOfTheMonthRequestPromise,
]))
    .debounceTime(510)
    .map((action) => {
    if (action.payload && action.payload.resolve && action.payload.reject) {
        return {
            ...action.payload,
            request: toMessage(action.payload.library, ProtocTypes.readinglist.Book)
        };
    }
    else {
        return { request: toMessage(action.payload, ProtocTypes.readinglist.Book) };
    }
})
    .flatMap((action) => {
    return Observable
        .defer(() => new Promise((resolve, reject) => {
        var host = store.getState().config.host.slice(0, -1) + ':9090';
        grpc.unary(ProtocServices.readinglist.ReadingList.CreateBook, {
            request: action.request,
            host: host,
            onEnd: (res) => {
                if (res.status != grpc.Code.OK) {
                    const err = createErrorObject(res.status, res.statusMessage);
                    reject(err);
                }
                if (res.message) {
                    resolve(res.message.toObject());
                }
            }
        });
    }))
        .retry(0)
        .timeout(3000)
        .map(resObj => {
        if (action.resolve) {
            action.resolve(resObj);
        }
        return protocActions.createBookOfTheMonthSuccess(resObj);
    })
        .catch(error => {
        const err = createErrorObject(error.code, error.message);
        if (action.reject) {
            action.reject(err);
        }
        return Observable.of(protocActions.createBookOfTheMonthFailure(err));
    });
})
    .takeUntil(action$.filter(isActionOf(protocActions.createBookOfTheMonthCancel)))
    .repeat();
export const getBookOfTheMonthEpic = (action$, store) => action$
    .filter(isActionOf([
    protocActions.getBookOfTheMonthRequest,
    protocActions.getBookOfTheMonthRequestPromise,
]))
    .debounceTime(510)
    .map((action) => {
    if (action.payload && action.payload.resolve && action.payload.reject) {
        return {
            ...action.payload,
            request: toMessage(action.payload.library, ProtocTypes.readinglist.Book)
        };
    }
    else {
        return { request: toMessage(action.payload, ProtocTypes.readinglist.Book) };
    }
})
    .flatMap((action) => {
    return Observable
        .defer(() => new Promise((resolve, reject) => {
        var host = store.getState().config.host.slice(0, -1) + ':9090';
        grpc.unary(ProtocServices.readinglist.ReadingList.ReadBook, {
            request: action.request,
            host: host,
            onEnd: (res) => {
                if (res.status != grpc.Code.OK) {
                    const err = createErrorObject(res.status, res.statusMessage);
                    reject(err);
                }
                if (res.message) {
                    resolve(res.message.toObject());
                }
            }
        });
    }))
        .retry(0)
        .timeout(3000)
        .map(resObj => {
        if (action.resolve) {
            action.resolve(resObj);
        }
        return protocActions.getBookOfTheMonthSuccess(resObj);
    })
        .catch(error => {
        const err = createErrorObject(error.code, error.message);
        if (action.reject) {
            action.reject(err);
        }
        return Observable.of(protocActions.getBookOfTheMonthFailure(err));
    });
})
    .takeUntil(action$.filter(isActionOf(protocActions.getBookOfTheMonthCancel)))
    .repeat();
export const updateBookOfTheMonthEpic = (action$, store) => action$
    .filter(isActionOf([
    protocActions.updateBookOfTheMonthRequest,
    protocActions.updateBookOfTheMonthRequestPromise,
]))
    .debounceTime(510)
    .map((action) => {
    if (action.payload && action.payload.resolve && action.payload.reject) {
        return {
            ...action.payload,
            request: toMessage(action.payload.library, ProtocTypes.readinglist.Book)
        };
    }
    else {
        return { request: toMessage(action.payload, ProtocTypes.readinglist.Book) };
    }
})
    .flatMap((action) => {
    return Observable
        .defer(() => new Promise((resolve, reject) => {
        var host = store.getState().config.host.slice(0, -1) + ':9090';
        grpc.unary(ProtocServices.readinglist.ReadingList.UpdateBook, {
            request: action.request,
            host: host,
            onEnd: (res) => {
                if (res.status != grpc.Code.OK) {
                    const err = createErrorObject(res.status, res.statusMessage);
                    reject(err);
                }
                if (res.message) {
                    resolve(res.message.toObject());
                }
            }
        });
    }))
        .retry(0)
        .timeout(3000)
        .map(resObj => {
        if (action.resolve) {
            action.resolve(resObj);
        }
        return protocActions.updateBookOfTheMonthSuccess(resObj);
    })
        .catch(error => {
        const err = createErrorObject(error.code, error.message);
        if (action.reject) {
            action.reject(err);
        }
        return Observable.of(protocActions.updateBookOfTheMonthFailure(err));
    });
})
    .takeUntil(action$.filter(isActionOf(protocActions.updateBookOfTheMonthCancel)))
    .repeat();
export const deleteBookOfTheMonthEpic = (action$, store) => action$
    .filter(isActionOf([
    protocActions.deleteBookOfTheMonthRequest,
    protocActions.deleteBookOfTheMonthRequestPromise,
]))
    .debounceTime(510)
    .map((action) => {
    if (action.payload && action.payload.resolve && action.payload.reject) {
        return {
            ...action.payload,
            request: toMessage(action.payload.library, ProtocTypes.readinglist.Book)
        };
    }
    else {
        return { request: toMessage(action.payload, ProtocTypes.readinglist.Book) };
    }
})
    .flatMap((action) => {
    return Observable
        .defer(() => new Promise((resolve, reject) => {
        var host = store.getState().config.host.slice(0, -1) + ':9090';
        grpc.unary(ProtocServices.readinglist.ReadingList.DeleteBook, {
            request: action.request,
            host: host,
            onEnd: (res) => {
                if (res.status != grpc.Code.OK) {
                    const err = createErrorObject(res.status, res.statusMessage);
                    reject(err);
                }
                if (res.message) {
                    resolve(res.message.toObject());
                }
            }
        });
    }))
        .retry(0)
        .timeout(3000)
        .map(resObj => {
        if (action.resolve) {
            action.resolve(resObj);
        }
        return protocActions.deleteBookOfTheMonthSuccess(resObj);
    })
        .catch(error => {
        const err = createErrorObject(error.code, error.message);
        if (action.reject) {
            action.reject(err);
        }
        return Observable.of(protocActions.deleteBookOfTheMonthFailure(err));
    });
})
    .takeUntil(action$.filter(isActionOf(protocActions.deleteBookOfTheMonthCancel)))
    .repeat();
export const getTimeoutBookEpic = (action$, store) => action$
    .filter(isActionOf([
    protocActions.getTimeoutBookRequest,
    protocActions.getTimeoutBookRequestPromise,
]))
    .debounceTime(510)
    .map((action) => {
    if (action.payload && action.payload.resolve && action.payload.reject) {
        return {
            ...action.payload,
            request: toMessage(action.payload.library, ProtocTypes.readinglist.Book)
        };
    }
    else {
        return { request: toMessage(action.payload, ProtocTypes.readinglist.Book) };
    }
})
    .flatMap((action) => {
    return Observable
        .defer(() => new Promise((resolve, reject) => {
        var host = store.getState().config.host.slice(0, -1) + ':9090';
        grpc.unary(ProtocServices.readinglist.ReadingList.ReadBook, {
            request: action.request,
            host: host,
            onEnd: (res) => {
                if (res.status != grpc.Code.OK) {
                    const err = createErrorObject(res.status, res.statusMessage);
                    reject(err);
                }
                if (res.message) {
                    resolve(res.message.toObject());
                }
            }
        });
    }))
        .retry(0)
        .timeout(3000)
        .map(resObj => {
        if (action.resolve) {
            action.resolve(resObj);
        }
        return protocActions.getTimeoutBookSuccess(resObj);
    })
        .catch(error => {
        const err = createErrorObject(error.code, error.message);
        if (action.reject) {
            action.reject(err);
        }
        return Observable.of(protocActions.getTimeoutBookFailure(err));
    });
})
    .takeUntil(action$.filter(isActionOf(protocActions.getTimeoutBookCancel)))
    .repeat();
export const customErrorBookEpic = (action$, store) => action$
    .filter(isActionOf([
    protocActions.customErrorBookRequest,
    protocActions.customErrorBookRequestPromise,
]))
    .debounceTime(510)
    .map((action) => {
    if (action.payload && action.payload.resolve && action.payload.reject) {
        return {
            ...action.payload,
            request: toMessage(action.payload.library, ProtocTypes.readinglist.Book)
        };
    }
    else {
        return { request: toMessage(action.payload, ProtocTypes.readinglist.Book) };
    }
})
    .flatMap((action) => {
    return Observable
        .defer(() => new Promise((resolve, reject) => {
        var host = store.getState().config.host.slice(0, -1) + ':9090';
        grpc.unary(ProtocServices.readinglist.ReadingList.ErrorOut, {
            request: action.request,
            host: host,
            onEnd: (res) => {
                if (res.status != grpc.Code.OK) {
                    const err = createErrorObject(res.status, res.statusMessage);
                    reject(err);
                }
                if (res.message) {
                    resolve(res.message.toObject());
                }
            }
        });
    }))
        .retry(0)
        .timeout(3000)
        .map(resObj => {
        if (action.resolve) {
            action.resolve(resObj);
        }
        return protocActions.customErrorBookSuccess(resObj);
    })
        .catch(error => {
        const err = createErrorObject(error.code, error.message);
        if (action.reject) {
            action.reject(err);
        }
        return Observable.of(protocActions.customErrorBookFailure(err));
    });
})
    .takeUntil(action$.filter(isActionOf(protocActions.customErrorBookCancel)))
    .repeat();
export const protocEpics = combineEpics(createLibraryEpic, listLibraryEpic, updateLibraryEpic, deleteLibraryEpic, createBookOfTheMonthEpic, getBookOfTheMonthEpic, updateBookOfTheMonthEpic, deleteBookOfTheMonthEpic, getTimeoutBookEpic, customErrorBookEpic);
//# sourceMappingURL=epics_pb.js.map