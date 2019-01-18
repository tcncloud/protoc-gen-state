import _ from 'lodash';
import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators, Dispatch  } from 'redux';

import * as protocActions from 'protos/BasicState/actions_pb';
import * as ProtocServices from 'protos/BasicState/protoc_services_pb';
import * as ProtocTypes from 'protos/BasicState/protoc_types_pb';
import { initialProtocState, ProtocState } from 'protos/BasicState/state_pb';
import { RootState } from '../rootState';

interface IState {
  bookOfTheMonth: ProtocTypes.readinglist.Book.AsObject;
}
interface IDispatch {
  customErrorBook: typeof protocActions.customErrorBookRequest;
  customErrorBookPromise: typeof protocActions.customErrorBookRequestPromise;
}
interface IProps extends IState, IDispatch {}

class GrpcErrorCodeTest extends React.Component<IProps, {}> {
  constructor(props: IProps) {
    super(props);
  }

  public getErrorPromise = (request: ProtocTypes.readinglist.Book.AsObject) => {
    return new Promise((resolve, reject) => {
      this.props.customErrorBookPromise(request, resolve, reject);
    })
  }

  public makeErrorAndSendPromise = () => {
    let myBook = new ProtocTypes.readinglist.Book();
    myBook.setAuthor("George Orwell");
    myBook.setTitle("1942 churchill sailed the ocean blue");

    this.getErrorPromise(myBook.toObject())
      .then((res) => console.log('bingo', res))
      .catch((err) => console.log('badgo', err))
  }

  public makeErrorAndSend = () => {
    let myBook = new ProtocTypes.readinglist.Book();
    myBook.setAuthor("George Orwell");
    myBook.setTitle("Animal House");

    this.props.customErrorBook(myBook.toObject())
  }

  public render() {
    console.log('props: ', this.props);

    return (
      <div style={{ marginLeft:'100px' }}>
        gRPC error codes:
        <button onClick={this.makeErrorAndSendPromise}>GetErrorPromise</button>
        <button onClick={this.makeErrorAndSend}>GetError</button>
      </div>
    );
  }
}

function mapStateToProps(state: RootState): IState {
  return {
    bookOfTheMonth: state.protoc.bookOfTheMonth.value,
  };
}

function mapDispatchToProps(dispatch: Dispatch<RootState>): IDispatch {
  return bindActionCreators({
    customErrorBookPromise: protocActions.customErrorBookRequestPromise,
    customErrorBook: protocActions.customErrorBookRequest,
  }, dispatch);
}
export default _.flowRight([
  connect<IState, IDispatch, {}>(mapStateToProps, mapDispatchToProps),
])(GrpcErrorCodeTest);

