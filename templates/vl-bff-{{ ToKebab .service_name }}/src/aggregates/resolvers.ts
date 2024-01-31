{{ $_ := . -}}
{{ $_.import_resolvers }}
/* eslint-disable @typescript-eslint/no-var-requires */
const { mergeResolvers } = require("@graphql-tools/merge");
const resolvers = mergeResolvers(
{{ $_.resolvers }}

)

export default resolvers
