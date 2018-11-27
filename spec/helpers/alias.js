const moduleAlias = require('module-alias');

moduleAlias.addAliases({
  'proto': '../generated/',
  '@App': 'neo/'
});

moduleAlias('package.json');
