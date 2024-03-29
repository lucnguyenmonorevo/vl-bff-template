{{ $_ := . -}}
import * as grpc from '@grpc/grpc-js'

import winston from 'winston'

import {
    BaseRequest,
    convertBaseRequest,
    convertBaseResponse,
    createBaseGraphQLError,
} from '../../../../utils/client'

import { {{ ToCamel $_.domain_name }}ServiceClient as Client } from '../../../../../proto_generated/{{ ToSnake $_.service_name }}/aggregates/{{ ToSnake $_.aggregate_name }}/{{ ToSnake $_.domain_name }}_service_grpc_pb'

import {
    {{ GetUpperCaseChars $_.domain_name }}{{ ToCamel $_.request_name }} as GraphQLRequest,
    // {{ GetUpperCaseChars $_.domain_name }}{{ ToCamel $_.request_name }}Data as GraphQLRequestData,
    {{ GetUpperCaseChars $_.domain_name }}{{ ToCamel $_.response_name }} as GraphQLResponse,
    // {{ GetUpperCaseChars $_.domain_name }}{{ ToCamel $_.response_name }}Data as GraphQLResponseData,
} from '../graphql_generated/graphql'

import {
    {{ convertImportGrpc $_.request_name }} as GRPCRequest,
    // {{ convertImportGrpc $_.request_name }}Data as GRPCRequestData,
    {{ convertImportGrpc $_.response_name }} as GRPCResponse,
    // {{ convertImportGrpc $_.response_name }}Data as GRPCResponseData,
} from '../../../../../proto_generated/{{ $_.request_path }}_pb'

import { {{ ToUpperSnake $_.service_name}}_GRPC_HOST as host, TIME_OUT_SECOND } from '../../../../../config'

const logger = winston.createLogger({
    transports: [new winston.transports.Console()],
})

// Convert request
function convertRequestGraphQLToGRPC(
    baseRequest: BaseRequest,
    graphqlReq: GraphQLRequest,
) {
    const rt = new GRPCRequest()
    const base = convertBaseRequest(baseRequest)
    // TODO: convert request
    rt.setBase(base)
    return rt
}

// Convert response
function convertResponseGRPCToGraphQL( grpcRes: GRPCResponse.AsObject ) {
    const grpcResData = grpcRes.data
    const rt: GraphQLResponse = {
        // TODO: convert response
    }
    return rt
}

export async function {{ ToLowerCamel .domain_name }}{{ ToCamel .name }}(
    baseRequest: BaseRequest,
    dataRequest: GraphQLRequest,
) {
    return new Promise((resolve, reject) => {
        const client = new Client(host, grpc.credentials.createInsecure())

        const deadline = new Date()
        deadline.setSeconds(deadline.getSeconds() + TIME_OUT_SECOND)

        client.waitForReady(deadline, (connectionError?: Error) => {
            if (connectionError) {
                logger.error(connectionError)
            } else {
                const req = convertRequestGraphQLToGRPC(baseRequest, dataRequest)
                try {
                    client.{{ ToLowerCamel $_.name }}(
                        req,
                        {},
                        (serverError: grpc.ServiceError | null, serverMessage?: GRPCResponse) => {
                            if (serverError) {
                                logger.error(serverError)
                                return reject(serverError)
                            }
                            if (serverMessage) {
                                const response = serverMessage.toObject()
                                const resBase = response.base
                                if (resBase) {
                                    const resBaseError = resBase.error
                                    if (resBaseError) {
                                        return reject(createBaseGraphQLError(resBaseError))
                                    }
                                }
                                const res = convertResponseGRPCToGraphQL(response)
                                return resolve(res)
                            }
                            return resolve
                        },
                    )
                } catch (error) {
                    logger.info(error)
                }
            }
        })
    })
}
