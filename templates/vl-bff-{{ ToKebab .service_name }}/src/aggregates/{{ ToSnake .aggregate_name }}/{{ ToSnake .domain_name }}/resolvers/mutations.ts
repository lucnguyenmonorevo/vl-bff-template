{{ $_ := . -}}
/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable @typescript-eslint/naming-convention */
import {
    MutationResolvers,
{{ $_.import_requests}}} from '../graphql_generated/graphql'

import {
    BaseRequest,
} from '../../../../utils/client'

{{ $_.import_convert }}
const mutations: MutationResolvers = {
  Mutation: {
    {{- $_.resolvers -}}
  },
}

export default mutations
