{{ $_ := . -}}
{{ $_.import_resolvers }}

const resolvers = {
{{ $_.resolvers }}
}

export default resolvers
