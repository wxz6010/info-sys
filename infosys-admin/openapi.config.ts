const { generateService } = require('@umijs/openapi')
generateService({
  schemaPath: 'http://localhost:8081/api/openapi',
  serversPath: './src/servers/gen-cache',

})