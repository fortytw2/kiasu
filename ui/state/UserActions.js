export const SEND_LOGIN_TOKEN_SUCCESS = "SEND_LOGIN_TOKEN_SUCCESS";

export function sendLoginTokenSuccess(apiKey) {
  return { type: SEND_LOGIN_TOKEN_SUCCESS, apiKey: apiKey };
}
