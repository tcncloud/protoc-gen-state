import { createEpicMiddleware } from 'redux-observable';
import { composeWithDevTools } from 'redux-devtools-extension';
import { Store, createStore, applyMiddleware } from "redux";

import { RootReducer } from './rootReducer';
import { RootState, InitialState } from './rootState';
import { RootAction } from './rootAction';
import { RootEpic } from './rootEpic';


function configureStore(initialState?: RootState): Store<RootState> {
  // configure middleware
  const epicMiddleware = createEpicMiddleware<RootAction, RootAction, RootState>()

  const middlewares = [
    epicMiddleware
  ];

  // compose enhancers with dev tools
  const enhancer = composeWithDevTools(
    applyMiddleware(...middlewares)
  );

  // create store
  const store = createStore(
    RootReducer,
    initialState!,
    enhancer
  );

  epicMiddleware.run(RootEpic)

  // Hot reload reducers
  // if (module.hot) {
  //   module.hot.accept(() => {
  //     // TODO hot reload epics
  //     // https://github.com/reactjs/react-redux/issues/602
  //     // const nextEpic = require('./rootEpic').RootEpic;
  //     // Store.replaceMiddleware(nextEpic);
  //     const nextReducer = require('./rootReducer').RootReducer;
  //     Store.replaceReducer(nextReducer);
  //   })
  // }

  return store;
}

export const store = configureStore(InitialState);
