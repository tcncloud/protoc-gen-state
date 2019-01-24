import { protocReducer } from 'protos/BasicState/reducer_pb';
import { initialProtocState } from 'protos/BasicState/state_pb';
import * as ProtocActions from 'protos/BasicState/actions_pb';
import * as ProtocTypes from 'protos/BasicState/protoc_types_pb';

describe('General Action Handling', () => {
  it('should return the initial state', () => {
    // can't send an empty object as action, so just using a junk one
    expect(protocReducer(undefined, ProtocActions.resetBookOfTheMonth())).toEqual(initialProtocState)
  })

  describe('Request Actions', () => {
    let book = new ProtocTypes.readinglist.Book();
    book.setTitle("Great Expectations");
    book.setAuthor("Charles Dickens");

    it('should mark isLoading as true', () => {
      expect(protocReducer(undefined, ProtocActions.createBookOfTheMonthRequest(book.toObject()))["bookOfTheMonth"]["isLoading"]).toEqual(true);
    })
  })

  describe('Success Actions', () => {
    let book = new ProtocTypes.readinglist.Book();
    book.setTitle("The Fountainhead");
    book.setAuthor("Ayn Rand");
    let errorStateSingle = protocReducer(undefined, ProtocActions.createBookOfTheMonthFailure(new Error("test")));
    let bookArray : ProtocTypes.readinglist.Book.AsObject[] = [];
    bookArray.push(book.toObject());

    it('should mark isLoading as false', () => {
      expect(protocReducer(undefined, ProtocActions.createBookOfTheMonthSuccess(book.toObject()))["bookOfTheMonth"]["isLoading"]).toEqual(false);
    })
    it('should set the value of a single field to the book', () => {
      expect(protocReducer(undefined, ProtocActions.createBookOfTheMonthSuccess(book.toObject()))["bookOfTheMonth"]["value"]).toEqual(book.toObject());
    })
    it('should set the value of a repeated field to the array', () => {
      expect(protocReducer(undefined, ProtocActions.createLibrarySuccess(book.toObject()))["library"]["value"]).toEqual(bookArray);
    })
    it('should clear the error if it exists', () => {
      expect(protocReducer(errorStateSingle, ProtocActions.createBookOfTheMonthSuccess(book.toObject()))["bookOfTheMonth"]["error"]).toEqual(null);
    })
  })

  describe('Failure Actions', () => {
    let error : NodeJS.ErrnoException = new Error("bigtime failure");
    error.code = "10";
    let book = new ProtocTypes.readinglist.Book();
    book.setTitle("The Fountainhead");
    book.setAuthor("Ayn Rand");
    let successStateSingle = protocReducer(undefined, ProtocActions.createBookOfTheMonthSuccess(book.toObject()));
    let successStateRepeated = protocReducer(undefined, ProtocActions.createLibrarySuccess(book.toObject()));
    let requestStateSingle = protocReducer(undefined, ProtocActions.createBookOfTheMonthRequest(book.toObject()));
    let bookArray : ProtocTypes.readinglist.Book.AsObject[] = [];
    bookArray.push(book.toObject());

    it('should mark isLoading as false', () => {
      expect(protocReducer(
        requestStateSingle,
        ProtocActions.createBookOfTheMonthFailure(error)
      )["bookOfTheMonth"]["isLoading"]).toEqual(false);
    })
    it('should set the error in the state', () => {
      expect(protocReducer(undefined, ProtocActions.createBookOfTheMonthFailure(error))["bookOfTheMonth"]["error"]).toEqual({
        message: error.message,
        code: error.code
      })
    })
    it('should reset a single value to null', () => {
      expect(protocReducer(
        successStateSingle,
        ProtocActions.createBookOfTheMonthFailure(error)
      )["bookOfTheMonth"]["value"]).toEqual(book.toObject())
    })
    it('should leave the value intact', () => {
      expect(protocReducer(
        successStateRepeated,
        ProtocActions.createLibraryFailure(error)
      )["library"]["value"]).toEqual(bookArray)
    })
  })

  describe('Cancel Actions', () => {
    let book = new ProtocTypes.readinglist.Book();
    book.setTitle("Ulysses");
    book.setAuthor("James Joyce");

    it('should mark isLoading as false', () => {
      let initState = protocReducer(undefined, ProtocActions.createBookOfTheMonthRequest(book.toObject()));
      expect(protocReducer(
        initState,
        ProtocActions.createBookOfTheMonthCancel()
      )["bookOfTheMonth"]["isLoading"]).toEqual(false)
    })
  })

  describe('Reset Actions', () => {
    describe('Single Entity', () => {
      let book = new ProtocTypes.readinglist.Book();
      book.setTitle("Don Quixote");
      book.setAuthor("Miguel de Cervantes");
      let error = new Error("bad");
      let loadingInitStateSingle = protocReducer(undefined, ProtocActions.createBookOfTheMonthRequest(book.toObject()));
      let successInitStateSingle = protocReducer(undefined, ProtocActions.createBookOfTheMonthSuccess(book.toObject()));
      let errorInitStateSingle = protocReducer(undefined, ProtocActions.createBookOfTheMonthFailure(error));

      it('should reset value to null', () => {
        expect(
          protocReducer(
            successInitStateSingle,
            ProtocActions.resetBookOfTheMonth()
          )["bookOfTheMonth"]["value"]
        ).toEqual(null);
      })
      it('should reset error to null', () => {
        expect(
          protocReducer(
            errorInitStateSingle,
            ProtocActions.resetBookOfTheMonth()
          )["bookOfTheMonth"]["error"]
        ).toEqual(null);
      })
      it('should reset isLoading state to false', () => {
        expect(
          protocReducer(
            loadingInitStateSingle,
            ProtocActions.resetBookOfTheMonth()
          )["bookOfTheMonth"]["isLoading"]
        ).toEqual(false);
      })
    })

    describe('Repeated Entity', () => {
      let book = new ProtocTypes.readinglist.Book();
      book.setTitle("Moby Dick");
      book.setAuthor("Herman Melville");
      let error = new Error("bad");
      let loadingInitStateRepeated = protocReducer(undefined, ProtocActions.createLibraryRequest(book.toObject()));
      let successInitStateRepeated = protocReducer(undefined, ProtocActions.createLibrarySuccess(book.toObject()));
      let errorInitStateRepeated = protocReducer(undefined, ProtocActions.createLibraryFailure(error));

      it('should reset value to null', () => {
        expect(
          protocReducer(
            successInitStateRepeated,
            ProtocActions.resetLibrary()
          )["library"]["value"]
        ).toEqual([]);
      })
      it('should reset error to null', () => {
        expect(
          protocReducer(
            errorInitStateRepeated,
            ProtocActions.resetLibrary()
          )["library"]["error"]
        ).toEqual(null);
      })
      it('should reset isLoading state to false', () => {
        expect(
          protocReducer(
            loadingInitStateRepeated,
            ProtocActions.resetLibrary()
          )["library"]["isLoading"]
        ).toEqual(false);
      })
    })
  })
})

