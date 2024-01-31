import { GraphQLError } from 'graphql/error'
import { Timestamp } from 'google-protobuf/google/protobuf/timestamp_pb'
import {
  GRPCRequest,
  GRPCResponse,
  Error as GRPCError,
} from '../../proto_generated/common/grpc/grpc_pb'

export function createBaseGraphQLError(resBaseError: GRPCError.AsObject): GraphQLError {
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

export type BaseRequest = {
  accessToken: string
  refreshToken: string
  pageId: number
  functionId: number
  baseId: number
  language: string
}

export type BaseResponse = {
  accessToken: string
  refreshToken: string
  err: Error
}

export type Error = {
  code: string
  exception: string
  messageOnLog: string
  messageOnFrontend: string
  httpCode: number
}

export function convertBaseRequest(req: BaseRequest) {
  const base = new GRPCRequest()
  base.setAccessToken(req.accessToken)
  base.setRefreshToken(req.refreshToken)
  base.setPageId(req.pageId)
  base.setFunctionId(req.functionId)
  base.setBaseId(req.baseId)
  base.setLanguage(req.language)
  return base
}

export function convertBaseResponse(req: GRPCResponse.AsObject) {
  const grpcError = req.error
  const error = new GRPCError()
  error.setCode(grpcError.code)
  error.setException(grpcError.exception)
  error.setMessageOnLog(grpcError.messageOnLog)
  error.setMessageOnFrontend(grpcError.messageOnFrontend)
  error.setHttpCode(grpcError.httpCode)

  const base = new GRPCResponse()
  base.setAccessToken(req.accessToken)
  base.setRefreshToken(req.refreshToken)
  base.setError(error)
  return base
}
