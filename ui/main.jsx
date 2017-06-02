import { Component, h, render } from "preact";
import { History, Store } from "./state/Store";
import { Route, Router } from "preact-router";

import Footer from "./components/Footer";
import Login from "./Components/Login";
import LoginCallback from "./Components/LoginCallback";
import Nav from "./Components/Nav";
import NotFound from "./Components/NotFound";
import NotificationWindow from "./Components/NotificationWindow";
import RehydrateProvider from "./Components/RehydrateProvider";
import TextContent from "./Components/TextContent";
import { initDevTools } from "./vendor/devtools";

initDevTools();

const App = function() {
  return (
    <RehydrateProvider store={Store}>
      <div class="min-vh-100">
        <Nav />
        <NotificationWindow />
        <Router history={History}>
          <Route
            path="/"
            component={TextContent}
            text="this is the home page"
          />
          <Route
            path="/about"
            component={TextContent}
            text="hi this is about page"
          />
          <Route path="/login" component={Login} />
          <Route path="/login-callback" component={LoginCallback} />
          <Route component={NotFound} default />
        </Router>
        <Footer />
      </div>
    </RehydrateProvider>
  );
};

render(<App />, document.body);
