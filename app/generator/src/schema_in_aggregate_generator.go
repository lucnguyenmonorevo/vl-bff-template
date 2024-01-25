package src

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"vl-template/app/domain"
	"vl-template/app/generator"
)

func (s *srcGenerator) generateSchemaInAggregates(domain []*domain.Domain) error {
	customConfMap := make(map[string]any)
	customConfMap["import_schemas"] = s.getSchemaInAggregatesImport()
	customConfMap["type_defs"] = s.getSchemaInAggregatesTypeDefs()
	customFuncMap := make(map[string]any)
	gen, err := generator.NewGenerator(domain[0], customConfMap, customFuncMap)
	if err != nil {
		return err
	}
	dirPath := "vl-{{ ToSnake .service_name }}/src/aggregates"
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
		str := fmt.Sprintf(`import { %sTypeDefs } from './%s/%s/schema'
`, domainName, aggregatePathName, domainPathName)
		rt += str
	}
	return rt
}

func (s *srcGenerator) getSchemaInAggregatesTypeDefs() string {
	rt := ""
	comma := ", "
	for i, domain := range s.Domains {
		if i == len(s.Domains)-1 {
			comma = ""
		}
		domainName := strcase.ToLowerCamel(domain.DomainName)
		str := fmt.Sprintf(`%sTypeDefs%s`, domainName, comma)
		rt += str
	}
	return rt
}
