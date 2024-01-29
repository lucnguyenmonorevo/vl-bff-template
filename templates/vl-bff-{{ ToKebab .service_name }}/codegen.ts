{{ $_ := . -}}
import type { CodegenConfig } from '@graphql-codegen/cli'

const config: CodegenConfig = {
  overwrite: true,
  emitLegacyCommonJSImports: false,
  generates: {
    {{ $_.code_gen_config }}
  },
}

export default config
