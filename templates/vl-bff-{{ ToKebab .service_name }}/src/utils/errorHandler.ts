import { GraphQLError } from 'graphql/error'
import { Timestamp } from 'google-protobuf/google/protobuf/timestamp_pb'
import { Error as ResBaseError } from '../../proto_generated/common/grpc/grpc_pb'

export function createBaseGraphQLError(resBaseError: ResBaseError.AsObject): GraphQLError {
  return new GraphQLError(resBaseError.messageOnFrontend, {
    extensions: {
      code: resBaseError.code,
      exception: resBaseError.exception,
      http: {
        status: resBaseError.httpCode,
      },
    },
  })
}
export function timestampToString(timestamp: Timestamp): string {
  if (!timestamp) {
    return ''
  }
  // eslint-disable-next-line @typescript-eslint/no-unsafe-member-access
  const date = timestamp.toDate() // Convert to JavaScript Date object
  // eslint-disable-next-line @typescript-eslint/no-unsafe-return, @typescript-eslint/no-unsafe-member-access
  return date.toISOString() // Convert to ISO string (or any other string format you want)
}
