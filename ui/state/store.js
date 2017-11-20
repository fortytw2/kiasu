import { applyMiddleware, combineReducers, compose, createStore } from "redux";
import { persistCombineReducers, persistStore } from "redux-persist";
import storage from 'redux-persist/es/storage' 
import { routerReducer, syncHistoryWithStore } from "preact-router-redux";

import createBrowserHistory from "history/createBrowserHistory";
import { login } from "./login/reducers";
import { notifications } from "./notifications/reducers";

const config = {
  key: 'hc-root',
  storage,
}

let Store = createStore(
  persistCombineReducers(config, {
    login: login,
    notifications: notifications,
    routing: routerReducer
  }),
  window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__()
);

// Create an enhanced history that syncs navigation events with the store
const History = syncHistoryWithStore(createBrowserHistory(), Store);

export { Store, History };
