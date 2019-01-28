import { protocReducer } from 'protos/BasicState/reducer_pb';
import { ProtocState, initialProtocState } from 'protos/BasicState/state_pb';
import * as ProtocActions from 'protos/BasicState/actions_pb';
import * as ProtocTypes from 'protos/BasicState/protoc_types_pb';

describe('Action Existance Tests', () => {
  describe('Single Entity', () => {
    it('should generate Create actions', () => {
      expect(typeof ProtocActions.createBookOfTheMonthRequest).toBe("function", "Request action not defined")
      expect(typeof ProtocActions.createBookOfTheMonthRequestPromise).toBe("function", "Request Promise action not defined")
      expect(typeof ProtocActions.createBookOfTheMonthSuccess).toBe("function", "Success action not defined")
      expect(typeof ProtocActions.createBookOfTheMonthFailure).toBe("function", "Failure action not defined")
      expect(typeof ProtocActions.createBookOfTheMonthCancel).toBe("function", "Cancel action not defined")
    })
    it('should generate Update actions', () => {
      expect(typeof ProtocActions.updateBookOfTheMonthRequest).toBe("function", "Request action not defined")
      expect(typeof ProtocActions.updateBookOfTheMonthRequestPromise).toBe("function", "Request Promise action not defined")
      expect(typeof ProtocActions.updateBookOfTheMonthSuccess).toBe("function", "Success action not defined")
      expect(typeof ProtocActions.updateBookOfTheMonthFailure).toBe("function", "Failure action not defined")
      expect(typeof ProtocActions.updateBookOfTheMonthCancel).toBe("function", "Cancel action not defined")
    })
    it('should generate Delete actions', () => {
      expect(typeof ProtocActions.deleteBookOfTheMonthRequest).toBe("function", "Request action not defined")
      expect(typeof ProtocActions.deleteBookOfTheMonthRequestPromise).toBe("function", "Request Promise action not defined")
      expect(typeof ProtocActions.deleteBookOfTheMonthSuccess).toBe("function", "Success action not defined")
      expect(typeof ProtocActions.deleteBookOfTheMonthFailure).toBe("function", "Failure action not defined")
      expect(typeof ProtocActions.deleteBookOfTheMonthCancel).toBe("function", "Cancel action not defined")
    })
    it('should generate Get actions', () => {
      expect(typeof ProtocActions.getBookOfTheMonthRequest).toBe("function", "Request action not defined")
      expect(typeof ProtocActions.getBookOfTheMonthRequestPromise).toBe("function", "Request Promise action not defined")
      expect(typeof ProtocActions.getBookOfTheMonthSuccess).toBe("function", "Success action not defined")
      expect(typeof ProtocActions.getBookOfTheMonthFailure).toBe("function", "Failure action not defined")
      expect(typeof ProtocActions.getBookOfTheMonthCancel).toBe("function", "Cancel action not defined")
    })
    it('should generate Reset action', () => {
      expect(typeof ProtocActions.resetBookOfTheMonth).toBe("function", "Reset action not defined")
    })
  })
  describe('Repeated Entities', () => {
    it('should generate Create actions', () => {
      expect(typeof ProtocActions.createLibraryRequest).toBe("function", "Request action not defined")
      expect(typeof ProtocActions.createLibraryRequestPromise).toBe("function", "Request Promise action not defined")
      expect(typeof ProtocActions.createLibrarySuccess).toBe("function", "Success action not defined")
      expect(typeof ProtocActions.createLibraryFailure).toBe("function", "Failure action not defined")
      expect(typeof ProtocActions.createLibraryCancel).toBe("function", "Cancel action not defined")
    })
    it('should generate Update actions', () => {
      expect(typeof ProtocActions.updateLibraryRequest).toBe("function", "Request action not defined")
      expect(typeof ProtocActions.updateLibraryRequestPromise).toBe("function", "Request Promise action not defined")
      expect(typeof ProtocActions.updateLibrarySuccess).toBe("function", "Success action not defined")
      expect(typeof ProtocActions.updateLibraryFailure).toBe("function", "Failure action not defined")
      expect(typeof ProtocActions.updateLibraryCancel).toBe("function", "Cancel action not defined")
    })
    it('should generate Delete actions', () => {
      expect(typeof ProtocActions.deleteLibraryRequest).toBe("function", "Request action not defined")
      expect(typeof ProtocActions.deleteLibraryRequestPromise).toBe("function", "Request Promise action not defined")
      expect(typeof ProtocActions.deleteLibrarySuccess).toBe("function", "Success action not defined")
      expect(typeof ProtocActions.deleteLibraryFailure).toBe("function", "Failure action not defined")
      expect(typeof ProtocActions.deleteLibraryCancel).toBe("function", "Cancel action not defined")
    })
    it('should generate List actions', () => {
      expect(typeof ProtocActions.listLibraryRequest).toBe("function", "Request action not defined")
      expect(typeof ProtocActions.listLibraryRequestPromise).toBe("function", "Request Promise action not defined")
      expect(typeof ProtocActions.listLibrarySuccess).toBe("function", "Success action not defined")
      expect(typeof ProtocActions.listLibraryFailure).toBe("function", "Failure action not defined")
      expect(typeof ProtocActions.listLibraryCancel).toBe("function", "Cancel action not defined")
    })
    it('should generate Reset action', () => {
      expect(typeof ProtocActions.resetLibrary).toBe("function", "Reset action not defined")
    })
  })
})

