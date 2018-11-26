import * as React from "react";

import "./App.css";
import Dog from './dog/component';
import Graph from './graph/component';
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
        <Dog parentPropsExample={"a prop"}/>
        <br />

        <Graph />

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
