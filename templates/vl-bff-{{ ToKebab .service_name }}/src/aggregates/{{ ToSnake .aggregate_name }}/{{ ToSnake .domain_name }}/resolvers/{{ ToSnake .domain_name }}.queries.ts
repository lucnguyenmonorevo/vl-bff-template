{{ $_ := . -}}
/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable @typescript-eslint/naming-convention */
import {
  QueryResolvers,
{{ $_.import_requests}}} from '../graphql_generated/graphql'

import {
  BaseRequest,
} from '../../../../utils/client'

{{ $_.import_convert }}
const toSnakeDomain_nameQueries: QueryResolvers = {
  Query: {
    {{- $_.resolvers -}}
  },
}
export default toSnakeDomain_nameQueries
