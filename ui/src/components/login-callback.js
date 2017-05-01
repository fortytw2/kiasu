import m from "mithril";

export default {
  view(vnode) {
    return m(
      "div",
      {
        class: "measure center"
      },
      m.route.param("key")
    );
  }
};
