import { Component, h } from "preact";

import Redux from "preact-redux";

class NotificationWindow extends Component {
  render() {
    return (
      <div>
        <h1>notifications</h1>
        {this.props.notifications.map(n => {
          return <h1>{n.message}</h1>;
        })}
      </div>
    );
  }
}

const mapStateToProps = state => {
  console.log("mapSTateTOProps", state);
  return {
    notifications: state.notifications
  };
};

export default Redux.connect(mapStateToProps)(NotificationWindow);
