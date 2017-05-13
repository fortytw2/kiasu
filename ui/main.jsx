import { Component, h, render } from "preact";
import { Route, Router } from "preact-router";
import { Store, history } from "./state/Store";

import Home from "./components/Home";
import Login from "./components/Login";
import Nav from "./components/Nav";
import Redux from "preact-redux";
import TextContent from "./components/TextContent";
import { initDevTools } from "./vendor/devtools";

initDevTools();

class Container extends Component {
  render() {
    return <App />;
  }
}

const App = function() {
  return (
    <Redux.Provider store={Store}>
      <div class="min-vh-100">
        <Nav />
        <Router history={history}>
          <Route path="/" component={Home} />
          <Route path="/about" component={TextContent} />
          <Route path="/login" component={Login} />
        </Router>
      </div>
    </Redux.Provider>
  );
};

render(<Container />, document.body);
