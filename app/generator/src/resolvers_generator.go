package src

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"vl-template/app/domain"
	"vl-template/app/generator"
	"vl-template/app/generator/util"
)

func (s *srcGenerator) generateResolvers(domainInput *domain.Domain) error {
	customConfMap := make(map[string]any)
	customConfMap["import_requests"] = s.generateResolversImportRequests(domainInput)
	customConfMap["import_convert"] = s.generateResolversImportConvert(domainInput)
	customConfMap["resolvers"] = s.generateResolversResolvers(domainInput)
	customConfMap["resolver_ts_import"] = s.generateResolverTsImport(domainInput)
	customConfMap["resolver_ts_resolvers"] = s.generateResolverTsQueries(domainInput)
	customFuncMap := make(map[string]any, 0)
	gen, err := generator.NewGenerator(domainInput, customConfMap, customFuncMap)
	if err != nil {
		return err
	}
	dirPath := "vl-bff-{{ ToKebab .service_name }}/src/aggregates/{{ ToSnake .aggregate_name }}/{{ ToSnake .domain_name }}/resolvers"
	fileName := ""
	if len(domainInput.MutationMethods) > 0 {
		fileName = "mutations.ts"
		if err := gen.GenerateFileByFile(dirPath, fileName, fileName, false); err != nil {
			return err
		}
	}
	if len(domainInput.QueryMethods) > 0 {
		fileName = "queries.ts"
		if err := gen.GenerateFileByFile(dirPath, fileName, fileName, false); err != nil {
			return err
		}
	}

	dirPath = "vl-bff-{{ ToKebab .service_name }}/src/aggregates/{{ ToSnake .aggregate_name }}/{{ ToSnake .domain_name }}"
	fileName = "resolver.ts"
	if err := gen.GenerateFileByFile(dirPath, fileName, fileName, false); err != nil {
		return err
	}
	return nil
}

func (s *srcGenerator) generateResolversImportRequests(domainInput *domain.Domain) string {
	rt := ""
	comma := `,
`
	for i, method := range domainInput.MutationMethods {
		if i == len(s.Domains)-1 {
			comma = ","
		}
		str := fmt.Sprintf(`	%s%s%s`, util.GetUpperCaseChars(domainInput.DomainName), strcase.ToCamel(method.RequestName), comma)
		rt += str
	}
	return rt
}

func (s *srcGenerator) generateResolversImportConvert(domainInput *domain.Domain) string {
	rt := ""
	comma := `
`
	for i, method := range domainInput.MutationMethods {
		if i == len(s.Domains)-1 {
			comma = ""
		}
		str := fmt.Sprintf(`import { %s } from '../grpc_clients/%s'%s`,
			strcase.ToSnake(domainInput.DomainName)+"_"+strcase.ToSnake(method.Name),
			strcase.ToSnake(domainInput.DomainName)+"_"+strcase.ToSnake(method.Name),
			comma)
		rt += str
	}
	return rt
}

func (s *srcGenerator) generateResolversResolvers(domainInput *domain.Domain) string {
	rt := ""
	comma := `,
`
	for i, method := range domainInput.MutationMethods {
		if i == len(s.Domains)-1 {
			comma = ","
		}
		str := fmt.Sprintf(`
    %s%s: async (
      _: any,
      params: { input: %s%s },
      base: BaseRequest,
    ) => {
      const res = await %s(base, params.input)
      return res
    }%s`, strcase.ToLowerCamel(domainInput.DomainName),
			strcase.ToCamel(method.Name),
			util.GetUpperCaseChars(domainInput.DomainName),
			strcase.ToCamel(method.RequestName),
			strcase.ToSnake(domainInput.DomainName)+"_"+strcase.ToSnake(method.Name),
			comma)
		rt += str
	}
	return rt
}

func (s *srcGenerator) generateResolverTsImport(domainInput *domain.Domain) string {
	rt := ""
	if len(domainInput.MutationMethods) > 0 {
		rt += `import mutations from './resolvers/mutations'
`
	}
	if len(domainInput.QueryMethods) > 0 {
		rt += `import queries from './resolvers/queries'
`
	}
	return rt
}

func (s *srcGenerator) generateResolverTsQueries(domainInput *domain.Domain) string {
	rt := ""
	coma := ""
	if len(domainInput.QueryMethods) > 0 {
		rt += `...queries`
		coma = ", "
	}
	if len(domainInput.MutationMethods) > 0 {
		rt += coma + `...mutations`
	}
	return rt
}
