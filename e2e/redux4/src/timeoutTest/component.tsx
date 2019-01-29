import _ from 'lodash';
import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators, Dispatch  } from 'redux';

import * as protocActions from 'protos/BasicState/actions_pb';
import * as ProtocServices from 'protos/BasicState/protoc_services_pb';
import * as ProtocTypes from 'protos/BasicState/protoc_types_pb';
import { initialProtocState } from 'protos/BasicState/state_pb';
import { RootAction } from '../rootAction';

console.log('----------- PROTOC REDUX ------------');
console.log('initialProtocState: ', initialProtocState);
console.log('protocActions: ', protocActions);
console.log('ProtocTypes: ', ProtocTypes);
console.log('ProtocServices: ', ProtocServices);
console.log('--------------------------------------');

interface IState { }

interface IDispatch {
  getTimeoutPromise: typeof protocActions.getTimeoutBookRequestPromise;
}
interface IProps extends IState, IDispatch {}

class TimeoutTest extends React.Component<IProps, {}> {
  constructor(props: IProps) {
    super(props);
  }

  public timeoutPromise = (request: ProtocTypes.readinglist.Book.AsObject) => {
    return new Promise((resolve, reject) => {
      this.props.getTimeoutPromise(request, resolve, reject);
    })
  }

  public doTimeout = () => {
    let myBook = new ProtocTypes.readinglist.Book();
    myBook.setTitle("George Orwell");

    this.timeoutPromise(myBook.toObject())
      .then((res) => console.log('promise resolve -- ', res))
      .catch((err) => console.log('promise reject (timeout) -- ', err))
  }

  public render() {
    console.log('props: ', this.props);

    return (
      <div style={{ marginLeft: "100px" }}>
        Timeout Action:
        <button onClick={this.doTimeout}>GetTimeout</button>
      </div>
    );
  }
}

function mapStateToProps(): IState {
  return { };
}

function mapDispatchToProps(dispatch: Dispatch<RootAction>): IDispatch {
  return bindActionCreators({
    getTimeoutPromise: protocActions.getTimeoutBookRequestPromise,
  }, dispatch);
}
export default _.flowRight([
  connect<IState, IDispatch, {}>(mapStateToProps, mapDispatchToProps),
])(TimeoutTest);

