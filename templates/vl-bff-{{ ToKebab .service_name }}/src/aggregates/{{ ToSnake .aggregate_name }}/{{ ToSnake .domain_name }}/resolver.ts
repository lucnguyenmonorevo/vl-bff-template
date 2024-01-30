{{ $_ := . -}}
import { Resolvers } from './graphql_generated/graphql'
{{ $_.resolver_ts_import -}}

// Note this "Resolvers" type isn't strictly necessary because we are already
// separately type checking our queries and resolvers. However, the "Resolvers"
// graphql_generated types is useful syntax if you are defining your resolvers
// in a single file.
const {{ ToLowerCamel $_.domain_name }}Resolvers: Resolvers = { {{ $_.resolver_ts_resolvers }} }

export default {{ ToLowerCamel $_.domain_name }}Resolvers
