const moduleAlias = require('module-alias');

moduleAlias.addAliases({
  'proto': '../protos/',
  '@App': 'neo/'
});

moduleAlias('package.json');
