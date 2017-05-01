import m from "mithril";
import user from "../models/user";

export default {
  view(vnode) {
    user.activateToken(m.route.param("key"));
    m.route.set("/");
  }
};
