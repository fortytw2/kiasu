import { Component, h } from "preact";

import Redux from "preact-redux";

class NotificationWindow extends Component {
  render() {
    return (
      <div>
        {this.props.notifications.map(n => (
          <div class="flex items-center justify-center pa2 bg-lightest-blue navy">
            <span class="lh-title ml1">{n.message}</span>
          </div>
        ))}
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