describe('Action Type Tests', () => {
  describe('Single Entity', () => {
    it('should follow the same format for all the type strings', () => {
      let payload : string = "BookOfTheMonth"
      let cruds : string[] = ["create", "update", "delete", "get"]
      let effects : string[] = ["Request", "Success", "Failure", "Cancel"]
      for(var cIndex : number = 0; cIndex < cruds.length; cIndex++){
        for(var eIndex : number = 0; eIndex < effects.length; eIndex++){
          expect(ProtocActions[cruds[cIndex]+payload+effects[eIndex]]().type)
            .toBe("PROTOC_" + cruds[cIndex].toUpperCase() + "_" + payload.toUpperCase() + "_" + effects[eIndex].toUpperCase())
        }
      }
    })
  })
  describe('Repeated Entity', () => {
    it('should follow the same format for all the type strings', () => {
      let payload : string = "Library"
      let cruds : string[] = ["create", "update", "delete", "list"]
      let effects : string[] = ["Request", "Success", "Failure", "Cancel"]
      for(var cIndex : number = 0; cIndex < cruds.length; cIndex++){
        for(var eIndex : number = 0; eIndex < effects.length; eIndex++){
          expect(ProtocActions[cruds[cIndex]+payload+effects[eIndex]]().type)
            .toBe("PROTOC_" + cruds[cIndex].toUpperCase() + "_" + payload.toUpperCase() + "_" + effects[eIndex].toUpperCase())
        }
      }
    })
  })
})

