import { h, Component, render } from "preact";
import { Link } from "preact-router/match";
import Fridge from "../state/Freezer";

class Nav extends Component {
  render(props, state) {
    return (
      <nav class="pa3 pa4-ns">
        <Link
          class="link dim black b f6 f5-ns dib mr3"
          activeClassName="red"
          href="/"
        >
          hydrocarbon - {state.loggedIn}
        </Link>
        <Link
          class="link dim gray f6 f5-ns dib mr3"
          activeClassName="red"
          href="/login"
        >
          login
        </Link>
      </nav>
    );
  }
}

export default Nav;
