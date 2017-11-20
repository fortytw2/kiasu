import babel from "rollup-plugin-babel";
import babel_minify from "rollup-plugin-babel-minify";
import commonjs from "rollup-plugin-commonjs";
import replace from "rollup-plugin-replace";
import resolve from "rollup-plugin-node-resolve";

export default {
  input: "./main.jsx",
  output: {
    format: "iife"
  },
  sourcemap: true,
  plugins: [
    babel({
      exclude: "node_modules/**"
    }),
    // redux puts this crap in their npm package...
    replace({
      "process.env.NODE_ENV": JSON.stringify("production")
    }),
    resolve({ jsnext: true, main: true, extensions: [".js", ".jsx"] }),
    commonjs({
      extensions: [".jsx", ".js"]
    }),
    babel_minify({
      comments: false
    })
  ]
};
