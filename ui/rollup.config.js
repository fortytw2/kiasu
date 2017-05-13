// rollup.config.js
import resolve from "rollup-plugin-node-resolve";
import commonjs from "rollup-plugin-commonjs";
import babel from "rollup-plugin-babel";
import babili from "rollup-plugin-babili";
import replace from "rollup-plugin-replace";

export default {
  entry: "./main.jsx",
  format: "iife",
  sourceMap: true,
  plugins: [
    babel({
      exclude: "node_modules/**"
      // plugins appears to be ignored. use .babelrc
    }),
    replace({
      "process.env.NODE_ENV": JSON.stringify("production")
    }),
    resolve({ jsnext: true, main: true }),
    commonjs({
      extensions: [".jsx", ".js"],
      include: ["node_modules/**/*"],
      namedExports: {
        "preact-redux": ["connect", "Provider", "connectAdvanced"]
      }
    }),
    babili({
      comments: false
    })
  ]
};
