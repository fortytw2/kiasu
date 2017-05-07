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
      .then(function(result) {
        console.log(result);
        flash.flashMessage = result.note;
      })
      .catch(function(error) {
        raven.captureException(error);
      });
  },
  activateToken: function(token) {
    return m
      .request({
        method: "POST",
        url: config.URL + "/token/activate",
        withCredentials: true,
        data: {
          token: token
        }
      })
      .then(function(result) {
        window.localStorage.setItem("hydrocarbon-key", result.key);
        window.localStorage.setItem("hydrocarbon-email", result.email);

        flash.flashMessage = "logged in ok";
      });
  },
  logout: function() {
    return m
      .request({
        method: "POST",
        url: config.URL + "/key/deactivate",
        withCredentials: true,
        data: {
          key: window.localStorage.getItem("hydrocarbon-key")
        }
      })
      .then(function(result) {
        window.localStorage.removeItem("hydrocarbon-key");
        window.localStorage.removeItem("hydrocarbon-email");

        flash.flashMessage = "logged out";
      });
  },
  loggedIn: function() {
    if (window.localStorage.getItem("hydrocarbon-key") !== null) {
      return true;
    }
    return false;
  },
  loggedInUser: function() {
    if (window.localStorage.getItem("hydrocarbon-key") !== null) {
      return window.localStorage.getItem("hydrocarbon-email");
    }
    return "";
  }
};
