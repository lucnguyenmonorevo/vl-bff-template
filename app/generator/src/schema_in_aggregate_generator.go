package src

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"vl-template/app/generator"
)

func (s *srcGenerator) generateSchemaInAggregates() error {
	customConfMap := make(map[string]any)
	customConfMap["import_schemas"] = s.getSchemaInAggregatesImport()
	customConfMap["type_defs"] = s.getSchemaInAggregatesTypeDefs()
	customConfMap["schemas_to_type_defs"] = s.convertSchemaToTypeDef()
	customFuncMap := make(map[string]any)
	gen, err := generator.NewGenerator(s.Domains[0], customConfMap, customFuncMap)
	if err != nil {
		return err
	}
	dirPath := "vl-bff-{{ ToKebab .service_name }}/src/aggregates"
	fileName := "schemas.ts"
	if err := gen.GenerateFileByFile(dirPath, fileName, fileName, false); err != nil {
		return err
	}
	return nil
}

func (s *srcGenerator) getSchemaInAggregatesImport() string {
	rt := ""
	for _, domain := range s.Domains {
		aggregatePathName := strcase.ToSnake(domain.AggregateName)
		domainPathName := strcase.ToSnake(domain.DomainName)
		domainName := strcase.ToLowerCamel(domain.DomainName)
		str := fmt.Sprintf(`import { %sSchema } from './%s/%s/schema'
`, domainName, aggregatePathName, domainPathName)
		rt += str
	}
	return rt
}

func (s *srcGenerator) getSchemaInAggregatesTypeDefs() string {
	rt := ""
	comma := `,`
	for i, domain := range s.Domains {
		if i == len(s.Domains)-1 {
			comma = `,
	`
		}
		domainName := strcase.ToLowerCamel(domain.DomainName)
		str := fmt.Sprintf(`
	%sTypeDef%s`, domainName, comma)
		rt += str
	}
	return rt
}

func (s *srcGenerator) convertSchemaToTypeDef() string {
	rt := ""
	for _, domain := range s.Domains {
		domainName := strcase.ToLowerCamel(domain.DomainName)
		str := fmt.Sprintf(`const %sTypeDef = gql(%sSchema)
`, domainName, domainName)
		rt += str
	}
	return rt
}
