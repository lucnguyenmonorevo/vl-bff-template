{{ $_ := . -}}
import { readFileSync } from 'fs'
import path from 'path'

const {{ ToLowerCamel $_.domain_name }}FilePath = path.resolve(__dirname, 'graphql/{{ ToSnake .domain_name }}.schemas.graphql')

export const {{ ToLowerCamel $_.domain_name }}Schema = readFileSync({{ ToLowerCamel $_.domain_name }}FilePath, { encoding: 'utf-8' })
