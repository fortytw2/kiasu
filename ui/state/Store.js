import * as reducers from "./Reducers";

import { combineReducers, compose, createStore } from "redux";

let initialState = {
  notifications: [],
  login: { apiKey: "", email: "" }
};

let Store = createStore(
  combineReducers(reducers),
  initialState,
  window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__()
);

export { Store };
