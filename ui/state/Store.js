import * as reducers from "./Reducers";

import { autoRehydrate, persistStore } from "redux-persist";
import { combineReducers, compose, createStore } from "redux";

let initialState = {
  notifications: []
};

let Store = createStore(
  combineReducers(reducers),
  initialState,
  window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__(),
  compose(autoRehydrate())
);

persistStore(Store, { blacklist: ["notifications"] }, () => {
  console.log("rehydrated");
});

export { Store };
