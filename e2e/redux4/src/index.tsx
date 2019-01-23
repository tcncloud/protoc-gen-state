import * as React from "react";
import * as ReactDOM from "react-dom";
import { Provider } from "react-redux";
import 'rxjs';
// import 'core-js';
// import 'es6-shim';

import { store } from './store';
import App from "./App";
// import "./index.css";


ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById("root")
);
