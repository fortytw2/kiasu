import babel from "rollup-plugin-babel";
import babili from "rollup-plugin-babili";
import commonjs from "rollup-plugin-commonjs";
import replace from "rollup-plugin-replace";
import resolve from "rollup-plugin-node-resolve";

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
