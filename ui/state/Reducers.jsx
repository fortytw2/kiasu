import {
  SEND_LOGIN_TOKEN_SUCCESS,
  SEND_LOGIN_TOKEN_FAILURE
} from "./UserActions";

export function sendLoginToken(state = [], action) {
  switch (action.type) {
    case SEND_LOGIN_TOKEN_SUCCESS:
      return { text: action.text, success: true };
    case SEND_LOGIN_TOKEN_FAILURE:
      return { text: action.text, success: false };
    default:
      return state;
  }
}
