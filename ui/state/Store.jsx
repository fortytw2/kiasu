import { combineReducers } from "redux";
import * as reducers from "./Reducers";
import { createStore } from "redux";

let store = createStore(
  combineReducers(reducers),
  window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__()
);

export default store;
