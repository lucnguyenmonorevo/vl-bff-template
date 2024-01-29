package src

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"vl-template/app/generator"
)

func (s *srcGenerator) generateCodegen() error {
	customConfMap := make(map[string]any)
	customConfMap["code_gen_config"] = s.getCodeGenConfig()
	customFuncMap := make(map[string]any)
	gen, err := generator.NewGenerator(s.Domains[0], customConfMap, customFuncMap)
	if err != nil {
		return err
	}
	dirPath := "vl-bff-{{ ToKebab .service_name }}"
	fileName := "codegen.ts"
	if err := gen.GenerateFileByFile(dirPath, fileName, fileName, false); err != nil {
		return err
	}
	return nil
}

func (s *srcGenerator) getCodeGenConfig() string {
	rt := ""
	comma := `,
`
	for i, domain := range s.Domains {
		if i == len(s.Domains)-1 {
			comma = ","
		}
		aggregateName := strcase.ToSnake(domain.AggregateName)
		domainName := strcase.ToSnake(domain.DomainName)
		str := fmt.Sprintf(`    './src/aggregates/%s/%s/graphql_generated/graphql.ts': {
      schema: './src/aggregates/%s/%s/graphql/*',
      plugins: ['typescript', 'typescript-resolvers', 'typescript-document-nodes'],
      config: {
        useIndexSignature: true,
      },
    }%s`, aggregateName, domainName, aggregateName, domainName, comma)
		rt += str
	}
	return rt
}
