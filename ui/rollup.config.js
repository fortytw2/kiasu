// rollup.config.js
import resolve from "rollup-plugin-node-resolve";
import commonjs from "rollup-plugin-commonjs";

export default {
  entry: "./main.js",
  format: "iife",
  sourceMap: true,
  plugins: [resolve({ jsnext: true, main: true }), commonjs()]
};