describe('Single Entity CLUDG Actions', () => {
  let book = new ProtocTypes.readinglist.Book();
  book.setTitle("Crime and Punishment");
  book.setAuthor("Fyodor Dostoevsky");
  let updatedBook = new ProtocTypes.readinglist.Book();
  book.setTitle("The Brothers Karamazov");
  book.setAuthor("Fyodor Dostoevsky");
  let initState = protocReducer(undefined, ProtocActions.createBookOfTheMonthSuccess(book.toObject()))

  describe('create Actions', () => {
    it('should store the entity under the value key', () => {
      expect(protocReducer(
        undefined,
        ProtocActions.createBookOfTheMonthSuccess(book.toObject())
      )["bookOfTheMonth"]["value"]).toEqual(book.toObject())
    })
  })
  describe('Update Actions', () => {
    it('should replace the entity under the value key', () => {
      expect(protocReducer(
        initState,
        ProtocActions.updateBookOfTheMonthSuccess(updatedBook.toObject())
      )["bookOfTheMonth"]["value"]).toEqual(updatedBook.toObject())
    })
  })
  describe('Delete Actions', () => {
    it('should reset the value key to the default value', () => {
      expect(protocReducer(
        initState,
        ProtocActions.deleteBookOfTheMonthSuccess(book.toObject())
      )["bookOfTheMonth"]["value"]).toEqual(initialProtocState.bookOfTheMonth.value)
    })
  })
  describe('Get Actions', () => {
    it('should replace the entity under the value key', () => {
      expect(protocReducer(
        initState,
        ProtocActions.getBookOfTheMonthSuccess(updatedBook.toObject())
      )["bookOfTheMonth"]["value"]).toEqual(updatedBook.toObject())
    })
  })
  // List Actions should not exist, error will be thrown during generation
})

describe('Repeated Entity CLUDG Actions', () => {
  let book = new ProtocTypes.readinglist.Book();
  book.setTitle("Forever");
  book.setAuthor("Judy Blume");
  let updatedBook = new ProtocTypes.readinglist.Book();
  book.setTitle("Forever");
  book.setAuthor("Pete Hamill");
  let initState = protocReducer(undefined, ProtocActions.createBookOfTheMonthSuccess(book.toObject()))
  let bookArray : ProtocTypes.readinglist.Book.AsObject[] = [];
  bookArray.push(book.toObject());

  describe('Create Actions', () => {
    it('should append the entity to the array', () => {
      expect(protocReducer(
        undefined,
        ProtocActions.createLibrarySuccess(book.toObject())
      )["library"]["value"]).toEqual(bookArray)
    })
  })
  describe('Update Actions', () => {
    it('should update in place the previous entity with the updated entity', () => {
      expect(protocReducer(
        initState,
        ProtocActions.updateLibrarySuccess({ prev: book.toObject(), updated: updatedBook.toObject() })
      )["library"]["value"]).toEqual([updatedBook.toObject()])
    })
  })
  describe('Delete Actions', () => {
    it('should remove the entity from the array', () => {
      expect(protocReducer(
        initState,
        ProtocActions.deleteLibrarySuccess(book.toObject())
      )["library"]["value"]).toEqual([])
    })
  })
  describe('List Actions', () => {
    it('should replace the entire array', () => {
      expect(protocReducer(
        initState,
        ProtocActions.listLibrarySuccess([updatedBook.toObject()])
      )["library"]["value"]).toEqual([updatedBook.toObject()])
    })
  })
  // Get Actions should not exist, error will be thrown during generation
})

// TODO: test map type
// TODO: test enum type
// TODO: test oneof type
