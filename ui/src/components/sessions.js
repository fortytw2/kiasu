import m from "mithril";
import user from "../models/user";

export default {
  sessions: [],
  oninit(vnode) {
    user.getSessions();
    console.log(user.sessions);
  },
  view(vnode) {
    return m(
      "div",
      user.sessions.map(s => {
        m("p", s);
      })
    );
  }
};
