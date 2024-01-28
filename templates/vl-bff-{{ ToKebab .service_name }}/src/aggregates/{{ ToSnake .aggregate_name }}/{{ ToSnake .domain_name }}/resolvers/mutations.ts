/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable @typescript-eslint/naming-convention */

import { MutationResolvers } from '../graphql_generated/graphql'
import { login } from '../grpc_clients/login'
import { logout } from '../grpc_clients/logout'

const mutations: MutationResolvers = {
  Mutation: {
    login: async (
      _: any,
      params: { username: string; password: string; organization_name: string },
    ) => {
      const resp = await login(params)
      return resp
    },
    logout: async (_: any, params: { access_token: string }, contextValue: { token: string }) => {
      const resp = await logout(params, contextValue)
      return resp
    },
  },
}

export default mutations
