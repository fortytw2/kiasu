// App entry point

import m from "mithril";
import layout from "./components/layout";
import login from "./components/login";
import logout from "./components/logout";
import logincallback from "./components/login-callback";
import settings from "./components/settings";
import sessions from "./components/sessions";
import config from "./config";
import raven from "raven-js";

m.route.prefix("");

m.route(document.body, "/", {
  "/": {
    render: function() {
      return m(layout, m("p", "this is the homepage"));
    }
  },
  "/about": {
    render: function() {
      return m(layout, m("p", "this is the about page"));
    }
  },
  "/privacy-policy": {
    render: function() {
      return m(layout, m("p", config.PRIVACY));
    }
  },
  "/terms-and-conditions": {
    render: function() {
      return m(layout, m("p", config.TERMS));
    }
  },
  "/promise": {
    render: function() {
      return m(layout, m("p", "this is the hydrocarbon promise"));
    }
  },
  "/login": {
    render: function() {
      return m(layout, m(login));
    }
  },
  "/logout": {
    render: function() {
      return m(layout, m(logout));
    }
  },
  "/login-callback": {
    render: function() {
      return m(layout, m(logincallback));
    }
  },
  "/settings": {
    render: function() {
      return m(layout, m(settings));
    }
  },
  "/settings/sessions": {
    render: function() {
      return m(layout, m(sessions));
    }
  }
});

if (config.SENTRY_PUBLIC_DSN !== "") {
  console.log("installing sentry", config.SENTRY_PUBLIC_DSN);
  raven
    .config(config.SENTRY_PUBLIC_DSN, {
      environment: "HYDROCARBON_ENV",
      autoBreadcrumbs: {
        xhr: true, // XMLHttpRequest
        console: true, // console logging
        dom: true, // DOM interactions, i.e. clicks/typing
        location: true // url changes, including pushState/popState
      }
    })
    .install();
}
