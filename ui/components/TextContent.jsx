import { h, Component, render } from "preact";

class TextContent extends Component {
  render({ text = "" }, {}) {
    return (
      <section>
        <p>{text}</p>
      </section>
    );
  }
}

export default TextContent;