describe('Action Payload Tests', () => {
  let bookArr : ProtocTypes.readinglist.Book.AsObject[] = [];
  let bookObj : ProtocTypes.readinglist.Book.AsObject;
  let error : NodeJS.ErrnoException;

  beforeAll(() => {
    let myBook = new ProtocTypes.readinglist.Book();
    myBook.setTitle("The Prince");
    myBook.setAuthor("Niccolo Machiavelli");
    bookObj = myBook.toObject();

    bookArr.push(bookObj);

    error = new Error("Test Error");
    error.code = "14";
  })

  describe('Create Actions', () => {
    describe('Single Entity Request', () => {
      it('should return the book as the payload', () => {
        let result = ProtocActions.createBookOfTheMonthRequest(bookObj);
        expect(result.payload).toEqual(bookObj);
        expect(result.meta).toEqual({ resolve: undefined, reject: undefined });
      })
    })
    describe('Single Entity Request Promise', () => {
      it('should return the book as the payload', () => {
        let resolve = function(){};
        let reject = function(){};
        let result = ProtocActions.createBookOfTheMonthRequestPromise(bookObj, resolve, reject);
        expect(result.payload).toEqual(bookObj);
        expect(result.meta).toEqual({ resolve, reject });
      })
    })
    describe('Single Entity Success', () => {
      it('should return the book as the payload', () => {
        let result = ProtocActions.createBookOfTheMonthSuccess(bookObj);
        expect(result.payload).toEqual(bookObj);
      })
    })
    describe('Single Entity Failure', () => {
      it('should return the error as the payload', () => {
        let result = ProtocActions.createBookOfTheMonthFailure(error);
        expect(result.payload).toEqual(error);
      })
    })
    describe('Single Entity Cancel', () => {
      it('should return no payload', () => {
        let result = ProtocActions.createBookOfTheMonthCancel();
        expect(result["payload"]).toEqual(undefined);
      })
    })
    describe('Repeated Entity Request', () => {
      it('should return the book as the payload', () => {
        let result = ProtocActions.createLibraryRequest(bookObj);
        expect(result.payload).toEqual(bookObj);
        expect(result.meta).toEqual({ resolve: undefined, reject: undefined });
      })
    })
    describe('Repeated Entity Request Promise', () => {
      it('should return the book as the payload', () => {
        let resolve = function(){};
        let reject = function(){};
        let result = ProtocActions.createLibraryRequestPromise(bookObj, resolve, reject);
        expect(result.payload).toEqual(bookObj);
        expect(result.meta).toEqual({ resolve, reject });
      })
    })
    describe('Repeated Entity Success', () => {
      it('should return the book as the payload', () => {
        let result = ProtocActions.createLibrarySuccess(bookObj);
        expect(result.payload).toEqual(bookObj);
      })
    })
    describe('Repeated Entity Failure', () => {
      it('should return the error as the payload', () => {
        let result = ProtocActions.createLibraryFailure(error);
        expect(result.payload).toEqual(error);
      })
    })
    describe('Repeated Entity Cancel', () => {
      it('should return no payload', () => {
        let result = ProtocActions.createLibraryCancel();
        expect(result["payload"]).toEqual(undefined);
      })
    })
  })

  describe('Read Actions', () => {
    describe('Single Entity Request', () => {
      it('should return the book as the payload', () => {
        let result = ProtocActions.getBookOfTheMonthRequest(bookObj);
        expect(result.payload).toEqual(bookObj);
        expect(result.meta).toEqual({ resolve: undefined, reject: undefined });
      })
    })
    describe('Single Entity Request Promise', () => {
      it('should return the book as the payload', () => {
        let resolve = function(){};
        let reject = function(){};
        let result = ProtocActions.getBookOfTheMonthRequestPromise(bookObj, resolve, reject);
        expect(result.payload).toEqual(bookObj);
        expect(result.meta).toEqual({ resolve, reject });
      })
    })
    describe('Single Entity Success', () => {
      it('should return the book as the payload', () => {
        let result = ProtocActions.getBookOfTheMonthSuccess(bookObj);
        expect(result.payload).toEqual(bookObj);
      })
    })
    describe('Single Entity Failure', () => {
      it('should return the error as the payload', () => {
        let result = ProtocActions.getBookOfTheMonthFailure(error);
        expect(result.payload).toEqual(error);
      })
    })
    describe('Single Entity Cancel', () => {
      it('should return no payload', () => {
        let result = ProtocActions.getBookOfTheMonthCancel();
        expect(result["payload"]).toEqual(undefined);
      })
    })
    describe('Repeated Entity Request', () => {
      it('should return the book as the payload', () => {
        let result = ProtocActions.listLibraryRequest(bookObj);
        expect(result.payload).toEqual(bookObj);
        expect(result.meta).toEqual({ resolve: undefined, reject: undefined });
      })
    })
    describe('Repeated Entity Request Promise', () => {
      it('should return the book as the payload', () => {
        let resolve = function(){};
        let reject = function(){};
        let result = ProtocActions.listLibraryRequestPromise(bookObj, resolve, reject);
        expect(result.payload).toEqual(bookObj);
        expect(result.meta).toEqual({ resolve, reject });
      })
    })
    describe('Repeated Entity Success', () => {
      it('should return the book as the payload', () => {
        let result = ProtocActions.listLibrarySuccess(bookArr);
        expect(result.payload).toEqual(bookArr);
      })
    })
    describe('Repeated Entity Failure', () => {
      it('should return the error as the payload', () => {
        let result = ProtocActions.listLibraryFailure(error);
        expect(result.payload).toEqual(error);
      })
    })
    describe('Repeated Entity Cancel', () => {
      it('should return no payload', () => {
        let result = ProtocActions.listLibraryCancel();
        expect(result["payload"]).toEqual(undefined);
      })
    })
  })

  describe('Update Actions', () => {
    let newBook = new ProtocTypes.readinglist.Book();
    newBook.setTitle('The Prince');
    newBook.setAuthor('Francine Rivers');
    let newBookObj : ProtocTypes.readinglist.Book.AsObject = newBook.toObject();

    describe('Single Entity Request', () => {
      it('should return the book as the payload', () => {
        let result = ProtocActions.updateBookOfTheMonthRequest(bookObj);
        expect(result.payload).toEqual(bookObj);
        expect(result.meta).toEqual({ resolve: undefined, reject: undefined });
      })
    })
    describe('Single Entity Request Promise', () => {
      it('should return the book as the payload', () => {
        let resolve = function(){};
        let reject = function(){};
        let result = ProtocActions.updateBookOfTheMonthRequestPromise(bookObj, resolve, reject);
        expect(result.payload).toEqual(bookObj);
        expect(result.meta).toEqual({ resolve, reject });
      })
    })
    describe('Single Entity Success', () => {
      it('should return the book as the payload', () => {
        let result = ProtocActions.updateBookOfTheMonthSuccess(bookObj);
        expect(result.payload).toEqual(bookObj);
      })
    })
    describe('Single Entity Failure', () => {
      it('should return the error as the payload', () => {
        let result = ProtocActions.updateBookOfTheMonthFailure(error);
        expect(result.payload).toEqual(error);
      })
    })
    describe('Single Entity Cancel', () => {
      it('should return no payload', () => {
        let result = ProtocActions.updateBookOfTheMonthCancel();
        expect(result["payload"]).toEqual(undefined);
      })
    })
    describe('Repeated Entity Request', () => {
      it('should return the book as the payload', () => {
        let result = ProtocActions.updateLibraryRequest(bookObj, newBookObj);
        expect(result.payload).toEqual({ prev: bookObj, updated: newBookObj });
        expect(result.meta).toEqual({ resolve: undefined, reject: undefined });
      })
    })
    describe('Repeated Entity Request Promise', () => {
      it('should return the book as the payload', () => {
        let resolve = function(){};
        let reject = function(){};
        let result = ProtocActions.updateLibraryRequestPromise(bookObj, newBookObj, resolve, reject);
        expect(result.payload).toEqual({ prev: bookObj, updated: newBookObj });
        expect(result.meta).toEqual({ resolve, reject });
      })
    })
    describe('Repeated Entity Success', () => {
      it('should return the book as the payload', () => { let result = ProtocActions.updateLibrarySuccess({ prev: bookObj, updated: newBookObj });
        expect(result.payload).toEqual({ prev: bookObj, updated: newBookObj });
      })
    })
    describe('Repeated Entity Failure', () => {
      it('should return the error as the payload', () => {
        let result = ProtocActions.updateLibraryFailure(error);
        expect(result.payload).toEqual(error);
      })
    })
    describe('Repeated Entity Cancel', () => {
      it('should return no payload', () => {
        let result = ProtocActions.updateLibraryCancel();
        expect(result["payload"]).toEqual(undefined);
      })
    })
  })

  describe('Delete Actions', () => {
    describe('Single Entity Request', () => {
      it('should return the book as the payload', () => {
        let result = ProtocActions.deleteBookOfTheMonthRequest(bookObj);
        expect(result.payload).toEqual(bookObj);
        expect(result.meta).toEqual({ resolve: undefined, reject: undefined });
      })
    })
    describe('Single Entity Request Promise', () => {
      it('should return the book as the payload', () => {
        let resolve = function(){};
        let reject = function(){};
        let result = ProtocActions.deleteBookOfTheMonthRequestPromise(bookObj, resolve, reject);
        expect(result.payload).toEqual(bookObj);
        expect(result.meta).toEqual({ resolve, reject });
      })
    })
    describe('Single Entity Success', () => {
      it('should return the book as the payload', () => {
        let result = ProtocActions.deleteBookOfTheMonthSuccess(bookObj);
        expect(result.payload).toEqual(bookObj);
      })
    })
    describe('Single Entity Failure', () => {
      it('should return the error as the payload', () => {
        let result = ProtocActions.deleteBookOfTheMonthFailure(error);
        expect(result.payload).toEqual(error);
      })
    })
    describe('Single Entity Cancel', () => {
      it('should return no payload', () => {
        let result = ProtocActions.deleteBookOfTheMonthCancel();
        expect(result["payload"]).toEqual(undefined);
      })
    })
    describe('Repeated Entity Request', () => {
      it('should return the book as the payload', () => {
        let result = ProtocActions.deleteLibraryRequest(bookObj);
        expect(result.payload).toEqual(bookObj);
        expect(result.meta).toEqual({ resolve: undefined, reject: undefined });
      })
    })
    describe('Repeated Entity Request Promise', () => {
      it('should return the book as the payload', () => {
        let resolve = function(){};
        let reject = function(){};
        let result = ProtocActions.deleteLibraryRequestPromise(bookObj, resolve, reject);
        expect(result.payload).toEqual(bookObj);
        expect(result.meta).toEqual({ resolve, reject });
      })
    })
    describe('Repeated Entity Success', () => {
      it('should return the book as the payload', () => {
        let result = ProtocActions.deleteLibrarySuccess(bookObj);
        expect(result.payload).toEqual(bookObj);
      })
    })
    describe('Repeated Entity Failure', () => {
      it('should return the error as the payload', () => {
        let result = ProtocActions.deleteLibraryFailure(error);
        expect(result.payload).toEqual(error);
      })
    })
    describe('Repeated Entity Cancel', () => {
      it('should return no payload', () => {
        let result = ProtocActions.deleteLibraryCancel();
        expect(result["payload"]).toEqual(undefined);
      })
    })
  })

  describe('Single Entity Reset', () => {
    it('should return no payload', () => {
      let result = ProtocActions.resetBookOfTheMonth();
      expect(result["payload"]).toEqual(undefined);
    })
  })
  describe('Repeated Entity Reset', () => {
    it('should return no payload', () => {
      let result = ProtocActions.resetLibrary();
      expect(result["payload"]).toEqual(undefined);
    })
  })
})

