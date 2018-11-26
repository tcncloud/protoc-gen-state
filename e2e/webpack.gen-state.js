const webpack = require('webpack');
const HtmlWebPackPlugin = require('html-webpack-plugin');
const path = require('path');

module.exports = (env) => {
  const config = {
    mode: env && env.development? 'development' : 'production',

    entry: {
      main: [
        require.resolve('./webpack.polyfills.js'),
        path.resolve(__dirname, 'src/index.tsx'),
      ],
    },

    output: {
      path: path.resolve(__dirname, 'gen-state-dist'),
      filename: 'bundle.js',
      publicPath: '/'
    },

    devtool: 'inline-source-map',

    resolve: {
      modules: [path.resolve(__dirname, 'node_modules')],
      alias: {
        '@App': path.resolve(__dirname, './'),
        'proto': path.resolve(__dirname, 'generated/'),
      },
      extensions: ['.js', '.jsx', '.ts', '.tsx'],
      symlinks: false,
    },

    // DEVSERVER SETTINGS (webpack-dev-server)
    devServer: {
      contentBase: path.join(__dirname, 'gen-state-dist'),
      historyApiFallback: true,
      overlay: true,
      port: 3000,
      stats: 'normal',
      open: true,
    },

    module: {
      rules: [
        {
          test: /\.css$/,
          use: [ 'style-loader', 'css-loader' ]
        },
        {
          test: /\.svg$/,
          loader: 'raw-loader'
        },
        {
          test: /\.tsx?$/, // ts || tsx
          exclude: /node_modules/,
          use: [
            { loader: 'babel-loader' },
            { loader: 'ts-loader',
              options: {
                configFile: "tsconfig.gen-state.json",
                // Uncomment these if you want to skip transpiling for some demonic reason
                // transpileOnly: true,
                // happyPackMode: true
              }
            },
            { loader: 'thread-loader' }
          ]
        },
      ]
    },

    resolveLoader: {
      modules: [path.resolve(__dirname, 'node_modules')],
    },

    plugins: [
      // Config for HTML Webpack Plugin (goes with HTML Loader)
      new HtmlWebPackPlugin({
        template: path.resolve(__dirname, 'src/public/index.html'),
        filename: 'index.html',
      }),
    ]
  };

  return config;
};
