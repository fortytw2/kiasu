// rollup.config.js
import resolve from "rollup-plugin-node-resolve";
import commonjs from "rollup-plugin-commonjs";
import babel from "rollup-plugin-babel";

export default {
  entry: "./main.js",
  format: "iife",
  sourceMap: true,
  plugins: [resolve({ jsnext: true, main: true }), commonjs(), babel()]
};
