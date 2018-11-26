const moduleAlias = require('module-alias');

moduleAlias.addAliases({
  'proto': '../proto/',
  '@App': 'neo/'
});

moduleAlias('package.json');