describe('Custom Action Tests', () => {
  describe('Action Existance', () => {
    it('should generate Custom actions', () => {
      expect(typeof ProtocActions.customErrorBookRequest).toBe("function", "Request action not defined")
      expect(typeof ProtocActions.customErrorBookRequestPromise).toBe("function", "Request Promise action not defined")
      expect(typeof ProtocActions.customErrorBookSuccess).toBe("function", "Success action not defined")
      expect(typeof ProtocActions.customErrorBookFailure).toBe("function", "Failure action not defined")
      expect(typeof ProtocActions.customErrorBookCancel).toBe("function", "Cancel action not defined")
    })
  })
  describe('Action Type Tests', () => {
    it('should follow the same format for all the type strings', () => {
      let payload : string = "ErrorBook"
      let cruds : string[] = ["custom"]
      let effects : string[] = ["Request", "Success", "Failure", "Cancel"]
      for(var cIndex : number = 0; cIndex < cruds.length; cIndex++){
        for(var eIndex : number = 0; eIndex < effects.length; eIndex++){
          expect(ProtocActions[cruds[cIndex]+payload+effects[eIndex]]().type)
            .toBe("PROTOC_" + cruds[cIndex].toUpperCase() + "_" + payload.toUpperCase() + "_" + effects[eIndex].toUpperCase())
        }
      }
    })
  })
  describe('Custom Action Payload Tests', () => {
    let bookObj : ProtocTypes.readinglist.Book.AsObject;
    let error : NodeJS.ErrnoException;

    beforeAll(() => {
      let myBook = new ProtocTypes.readinglist.Book();
      myBook.setTitle("The Prince");
      myBook.setAuthor("Niccolo Machiavelli");
      bookObj = myBook.toObject();

      error = new Error("Test Error");
      error.code = "14";
    })

    it('should return the book as the payload', () => {
      let result = ProtocActions.customErrorBookRequest(bookObj);
      expect(result.payload).toEqual(bookObj);
      expect(result.meta).toEqual({ resolve: undefined, reject: undefined });
    })
    it('should return the book as the payload', () => {
      let resolve = function(){};
      let reject = function(){};
      let result = ProtocActions.customErrorBookRequestPromise(bookObj, resolve, reject);
      expect(result.payload).toEqual(bookObj);
      expect(result.meta).toEqual({ resolve, reject });
    })
    it('should return the book as the payload', () => {
      let result = ProtocActions.customErrorBookSuccess(bookObj);
      expect(result.payload).toEqual(bookObj);
    })
    it('should return the error as the payload', () => {
      let result = ProtocActions.customErrorBookFailure(error);
      expect(result.payload).toEqual(error);
    })
    it('should return no payload', () => {
      let result = ProtocActions.customErrorBookCancel();
      expect(result["payload"]).toEqual(undefined);
    })
  })
})

// TODO: test map type
// TODO: test enum type
// TODO: test oneof type
