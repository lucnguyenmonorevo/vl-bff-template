package graphql

import (
	"vl-template/app/domain"
	"vl-template/app/generator"
)

type SchemaGenerator interface {
	Generate() error
}
type schemaGenerator struct {
	Domains []*domain.Domain
}

func NewSchemaGenerator(domains []*domain.Domain) SchemaGenerator {
	return &schemaGenerator{Domains: domains}
}

func (s *schemaGenerator) Generate() error {
	for _, domain := range s.Domains {
		if err := s.generate(domain); err != nil {
			return err
		}
	}
	return nil
}

func (s *schemaGenerator) generate(domain *domain.Domain) error {
	customConfMap := make(map[string]any)
	customFuncMap := make(map[string]any)
	gen, err := generator.NewGenerator(domain, customConfMap, customFuncMap)
	if err != nil {
		return err
	}
	dirPath := "vl-{{ .service_name }}/src/aggregates/{{ .aggregate_name }}/{{ .domain_name }}/graphql"
	fileName := "{{ .domain_name }}.schemas.graphql"
	if err := gen.GenerateFileByFile(dirPath, fileName, fileName, false); err != nil {
		return err
	}
	return nil
}
