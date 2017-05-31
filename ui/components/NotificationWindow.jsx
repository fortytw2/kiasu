import { Component, h } from "preact";

import Notification from "./Notification";
import Redux from "preact-redux";

class NotificationWindow extends Component {
  render(props, state) {
    return (
      <div>
        {props.notifications.map(function(n) {
          return (
            <Notification
              dispatch={props.dispatch}
              sKey={n.key}
              message={n.message}
              level={n.level}
            />
          );
        })}
      </div>
    );
  }
}

const mapStateToProps = state => {
  return {
    notifications: state.notifications
  };
};

export default Redux.connect(mapStateToProps)(NotificationWindow);
