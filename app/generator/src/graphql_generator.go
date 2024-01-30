package src

import (
	"vl-template/app/domain"
	"vl-template/app/generator"
)

func (s *srcGenerator) generateGraphql(domain *domain.Domain) error {
	customConfMap := make(map[string]any)
	customFuncMap := make(map[string]any)
	gen, err := generator.NewGenerator(domain, customConfMap, customFuncMap)
	if err != nil {
		return err
	}
	dirPath := "vl-bff-{{ ToKebab .service_name }}/src/aggregates/{{ ToSnake .aggregate_name }}/{{ ToSnake .domain_name }}/graphql"
	fileName := "{{ ToSnake .domain_name }}.schemas.graphql"
	if err := gen.GenerateFileByFile(dirPath, fileName, fileName, false); err != nil {
		return err
	}
	return nil
}
