import { Component, h } from "preact";

import { NOTIFICATION_LEVEL_WARNING } from "../state/Notifications";
import Redux from "preact-redux";

class NotificationWindow extends Component {
  render(props, state) {
    return (
      <div>
        {props.notifications.map(function(n) {
          var style = n.level === NOTIFICATION_LEVEL_WARNING
            ? "flex items-center justify-center pa2 " + "bg-washed-red"
            : "flex items-center justify-center pa2 " + "bg-lightest-blue";
          return (
            <div class={style}>
              <span class="lh-title ml1">{n.message}</span>
            </div>
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
