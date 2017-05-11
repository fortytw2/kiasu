import { h, Component, render } from "preact";
import Fridge from "../state/Freezer";

class Home extends Component {
  setText = e => {
    Fridge.trigger("changeHomepage", e.target.value);
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
