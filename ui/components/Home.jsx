import { h, Component, render } from "preact";

class Home extends Component {
  render({}, { text = "Some Text" }) {
    return (
      <section class="home">
        <input value={text} />
        <div>In caps: {text.toUpperCase()}</div>
      </section>
    );
  }
}

export default Home;
