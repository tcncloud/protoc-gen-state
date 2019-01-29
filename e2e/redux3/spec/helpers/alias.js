const moduleAlias = require('module-alias');

moduleAlias.addAliases({
  'protos': '../../protos/',
  '@App': '../../'
});

moduleAlias('package.json');
