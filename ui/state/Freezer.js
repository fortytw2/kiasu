import Freezer from "freezer-js";

var fridge = new Freezer({
  loggedIn: {}
});

fridge.on("changeHomepage", function(val) {
  fridge.get().set({ loggedIn: val });
});

fridge.on("update", function(currentState, prevState) {
  console.log("event", currentState, prevState);
});

export default fridge;
