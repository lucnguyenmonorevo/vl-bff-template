/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable @typescript-eslint/naming-convention */

import { QueryResolvers } from '../graphql_generated/graphql'
import { getUserLoggedInfo } from '../grpc_clients/user_logged_info'

const queries: QueryResolvers = {
  Query: {
    getUserLoggedInfo: async (
      _: any,
      params: { access_token: string },
      contextValue: { token: string },
    ) => {
      const resp = await getUserLoggedInfo(params, contextValue)
      return resp
    },
  },
}
export default queries
