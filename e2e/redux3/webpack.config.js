
// const fs = require('fs');
const webpack = require('webpack');
const path = require('path');
const HtmlWebPackPlugin = require('html-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const BundleAnalyzerPlugin = require('webpack-bundle-analyzer').BundleAnalyzerPlugin;
const ForkTsCheckerWebpackPlugin = require('fork-ts-checker-webpack-plugin');
const SpeedMeasurePlugin = require("speed-measure-webpack-plugin");

// Behold. The Webpack Config.
module.exports = (env, argv) => {
  const buildMode = argv && argv.mode === 'development' ? 'development' : 'production'
  console.log('Webpack Mode: ', buildMode);

  const config = {
    mode: buildMode,
    entry: {
      main: [
        require.resolve('./webpack.polyfills.js'),
        path.resolve(__dirname, 'src/index.tsx'),
      ],
    },

    output: {
      filename: 'static/js/[contenthash:8].js',
      path: path.resolve(__dirname, 'dist'),
      publicPath: '/',
    },

    stats: {
      // copied from `'minimal'`
      all: false,
      modules: false,
      maxModules: 0,
      errors: true,
      warnings: true,
      // our additional option
      assets: true,
      moduleTrace: false,
      errorDetails: false,
      builtAt: true,
      cached: true,
      colors: true,
      env: true,
      timings: true,
      depth: false,
      entrypoints: false,
      publicPath: false,
      reasons: false,
    },

    // WEBPACK PLUGINS CONFIG
    plugins: [
      // Config for HTML Webpack Plugin (goes with HTML Loader)
      new HtmlWebPackPlugin({
        template: path.resolve(__dirname, 'src/public/index.html'),
        filename: 'index.html',
      }),

      // Works with CSS Loader to extract CSS to its own file
      new MiniCssExtractPlugin({
        filename: 'static/css/[contenthash:8].css',
        chunkFilename: 'static/css/[contenthash:8].css',
      }),

      // Runs Typescript lint checking and type checking in seperate process
      new ForkTsCheckerWebpackPlugin(),

      // Inject build time variables.
      new webpack.DefinePlugin({
        'process.env.VERSION': JSON.stringify(env && env.version ? env.version : 'development'),
        'process.env.NODE_ENV': JSON.stringify(buildMode === 'production' ? 'production' : 'development')
      })

    ],

    // WEBPACK RESOLVER CONFIG
    resolve: {
      modules: [
        path.resolve(__dirname, 'node_modules'),
      ],
      alias: {
        'protos': path.resolve(__dirname, 'protos/'),// the bazel-genfiles path is modified by our webpack rule
        '@App': path.resolve(__dirname, './'),
      },
      extensions: ['.ts', '.tsx', '.js', '.jsx'],
      symlinks: false,
    },

    resolveLoader: {
      modules: [
        path.resolve(__dirname, 'node_modules'),
      ],
      symlinks: false,
    },

    // DEVSERVER SETTINGS (webpack-dev-server)
    devServer: {
      contentBase: path.join(__dirname, 'dist'),
      historyApiFallback: true,
      overlay: true,
      port: 3000,
      clientLogLevel: 'none',
      stats: {
        // copied from `'minimal'`
        all: false,
        modules: false,
        maxModules: 0,
        errors: true,
        warnings: true,
        // our additional options
        moduleTrace: false,
        errorDetails: false,
        builtAt: true,
        cached: true,
        colors: true,
        env: true,
        timings: true,
        depth: false,
        entrypoints: false,
        publicPath: false,
        reasons: false,
      },
      open: true,
    },

    // Vendor chunking
    optimization: {
      runtimeChunk: false,
      splitChunks: {
        cacheGroups: {
          default: false,
          node_vendors: {
            test: /[\\/]node_modules[\\/]/,
            name: 'node_modules',
            chunks: 'all',
            minChunks: 2,
          },
        },
      },
    },

    // LOADERS AND RULES CONFIG
    module: {
      rules: [
        // Babel Loader
        {
          test: /\.(j|t)sx?$/, // js || jsx || ts || tsx
          include: [
            path.resolve(__dirname, 'protos/'),
            path.resolve(__dirname, 'src/'),
          ],
          exclude: [
            /node_modules/,
            /webphone/,
          ],
          use: [
            {
              loader: 'babel-loader',
              query: {
                cacheDirectory: true, // adds caching
                babelrc: false, // dont look for a babelrc file
                // sourceType Fixes issue with exports in ts-protoc-gen becacuse they are common js (rather than es modules)
                // when using useBuiltIns: 'usage'. Not needed when using useBuiltIns: false.
                // https://babeljs.io/docs/en/options#sourcetype
                // sourceType: 'unambiguous',
                presets: [ // Presets run right to left (bottom to top)
                  [
                    require.resolve('@babel/preset-env'), // https://babeljs.io/docs/en/babel-preset-env
                    {
                      // Which browsers / envrionments to target. Uses browserlist compat query (https://github.com/browserslist/browserslist)
                      targets: '> 0.5%, not ie <= 10, not dead',

                      // Includes core-js polyfills
                      // https://babeljs.io/docs/en/babel-preset-env#usebuiltins
                      // Requires peer dependency of @babel/polyfill (in deps, not dev deps).
                      // see link above for info about 'usage'. Has some side effects due to ts-protoc-gen.
                      // Don't recommend using right now (12/5/2018)
                      // Using the false value requires adding the polyfills to the inputs array.
                      // They are included in webpack.polyfills.js
                      useBuiltIns: 'entry', // false, //'usage',
                    }
                  ],
                  require.resolve('@babel/preset-react'),
                  require.resolve('@babel/preset-typescript'),
                ],
                plugins: [
                  // Plugins not provided by preset env
                  require.resolve('@babel/plugin-proposal-class-properties'),
                  require.resolve('@babel/plugin-proposal-object-rest-spread'),
                  require.resolve('@babel/plugin-syntax-dynamic-import'),
                  [require.resolve('@babel/plugin-transform-runtime'), {
                    'regenerator': true
                  }]
                ]
              }
            },
          ],
        },

        // Normal images within project
        {
          test: /\.(png|jpg|gif|svg|ico)$/,
          use: [
            {
              loader: 'file-loader',
              options: {
                name: 'static/images/[name].[ext]',
              },
            },
          ],
        },

        // Fonts that are part of the project.
        {
          test: /\.(ttf|eot|woff|woff2)$/,
          use: [
            {
              loader: 'file-loader',
              options: {
                name: 'static/fonts/[name].[ext]',
              },
            },
          ],
        },

        // SASS and CSS Loading
        // TODO: Configure browserlistrc
        {
          test: /\.scss$/,
          use: [
            MiniCssExtractPlugin.loader,
            'css-loader',
            {
              loader: 'postcss-loader',
              options: {
                plugins: () => [
                  require('autoprefixer'),
                ],
              },
            },
            {
              loader: 'sass-loader',
              options: {
                includePaths: [
                  __dirname,
                ],
              },
            },
          ],
        },

        {
          test: /\.css$/,
          use: [
            MiniCssExtractPlugin.loader,
            'css-loader',
            {
              loader: 'postcss-loader',
              options: {
                plugins: () => [
                  require('autoprefixer'),
                ],
              },
            },
          ],
        },

        // HTML Loader
        {
          test: /\.html$/,
          exclude: [/node_modules/],
          use: [
            {
              loader: 'html-loader',
              options: { minimize: true },
            },
          ],
        },

      ],
    },
  };

  // Analyzes the bundles in a visual fashion if you add the flag --env.analyze
  if (env && env.analyze_size) {
    config.plugins.push(new BundleAnalyzerPlugin());
  }

  if(env && env.analyze_speed) {
    const smp = new SpeedMeasurePlugin();
    return smp.wrap(config);
  }

  return config;
};
