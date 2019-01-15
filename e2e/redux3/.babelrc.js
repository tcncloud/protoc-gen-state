const presets = [require.resolve('@babel/preset-env'), require.resolve('@babel/preset-react')];
const plugins = [
  require.resolve('@babel/plugin-syntax-dynamic-import'),
  require.resolve('@babel/plugin-proposal-class-properties'),
  require.resolve('@babel/plugin-proposal-object-rest-spread'),
  [require.resolve('@babel/plugin-transform-runtime'), {
    'regenerator': true
  }]
];

module.exports = { presets, plugins };
