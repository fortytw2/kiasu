import { h, Component, render } from "preact";
import Redux from "preact-redux";
import Router from "preact-router";

import Store from "./state/Store";

import Nav from "./components/Nav";
import Home from "./components/Home";
import Login from "./components/Login";
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
        <Nav state={this.state} />
        <Router>
          <Home path="/" />
          <TextContent path="/about" text="this is the about page" />
          <Login path="/login" />
        </Router>
      </div>
    </Redux.Provider>
  );
};

render(<Container />, document.body);
