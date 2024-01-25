/* eslint-disable @typescript-eslint/no-explicit-any */
import { v4 as uuidv4 } from 'uuid'
import winston from 'winston'

const stringify = (obj: unknown) => JSON.stringify(obj)
const logger = winston.createLogger({
  transports: [new winston.transports.Console()],
})

interface Options {
  config?: {
    enableDidEncounterErrors?: boolean
    enableDidResolveOperation?: boolean
    enableExecutionDidStart?: boolean
    enableParsingDidStart?: boolean
    enableResponseForOperation?: boolean
    enableValidationDidStart?: boolean
    enableWillSendResponse?: boolean
    enableRequestDidStart?: boolean
  }

  winstonInstance?: winston.Logger

  levels?: {
    debug?: string
    info?: string
    error?: string
  }
}

const apolloWinstonLoggingPlugin = (opts: Options = {}): any => {
  const {
    enableDidEncounterErrors = true,
    enableDidResolveOperation = false,
    enableExecutionDidStart = false,
    enableParsingDidStart = false,
    enableResponseForOperation = false,
    enableValidationDidStart = false,
    enableWillSendResponse = true,
    enableRequestDidStart = true,
  } = opts.config || {}

  const { debug = 'debug', info = 'info', error = 'error' } = opts.levels || {}
  const { winstonInstance = logger } = opts

  return {
    requestDidStart(context: {
      request: { query: string; operationName: string; variables: string }
    }) {
      const id: string = uuidv4()
      const { query, operationName, variables } = context.request
      if (enableRequestDidStart) {
        winstonInstance.log(
          info,
          stringify({
            id,
            event: 'request',
            operationName,
            query: query?.replace(/\s+/g, ' '),
            variables,
          }),
        )
      }
      const handlers = {
        didEncounterErrors({ errors }) {
          if (enableDidEncounterErrors) {
            winstonInstance.log(error, stringify({ id, event: 'errors', errors }))
          }
        },

        didResolveOperation(ctx) {
          if (enableDidResolveOperation) {
            winstonInstance.log(
              debug,
              stringify({
                id,
                event: 'didResolveOperation',
                ctx,
              }),
            )
          }
        },

        executionDidStart(ctx) {
          if (enableExecutionDidStart) {
            winstonInstance.log(debug, stringify({ id, event: 'executionDidStart', ctx }))
          }
        },

        parsingDidStart(ctx) {
          if (enableParsingDidStart) {
            winstonInstance.log(debug, stringify({ id, event: 'parsingDidStart', ctx }))
          }
        },

        responseForOperation(ctx) {
          if (enableResponseForOperation) {
            winstonInstance.log(
              debug,
              stringify({
                id,
                event: 'responseForOperation',
                ctx,
              }),
            )
          }
          return null
        },

        validationDidStart(ctx) {
          if (enableValidationDidStart) {
            winstonInstance.log(debug, stringify({ id, event: 'validationDidStart', ctx }))
          }
        },

        willSendResponse({ response }: { response: { data: any } }) {
          if (enableWillSendResponse) {
            winstonInstance.log(
              debug,
              stringify({
                id,
                event: 'response',
                response: response.data,
              }),
            )
          }
        },
      }
      return handlers
    },
  }
}

export default apolloWinstonLoggingPlugin
