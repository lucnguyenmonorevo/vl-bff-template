package src

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"vl-template/app/generator"
)

func (s *srcGenerator) generateResolverInAggregates() error {
	customConfMap := make(map[string]any)
	customConfMap["import_resolvers"] = s.getResolverInAggregatesImport()
	customConfMap["resolvers"] = s.getResolverInAggregatesResolvers()
	customFuncMap := make(map[string]any)
	gen, err := generator.NewGenerator(s.Domains[0], customConfMap, customFuncMap)
	if err != nil {
		return err
	}
	dirPath := "vl-bff-{{ ToKebab .service_name }}/src/aggregates"
	fileName := "resolvers.ts"
	if err := gen.GenerateFileByFile(dirPath, fileName, fileName, false); err != nil {
		return err
	}
	return nil
}

func (s *srcGenerator) getResolverInAggregatesImport() string {
	rt := ""
	for _, domain := range s.Domains {
		aggregatePathName := strcase.ToSnake(domain.AggregateName)
		domainPathName := strcase.ToSnake(domain.DomainName)
		domainName := strcase.ToLowerCamel(domain.DomainName)
		str := fmt.Sprintf(`import %sResolvers from './%s/%s/resolver'
`, domainName, aggregatePathName, domainPathName)
		rt += str
	}
	return rt
}

func (s *srcGenerator) getResolverInAggregatesResolvers() string {
	rt := ""
	comma := `,
`
	for i, domain := range s.Domains {
		if i == len(s.Domains)-1 {
			comma = `,`
		}
		domainName := strcase.ToLowerCamel(domain.DomainName)
		str := fmt.Sprintf(`	...%sResolvers%s`, domainName, comma)
		rt += str
	}
	return rt
}
