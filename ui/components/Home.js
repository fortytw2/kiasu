import { h, Component, render } from "preact";

class Home extends Component {
  setText = e => {
    this.setState({ text: e.target.value });
  };
  render({}, { text = "Some Text" }) {
    return (
      <section class="home">
        <input value={text} onInput={this.setText} />
        <div>In caps: {text.toUpperCase()}</div>
      </section>
    );
  }
}

export default Home;
