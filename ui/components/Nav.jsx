import { Component, h } from "preact";

import { Link } from "preact-router/match";
import Redux from "preact-redux";

class Nav extends Component {
  loggedIn() {
    if (this.props.login.email !== "") {
      return true;
    }
    return false;
  }
  render(props, state) {
    return (
      <nav class="pa1 pa2-ns">
        <Link
          class="link dim black b f6 f5-ns dib mr3"
          activeClassName="blue"
          href="/"
        >
          hydrocarbon
        </Link>
        {this.loggedIn()
          ? <Link
              class="link dim gray f6 f5-ns dib mr3"
              activeClassName="blue"
              href="/feed"
            >
              {props.login.email}
            </Link>
          : <Link
              class="link dim gray f6 f5-ns dib mr3"
              activeClassName="blue"
              href="/login"
            >
              login
            </Link>}

      </nav>
    );
  }
}

const mapStateToProps = state => {
  return {
    login: state.login
  };
};

export default Redux.connect(mapStateToProps)(Nav);
