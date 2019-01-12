// package: readinglist
// file: e2e/redux4/protos/readinglist/readinglist.proto

var e2e_redux4_protos_readinglist_readinglist_pb = require("../../../../e2e/redux4/protos/readinglist/readinglist_pb");
var grpc = require("grpc-web-client").grpc;

var ReadingList = (function () {
  function ReadingList() {}
  ReadingList.serviceName = "readinglist.ReadingList";
  return ReadingList;
}());

ReadingList.ReadAllBooks = {
  methodName: "ReadAllBooks",
  service: ReadingList,
  requestStream: false,
  responseStream: true,
  requestType: e2e_redux4_protos_readinglist_readinglist_pb.Empty,
  responseType: e2e_redux4_protos_readinglist_readinglist_pb.Book
};

ReadingList.CreateBook = {
  methodName: "CreateBook",
  service: ReadingList,
  requestStream: false,
  responseStream: false,
  requestType: e2e_redux4_protos_readinglist_readinglist_pb.Book,
  responseType: e2e_redux4_protos_readinglist_readinglist_pb.Book
};

ReadingList.ReadBook = {
  methodName: "ReadBook",
  service: ReadingList,
  requestStream: false,
  responseStream: false,
  requestType: e2e_redux4_protos_readinglist_readinglist_pb.Book,
  responseType: e2e_redux4_protos_readinglist_readinglist_pb.Book
};

ReadingList.UpdateBook = {
  methodName: "UpdateBook",
  service: ReadingList,
  requestStream: false,
  responseStream: false,
  requestType: e2e_redux4_protos_readinglist_readinglist_pb.Book,
  responseType: e2e_redux4_protos_readinglist_readinglist_pb.Book
};

ReadingList.DeleteBook = {
  methodName: "DeleteBook",
  service: ReadingList,
  requestStream: false,
  responseStream: false,
  requestType: e2e_redux4_protos_readinglist_readinglist_pb.Book,
  responseType: e2e_redux4_protos_readinglist_readinglist_pb.Book
};

ReadingList.ErrorOut = {
  methodName: "ErrorOut",
  service: ReadingList,
  requestStream: false,
  responseStream: false,
  requestType: e2e_redux4_protos_readinglist_readinglist_pb.Book,
  responseType: e2e_redux4_protos_readinglist_readinglist_pb.Book
};

exports.ReadingList = ReadingList;

function ReadingListClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

ReadingListClient.prototype.readAllBooks = function readAllBooks(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(ReadingList.ReadAllBooks, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    onMessage: function (responseMessage) {
      listeners.data.forEach(function (handler) {
        handler(responseMessage);
      });
    },
    onEnd: function (status, statusMessage, trailers) {
      listeners.end.forEach(function (handler) {
        handler();
      });
      listeners.status.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners = null;
    }
  });
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

ReadingListClient.prototype.createBook = function createBook(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(ReadingList.CreateBook, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

ReadingListClient.prototype.readBook = function readBook(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(ReadingList.ReadBook, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

ReadingListClient.prototype.updateBook = function updateBook(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(ReadingList.UpdateBook, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

ReadingListClient.prototype.deleteBook = function deleteBook(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(ReadingList.DeleteBook, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

ReadingListClient.prototype.errorOut = function errorOut(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(ReadingList.ErrorOut, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

exports.ReadingListClient = ReadingListClient;

