import _ from 'lodash';
import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators, Dispatch  } from 'redux';

import * as ProtocTypes from 'protos/BasicState/protoc_types_pb';
import { RootState } from '../rootState';

import * as actions from './actions';

interface IState { }

interface IDispatch {
  retryRequestPromise: typeof actions.retryRequestPromise;
  timeoutRequestPromise: typeof actions.timeoutRequestPromise;
  codeRequestPromise: typeof actions.codeRequestPromise;
}
interface IProps extends IState, IDispatch {}

class TimeoutRetryTest extends React.Component<IProps, {}> {
  constructor(props: IProps) {
    super(props);
  }

  public timeoutPromise = (request: ProtocTypes.readinglist.Book.AsObject) => {
    return new Promise((resolve, reject) => {
      this.props.timeoutRequestPromise(request, resolve, reject);
    })
  }

  public retryPromise = (request: ProtocTypes.readinglist.Book.AsObject) => {
    return new Promise((resolve, reject) => {
      this.props.retryRequestPromise(request, resolve, reject);
    })
  }

  public codePromise = (request: ProtocTypes.readinglist.Book.AsObject) => {
    return new Promise((resolve, reject) => {
      this.props.codeRequestPromise(request, resolve, reject);
    })
  }

  public sendBookCode = () => {
    let myBook = new ProtocTypes.readinglist.Book();
    myBook.setAuthor("George Orwell");
    myBook.setAuthor("1942 churchill sailed the ocean blue");

    this.codePromise(myBook.toObject())
      .then((res) => console.log('code resolve -- ', res))
      .catch((err) => console.log('code reject -- ', err))
  }

  public sendBookTimeout = () => {
    let myBook = new ProtocTypes.readinglist.Book();
    myBook.setAuthor("George Orwell");

    this.timeoutPromise(myBook.toObject())
      .then((res) => console.log('promise resolve -- ', res))
      .catch((err) => console.log('promise reject -- ', err))
  }

  public sendBookRetry = () => {
    let myBook = new ProtocTypes.readinglist.Book();
    myBook.setAuthor("George Orwell");
    myBook.setTitle("Animal Farm");

    this.retryPromise(myBook.toObject())
      .then((res) => console.log('promise resolve -- ', res))
      .catch((err) => console.log('promise reject -- ', err))
  }

  public render() {
    console.log('props: ', this.props);

    return (
      <div style={{ margin:'100px', padding:'100px' }}>
        "Timeout: "
        <button onClick={this.sendBookTimeout}>GetOrg</button>
        <br />

        "Retry: "
        <button onClick={this.sendBookRetry}>GetOrg</button>
        <br />

        "Code: "
        <button onClick={this.sendBookCode}>GetError</button>
      </div>
    );
  }
}

function mapStateToProps(): IState {
  return { };
}

function mapDispatchToProps(dispatch: Dispatch<RootState>): IDispatch {
  return bindActionCreators({
    retryRequestPromise: actions.retryRequestPromise,
    timeoutRequestPromise: actions.timeoutRequestPromise,
    codeRequestPromise: actions.codeRequestPromise,
  }, dispatch);
}
export default _.flowRight([
  connect<IState, IDispatch, {}>(mapStateToProps, mapDispatchToProps),
])(TimeoutRetryTest);

