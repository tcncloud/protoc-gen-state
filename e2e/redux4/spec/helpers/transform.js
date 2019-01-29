console.log('Jasmine Helpers: Requiring transforms...');

require.extensions['.svg'] = () => { 'module.exports = ""'; };
require.extensions['.css'] = () => { 'module.exports = ""'; };
require.extensions['.scss'] = () => { 'module.exports = ""'; };
