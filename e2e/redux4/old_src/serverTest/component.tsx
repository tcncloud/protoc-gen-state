import * as React from "react";

import { grpc } from "grpc-web-client";
import * as ProtocTypes from 'protos/BasicState/protoc_types_pb';
import * as ProtocActions from 'protos/BasicState/actions_pb';
import * as ProtocServices from 'protos/BasicState/protoc_services_pb'
console.log('Check services', ProtocServices)

console.log("acionts: ",ProtocActions);
console.log("types: ",ProtocTypes);


interface PropsFromComponent {}

interface ComponentLocalState {
  readonly localStateExample: string,
}

class ServerTest extends React.Component<PropsFromComponent | ComponentLocalState> {
  state: ComponentLocalState;
  constructor(props: PropsFromComponent) {
    super(props);

    this.state = {
      localStateExample: "money",
    };

    this.createBook = this.createBook.bind(this)
    this.deleteBook = this.deleteBook.bind(this)
  }

  componentWillReceiveProps() {
  }

  // deleteBook(e: React.FormEvent<HTMLButtonElement>): void {
  deleteBook(): void {
    const host = "http://localhost:9090";
    const removal = new ProtocTypes.readinglist.Book();
    removal.setTitle("To Kill a Mockingbird")
    removal.setAuthor("Harper Lee")
    const client = grpc.client(ProtocServices.readinglist.ReadingList.DeleteBook, {
      host: host,
    });
    client.onHeaders((headers: grpc.Metadata) => {
      console.log("queryBooks.onHeaders", headers);
    });
    client.onMessage((message: ProtocTypes.readinglist.Book) => {
      console.log("queryBooks.onMessage", message.toObject());
    });
    client.onEnd((code: grpc.Code, msg: string, trailers: grpc.Metadata) => {
      console.log("queryBooks.onEnd", code, msg, trailers);
    });
    client.start(new grpc.Metadata({ "authorization": `Bearer ASDFQWEQWERQGaseqwr3qrwe`}));
    client.send(removal);
  }

  // createBook(e: React.FormEvent<HTMLButtonElement>): void {
  createBook(): void {
    const host = "http://localhost:9090";
    const magazine = new ProtocTypes.readinglist.Book();
    magazine.setTitle("To Kill a Mockingbird")
    magazine.setAuthor("Harper Lee")
    console.log(magazine)
    grpc.unary(ProtocServices.readinglist.ReadingList.CreateBook, {
      request: magazine,
      host: host,
      metadata: new grpc.Metadata({ "authorization": `Bearer ASDFQWEQWERQGaseqwr3qrwe`}),
      onEnd: (res:any) => {
        console.log("onEnd", res);
      }
    })
  }

  render() {
    return (
      <div className="App" style={{"marginTop":"20px", paddingTop: '30px', "background": "#f1f1f1", "paddingBottom":"50px"}}>
        <button onClick={this.createBook}>Add Book Grpc Call</button>
        <button onClick={this.deleteBook}>Delete Book Grpc Call</button>
      </div>
    );
  }
}

export default ServerTest;
