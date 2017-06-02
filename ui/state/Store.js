import * as reducers from "./Reducers";

import { applyMiddleware, combineReducers, compose, createStore } from "redux";
import { autoRehydrate, persistStore } from "redux-persist";

let initialState = {
  notifications: [],
  login: { apiKey: "", email: "" }
};

let Store = compose(autoRehydrate({ log: true }))(createStore)(
  combineReducers(reducers),
  window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__()
);

export { Store };
