const HtmlWebPackPlugin = require('html-webpack-plugin');
const path = require('path');
const htmlWebpackPlugin = new HtmlWebPackPlugin({
  template: './src/index.html',
  filename: './index.html'
});
module.exports = {
  entry: [
    '@babel/polyfill',
    './src/index.js',
  ],
  output: {
    filename: 'app.js',
    // path: __dirname + '/dist'
    path: path.resolve(__dirname, '../public/build'),
  },
  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: /node_modules/,
        use: {
          loader: "babel-loader"
        }
      },
      {
        test: /\.css$/,
        use: [
          {
            loader: "style-loader"
          },
          {
            loader: "css-loader",
            options: {
              modules: true,
              importLoaders: 1,
              localIdentName: "[name]_[local]_[hash:base64]",
              sourceMap: true,
              minimize: true
            }
           }
         ]
       }
     ]
  },
  devServer: {
    historyApiFallback: false,
  },
  plugins: [htmlWebpackPlugin],
  watchOptions: {
    aggregateTimeout: 300,
    poll: 1000
  }
};