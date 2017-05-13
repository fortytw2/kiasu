// rollup.config.js
import resolve from "rollup-plugin-node-resolve";
import commonjs from "rollup-plugin-commonjs";
import babel from "rollup-plugin-babel";
export default {
  entry: "./main.jsx",
  format: "iife",
  sourceMap: true,
  plugins: [
    babel({
      exclude: "node_modules/**"
      // plugins appears to be ignored. use .babelrc
    }),
    resolve({ jsnext: true, main: true }),
    commonjs({
      extensions: [".jsx", ".js"]
    })
  ]
};
