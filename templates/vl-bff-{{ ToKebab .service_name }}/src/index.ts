/* eslint-disable no-promise-executor-return */
/* eslint-disable @typescript-eslint/no-unsafe-assignment */
import express from 'express'
import cors from 'cors'
import { json } from 'body-parser'
import http from 'http'

import { buildSubgraphSchema } from '@apollo/subgraph'
import { ApolloServer, BaseContext } from '@apollo/server'
import { expressMiddleware } from '@apollo/server/express4'
import { ApolloServerPluginDrainHttpServer } from '@apollo/server/plugin/drainHttpServer'
import winston from 'winston'
import * as dotenv from 'dotenv'
import gql from 'graphql-tag'
// import resolvers from './aggregates/resolvers'
const resolvers  = {}
import schemas from './aggregates/schemas'
import apolloWinstonLoggingPlugin from './utils/logger'
import { PORT } from '../config'

dotenv.config()

const logger = winston.createLogger({
  format: winston.format.simple(),
  transports: [
    new winston.transports.Console(),
    new winston.transports.File({ filename: 'logs.log' }),
  ],
})

const typeDefs = gql(schemas)
export interface ApolloServerContext extends BaseContext {
  token?: string
}

async function startApolloServer() {
  const app = express()
  const httpServer = http.createServer(app)

  const server = new ApolloServer<ApolloServerContext>({
    schema: buildSubgraphSchema({ typeDefs, resolvers }),
    plugins: [
      ApolloServerPluginDrainHttpServer({ httpServer }),
      apolloWinstonLoggingPlugin({
        winstonInstance: logger,
        config: { enableWillSendResponse: true },
      }),
    ],
    includeStacktraceInErrorResponses: false,
  })

  await server.start()

  app.get('/healthcheck', (req, res) => {
    res.status(200).send('healthy')
  })

  app.use(
    '/graphql',
    cors<cors.CorsRequest>(),
    json(),
    expressMiddleware(server, {
      context: async ({ req }) => {
        const promise = Promise.resolve('value')
        await promise
        return {
          token: req.headers.token,
        }
      },
    }),
  )

  await new Promise<void>((resolve) =>
    app.listen(PORT, (): void => {
      logger.info(`🌏 Express server started at http://localhost:${PORT}`)
      logger.info(`🚀 Graphql ready at HTTP ${PORT} http://localhost:${PORT}/graphql`)
      resolve()
    }),
  )
}

startApolloServer().catch((err) => {
  logger.info('Failed to start Apollo Server:', err)
})
