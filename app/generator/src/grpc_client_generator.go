package src

import (
	"vl-template/app/domain"
	"vl-template/app/generator"
)

func (s *srcGenerator) generateGrpcClient(domainInput *domain.Domain) error {
	methods := make([]*domain.Method, 0)
	methods = append(methods, domainInput.QueryMethods...)
	methods = append(methods, domainInput.MutationMethods...)
	for _, method := range methods {
		customConfMap := make(map[string]any)
		customConfMap["name"] = method.Name
		customConfMap["request_name"] = method.RequestName
		customConfMap["response_name"] = method.ResponseName
		customConfMap["request_path"] = s.getRequestPath(method)
		customConfMap["response_path"] = s.getResponsePath(method)
		customConfMap["requests"] = method.Requests
		customConfMap["responses"] = method.Responses

		customFuncMap := make(map[string]any, 0)
		gen, err := generator.NewGenerator(domainInput, customConfMap, customFuncMap)
		if err != nil {
			return err
		}

		dirPath := "vl-bff-{{ ToKebab .service_name }}/src/aggregates/{{ ToSnake .aggregate_name }}/{{ ToSnake .domain_name }}/grpc_clients"
		fileName := "{{ ToSnake .domain_name }}_{{ ToSnake .name }}.ts"
		if err := gen.GenerateFileByFile(dirPath, fileName, fileName, false); err != nil {
			return err
		}
	}
	return nil
}

func (s *srcGenerator) getRequestPath(method *domain.Method) string {
	for _, payload := range method.Requests {
		if payload.Name != method.RequestName {
			continue
		}
		return payload.Path
	}
	return ""
}

func (s *srcGenerator) getResponsePath(method *domain.Method) string {
	for _, payload := range method.Responses {
		if payload.Name != method.ResponseName {
			continue
		}
		return payload.Path
	}
	return ""
}
