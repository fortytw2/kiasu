import { NOTIFICATION_ADD, NOTIFICATION_REMOVE } from "./Notifications";

import { SEND_LOGIN_TOKEN_SUCCESS } from "./UserActions";
import _ from "lodash";

export function apiKey(state = "", action) {
  switch (action.type) {
    case SEND_LOGIN_TOKEN_SUCCESS:
      return action.apiKey;
    default:
      return state;
  }
}

export function notifications(state = [], action) {
  switch (action.type) {
    case NOTIFICATION_ADD:
      return _.concat(state, [
        {
          message: action.message,
          level: action.level
        }
      ]);
    default:
      return state;
  }
}
