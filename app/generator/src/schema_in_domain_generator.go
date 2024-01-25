package src

import (
	"vl-template/app/domain"
	"vl-template/app/generator"
)

func (s *srcGenerator) generateSchemaInDomain(domain *domain.Domain) error {
	customConfMap := make(map[string]any)
	customFuncMap := make(map[string]any)
	gen, err := generator.NewGenerator(domain, customConfMap, customFuncMap)
	if err != nil {
		return err
	}
	dirPath := "vl-{{ ToSnake .service_name }}/src/aggregates/{{ ToSnake .aggregate_name }}/{{ ToSnake .domain_name }}"
	fileName := "schema.ts"
	if err := gen.GenerateFileByFile(dirPath, fileName, fileName, false); err != nil {
		return err
	}
	return nil
}
