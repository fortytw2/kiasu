import Router from "preact-router";
import { h, render } from "preact";

import Nav from "./components/Nav";
import Home from "./components/Home";
import Login from "./components/Login";
import TextContent from "./components/TextContent";

const Main = () => (
  <div class="min-vh-100">
    <Nav />
    <Router>
      <Home path="/" />
      <TextContent path="/about" text="this is the about page" />
      <Login path="/login" />
    </Router>
  </div>
);

render(<Main />, document.body);
