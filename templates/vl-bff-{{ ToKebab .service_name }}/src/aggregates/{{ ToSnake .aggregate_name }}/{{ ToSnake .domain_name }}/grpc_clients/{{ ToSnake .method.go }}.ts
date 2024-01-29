{{ $_ := . -}}

import * as grpc from '@grpc/grpc-js'

import winston from 'winston'

import { CustomerOrganizationServiceClient as Client } from '../proto_generated/business/aggregates/sale/customer_organization_service_grpc_pb'

import { GRPCRequest } from '../proto_generated/common/grpc/grpc_pb'
import { CustomerManagementCompanyInfoCreateInput as GraphQLInput } from '../graphql_generated/graphql'
import {
    CustomerOrganizationCreateRequest as GRPCRequestObj,
    CustomerOrganizationCreateResponse as GRPCResponseObj,
} from '../proto_generated/business/aggregates/sale/customer_organization_create_pb'

import { createBaseGraphQLError } from './common'
import { convertGraphQLToGRPC, convertGRPCToGraphQL } from './companyInfoCreateConvert'
import { BUSINESS_GRPC_HOST as host, TIME_OUT_SECOND } from '../../../../../config'

const logger = winston.createLogger({
    transports: [new winston.transports.Console()],
})

export async function companyInfoCreate(contextValue: { token: string }, params: GraphQLInput) {
    return new Promise((resolve, reject) => {
        const client = new Client(host, grpc.credentials.createInsecure())

        const deadline = new Date()
        deadline.setSeconds(deadline.getSeconds() + TIME_OUT_SECOND)

        client.waitForReady(deadline, (connectionError?: Error) => {
            if (connectionError) {
                logger.error(connectionError)
            } else {
                const req = new GRPCRequestObj()
                const base = new GRPCRequest()
                const reqData = convertGraphQLToGRPC(params)
                req.setData(reqData)

                base.setToken(contextValue.token)
                req.setBase(base)

                try {
                    client.create(
                        req,
                        {},
                        (serverError: grpc.ServiceError | null, serverMessage?: GRPCResponseObj) => {
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
                                const res = convertGRPCToGraphQL(serverMessage.getData())
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
