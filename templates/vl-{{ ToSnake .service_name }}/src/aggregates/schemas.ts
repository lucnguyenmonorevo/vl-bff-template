{{ $_ := . -}}
{{ $_.import_schemas}}

const typeDefs = [{{ $_.type_defs }}]

export default typeDefs
