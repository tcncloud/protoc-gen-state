/// <reference path="neo/node_modules/@types/jasmine/index.d.ts" />
var Jasmine = require('jasmine');
var path = require('path');

var jasmine = new Jasmine();

var sandbox = path.dirname(path.resolve(__dirname));
process.env.BABEL_CACHE_PATH = sandbox + '/.babel.json';

jasmine.loadConfig({
  spec_dir: 'bazel/protoc-gen-state/example/fullApp',
  helpers: [
    // Explictly define some here so they get loaded in that order.
    'spec/helpers/babel.js',
    'spec/helpers/enzyme.js',
    'spec/helpers/jsdom.js',
    'spec/helpers/**/*.js',
  ],
  spec_files: [
    'src/**/*.test.[jt]s'
  ],
});

jasmine.execute();
