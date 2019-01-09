// package: readinglist
// file: e2e/protos/readinglist/readinglist.proto

import * as e2e_protos_readinglist_readinglist_pb from "../../../e2e/protos/readinglist/readinglist_pb";
import {grpc} from "grpc-web-client";

type ReadingListReadAllBooks = {
  readonly methodName: string;
  readonly service: typeof ReadingList;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof e2e_protos_readinglist_readinglist_pb.Empty;
  readonly responseType: typeof e2e_protos_readinglist_readinglist_pb.Book;
};

type ReadingListCreateBook = {
  readonly methodName: string;
  readonly service: typeof ReadingList;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof e2e_protos_readinglist_readinglist_pb.Book;
  readonly responseType: typeof e2e_protos_readinglist_readinglist_pb.Book;
};

type ReadingListReadBook = {
  readonly methodName: string;
  readonly service: typeof ReadingList;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof e2e_protos_readinglist_readinglist_pb.Book;
  readonly responseType: typeof e2e_protos_readinglist_readinglist_pb.Book;
};

type ReadingListUpdateBook = {
  readonly methodName: string;
  readonly service: typeof ReadingList;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof e2e_protos_readinglist_readinglist_pb.Book;
  readonly responseType: typeof e2e_protos_readinglist_readinglist_pb.Book;
};

type ReadingListDeleteBook = {
  readonly methodName: string;
  readonly service: typeof ReadingList;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof e2e_protos_readinglist_readinglist_pb.Book;
  readonly responseType: typeof e2e_protos_readinglist_readinglist_pb.Book;
};

type ReadingListErrorOut = {
  readonly methodName: string;
  readonly service: typeof ReadingList;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof e2e_protos_readinglist_readinglist_pb.Book;
  readonly responseType: typeof e2e_protos_readinglist_readinglist_pb.Book;
};

export class ReadingList {
  static readonly serviceName: string;
  static readonly ReadAllBooks: ReadingListReadAllBooks;
  static readonly CreateBook: ReadingListCreateBook;
  static readonly ReadBook: ReadingListReadBook;
  static readonly UpdateBook: ReadingListUpdateBook;
  static readonly DeleteBook: ReadingListDeleteBook;
  static readonly ErrorOut: ReadingListErrorOut;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }
export type ServiceClientOptions = { transport: grpc.TransportConstructor }

interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: () => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}

export class ReadingListClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: ServiceClientOptions);
  readAllBooks(requestMessage: e2e_protos_readinglist_readinglist_pb.Empty, metadata?: grpc.Metadata): ResponseStream<e2e_protos_readinglist_readinglist_pb.Book>;
  createBook(
    requestMessage: e2e_protos_readinglist_readinglist_pb.Book,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: e2e_protos_readinglist_readinglist_pb.Book|null) => void
  ): void;
  createBook(
    requestMessage: e2e_protos_readinglist_readinglist_pb.Book,
    callback: (error: ServiceError, responseMessage: e2e_protos_readinglist_readinglist_pb.Book|null) => void
  ): void;
  readBook(
    requestMessage: e2e_protos_readinglist_readinglist_pb.Book,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: e2e_protos_readinglist_readinglist_pb.Book|null) => void
  ): void;
  readBook(
    requestMessage: e2e_protos_readinglist_readinglist_pb.Book,
    callback: (error: ServiceError, responseMessage: e2e_protos_readinglist_readinglist_pb.Book|null) => void
  ): void;
  updateBook(
    requestMessage: e2e_protos_readinglist_readinglist_pb.Book,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: e2e_protos_readinglist_readinglist_pb.Book|null) => void
  ): void;
  updateBook(
    requestMessage: e2e_protos_readinglist_readinglist_pb.Book,
    callback: (error: ServiceError, responseMessage: e2e_protos_readinglist_readinglist_pb.Book|null) => void
  ): void;
  deleteBook(
    requestMessage: e2e_protos_readinglist_readinglist_pb.Book,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: e2e_protos_readinglist_readinglist_pb.Book|null) => void
  ): void;
  deleteBook(
    requestMessage: e2e_protos_readinglist_readinglist_pb.Book,
    callback: (error: ServiceError, responseMessage: e2e_protos_readinglist_readinglist_pb.Book|null) => void
  ): void;
  errorOut(
    requestMessage: e2e_protos_readinglist_readinglist_pb.Book,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: e2e_protos_readinglist_readinglist_pb.Book|null) => void
  ): void;
  errorOut(
    requestMessage: e2e_protos_readinglist_readinglist_pb.Book,
    callback: (error: ServiceError, responseMessage: e2e_protos_readinglist_readinglist_pb.Book|null) => void
  ): void;
}

