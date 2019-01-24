import * as React from "react";

import "./App.css";
import BookComponent from './book/component';
import ServerTest from './serverTest/component';
import PromiseTest from './promiseTest/component';
import GrpcErrorCodeTest from './grpcErrorCodeTest/component';
import TimeoutTest from './timeoutTest/component';
import TimeoutRetryTest from './timeoutRetryTest/component';

import { RootState } from './rootState';



export default class App extends React.Component<any,RootState> {
  render() {
    return (
      <div>
        <BookComponent />

        <ServerTest />

        <PromiseTest />

        <GrpcErrorCodeTest />

        <TimeoutTest />

        <TimeoutRetryTest />
      </div>
    )
  }
};
