package src

import (
	"vl-template/app/domain"
	"vl-template/app/generator"
)

type SchemaGenerator interface {
	Generate() error
}
type srcGenerator struct {
	Domains []*domain.Domain
}

func NewSrcGenerator(domains []*domain.Domain) SchemaGenerator {
	return &srcGenerator{Domains: domains}
}

func (s *srcGenerator) Generate() error {
	if len(s.Domains) == 0 {
		return nil
	}
	for _, domain := range s.Domains {
		if err := s.generate(domain); err != nil {
			return err
		}
		if err := s.generateSchemaInDomain(domain); err != nil {
			return err
		}
	}
	if err := s.generateSchemaInAggregates(); err != nil {
		return err
	}
	if err := s.generateUtils(); err != nil {
		return err
	}
	if err := s.generateCodegen(); err != nil {
		return err
	}
	return nil
}

func (s *srcGenerator) generate(domain *domain.Domain) error {
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
