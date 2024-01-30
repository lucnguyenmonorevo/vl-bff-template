package src

import (
	"vl-template/app/domain"
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
		if err := s.generateGraphql(domain); err != nil {
			return err
		}
		if err := s.generateSchemaInDomain(domain); err != nil {
			return err
		}
		if err := s.generateGrpcClient(domain); err != nil {
			return err
		}
		if err := s.generateResolvers(domain); err != nil {
			return err
		}
	}
	if err := s.generateSchemaInAggregates(); err != nil {
		return err
	}
	if err := s.generateResolverInAggregates(); err != nil {
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
