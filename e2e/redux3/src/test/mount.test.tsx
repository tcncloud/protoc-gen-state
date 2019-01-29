// components
import App from '@App/src/App';
import BookComponent from '@App/src/book/component';
import ServerTest from '@App/src/serverTest/component';
import PromiseTest from '@App/src/promiseTest/component';
import GrpcErrorCodeTest from '@App/src/grpcErrorCodeTest/component';
import TimeoutTest from '@App/src/timeoutTest/component';
import TimeoutRetryTest from '@App/src/timeoutRetryTest/component';

// general
import React from 'react';
import { shallow } from 'enzyme';

// import { Provider } from 'redux';
// import configureMockStore from 'redux-mock-store';
// const mockStore = configureMockStore();

describe('Mounting test to see if components can shallow render', () => {

  // it('Mount the Book component', () => {
  //   shallow(
  //     <Provider store={store} />
  //       <BookComponent />
  //     </Provider>
  //   );
  // })

  // it('Mount the ServerTest component', () => {
  //   shallow(<ServerTest  store={store} />);
  // })

  // it('Mount the PromiseTest component', () => {
  //   shallow(<PromiseTest />);
  // })

  // it('Mount the GrpcErrorCodeTest component', () => {
  //   shallow(<GrpcErrorCodeTest />);
  // })

  // it('Mount the TimeoutTest component', () => {
  //   shallow(<TimeoutTest />);
  // })

  // it('Mount the TimeoutRetryTest component', () => {
  //   shallow(<TimeoutRetryTest />);
  // })

  it('Mount the entire App component', () => {
    shallow(<App />);
  })

})
