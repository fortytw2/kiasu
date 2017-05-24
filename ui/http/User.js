import {
  NOTIFICATION_LEVEL_INFO,
  addNotification,
} from "../state/Notifications";

import { Store } from "../state/Store";
import { ajax } from "nanoajax";

export function RequestLoginToken(email) {
  ajax(
    {
      url: "/api/token/request",
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        email: email
      })
    },
    function(code, responseText, request) {
      if (code !== 200) {
        alert(responseText);
        return;
      }
      console.log(JSON.parse(responseText));
      Store.dispatch(
        addNotification(NOTIFICATION_LEVEL_INFO, JSON.parse(responseText).note)
      );
    }
  );
}
