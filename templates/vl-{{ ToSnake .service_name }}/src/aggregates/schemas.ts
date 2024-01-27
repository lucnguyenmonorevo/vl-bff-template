{{ $_ := . -}}
import gql from 'graphql-tag'
{{ $_.import_schemas}}

{{ $_.schemas_to_type_defs }}

const typeDefs = [{{ $_.type_defs }}]

export default typeDefs
