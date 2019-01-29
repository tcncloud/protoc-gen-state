import { ActionsObservable, StateObservable } from 'redux-observable'
import { grpc } from '@improbable-eng/grpc-web'
import { Subject } from 'rxjs'

import * as ProtocEpics from 'protos/BasicState/epics_pb'
import * as ProtocActions from 'protos/BasicState/actions_pb'
import * as ProtocTypes from 'protos/BasicState/protoc_types_pb'

var api = {
  unary: (service:any, { request, host, metadata, onEnd }:any) => {
    console.log('******')
    console.log('in unary', request)
    onEnd({message: request, status: grpc.Code.OK})
  },
  client: grpc.client,
  // client: () => {
  //   return {
  //     // TODO look into these events
  //     onMessage: () => {
  //     },
  //     onEnd: (code, msg, trailers) => {
  //     },
  //     start: (metadata) => {
  //     },
  //     send: (pb_message) => {
  //     }
  // }
}

describe('Epic tests', () => {
  let bookArr : ProtocTypes.readinglist.Book.AsObject[] = []
  let bookObj : ProtocTypes.readinglist.Book.AsObject
  let error : NodeJS.ErrnoException
  // let store

  // TODO use this method
  // beforeEach(() => {
  //   const epicMiddleware = createEpicMiddleware()
  //   const mockStore = configureStore([epicMiddleware])
  //   store = mockStore(() => ({
  //     config: { host 'http://localhost:9090' } 
  //   }))

  //   epicMiddleware.run(j
  // })

  beforeAll(() => {
    let myBook = new ProtocTypes.readinglist.Book()
    myBook.setTitle("The Prince")
    myBook.setAuthor("Niccolo Machiavelli")
    bookObj = myBook.toObject()

    bookArr.push(bookObj)

    error = new Error("Test Error")
    error.code = "14"
  })

  xit('should test calling an epic', (done: any) => {
    const action$ = ActionsObservable.of(ProtocActions.getBookOfTheMonthRequest(bookObj))

    const state$ = new StateObservable(new Subject(), { config: { token: 'idklol', host: 'http://localhost:9090' } })

    const epic$ = ProtocEpics.getBookOfTheMonthEpic(action$, state$, api)

    epic$.subscribe(result => {
      // This will only work if the repeat() operator is taken off of the epics
      expect(result).toEqual(ProtocActions.getBookOfTheMonthSuccess(bookObj))
      done()
    })

  })

  it('should dispatch actions and watch the payload of the returned actions by using the store', () => {
    // TODO
    expect(true).toBe(true)
  })

})

