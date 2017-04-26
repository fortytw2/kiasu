import m from "mithril";
import flash from "../models/flash";

function binds(data) {
  return {
    onchange: function(e) {
      data[e.target.name] = e.target.value;
    }
  };
}

export default {
  view(vnode) {
    if (flash.flashMessage !== "") {
      var msg = flash.flashMessage
      flash.flashMessage = ""

      return m("div", binds(flash.flashMessage), msg);
    } else {
      return;
    }
  }
};
