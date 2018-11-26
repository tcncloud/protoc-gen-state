import _ from 'lodash';
import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators, Dispatch  } from 'redux';

import * as ProtocActions from 'proto/BasicState/actions_pb';
import * as ProtocTypes from 'proto/BasicState/protoc_types_pb';
import * as ProtocServices from 'proto/BasicState/protoc_services_pb';
import { initialProtocState, ProtocState } from 'proto/BasicState/state_pb';

import { RootState } from '../rootState';

console.log('----------- PROTOC REDUX ------------');
console.log('initialProtocState: ', initialProtocState);
console.log('ProtocActions: ', ProtocActions);
console.log('ProtocTypes: ', ProtocTypes);
console.log('ProtocServices: ', ProtocServices);
console.log('--------------------------------------');

interface IState {
  bookOfTheMonth: ProtocTypes.readinglist.Book.AsObject;
}

interface IDispatch {
  getBOTMRequestPromise: typeof ProtocActions.getBookOfTheMonthRequestPromise;
}
interface IProps extends IState, IDispatch {}

class PromiseTest extends React.Component<IProps, {}> {
  constructor(props: IProps) {
    super(props);
  }

  public getBOTMPromise = (request: ProtocTypes.readinglist.Book.AsObject) => {
    return new Promise((resolve, reject) => {
      this.props.getBOTMRequestPromise(request, resolve, reject);
    })
  }

  public makeBookAndSend = () => {
    let myBook = new ProtocTypes.readinglist.Book();
    myBook.setTitle("George Orwell");

    this.getBOTMPromise(myBook.toObject())
      .then((res) => console.log('promise resolve -- ', res))
      .catch((err) => console.log('promise reject -- ', err))
  }

  public render() {
    console.log('props: ', this.props);

    return (
      <div style={{ marginLeft:'100px' }}>
        Promisified Action:
        <button onClick={this.makeBookAndSend}>GetBook</button> (George Orwell)
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
    getBOTMRequestPromise: ProtocActions.getBookOfTheMonthRequestPromise,
  }, dispatch);
}
export default _.flowRight([
  connect<IState, IDispatch, {}>(mapStateToProps, mapDispatchToProps),
])(PromiseTest);

