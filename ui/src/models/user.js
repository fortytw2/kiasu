// src/models/User.js
import m from "mithril";
import config from "../config";
import flash from "./flash";
import raven from "raven-js";

export default {
  requestToken: function(email) {
    return m
      .request({
        method: "POST",
        url: config.URL + "/token/request",
        withCredentials: true,
        data: {
          email: email
        }
      })
      .then(
        function(result) {
          console.log(result);
          flash.flashMessage = result.note;
        },
        function(error) {
          raven.captureException(error);
        }
      );
  }
};
