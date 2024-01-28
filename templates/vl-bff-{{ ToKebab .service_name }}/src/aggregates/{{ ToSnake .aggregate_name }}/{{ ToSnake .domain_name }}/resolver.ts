import { Resolvers } from './graphql_generated/graphql'
import mutations from './resolvers/mutations'
import queries from './resolvers/queries'

// Note this "Resolvers" type isn't strictly necessary because we are already
// separately type checking our queries and resolvers. However, the "Resolvers"
// graphql_generated types is useful syntax if you are defining your resolvers
// in a single file.
const userResolvers: Resolvers = { ...queries, ...mutations }

export default userResolvers
