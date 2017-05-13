import Router from "preact-router";
import { h, Component, render } from "preact";

import Nav from "./components/Nav";
import Home from "./components/Home";
import Login from "./components/Login";
import TextContent from "./components/TextContent";
import Fridge from "./state/Freezer";

import { initDevTools } from "./vendor/devtools";

initDevTools();

class Container extends Component {
  render() {
    // 1. Your app receives the state
    var state = Fridge.get();
    return <App state={state} />;
  }
  componentDidMount() {
    var me = this;
    // 2. Your app get re-rendered on any state change
    Fridge.on("update", function() {
      me.setState({});
    });
  }
}

const App = function() {
  return (
    <div class="min-vh-100">
      <Nav state={this.state} />
      <Router>
        <Home path="/" />
        <TextContent path="/about" text="this is the about page" />
        <Login path="/login" />
      </Router>
    </div>
  );
};

render(<Container />, document.body);
