import m from "mithril";
import user from "../models/user";

export default {
  oncreate(vnode) {
    user.logout();
  },
  view(vnode) {
    return m.route.set("/login");
  }
};
