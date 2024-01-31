package src

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"strings"
	"vl-template/app/domain"
	"vl-template/app/generator"
	"vl-template/app/generator/util"
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
		customConfMap["convert_request"] = s.convertRequest(method)
		customConfMap["convert_response"] = s.convertResponse(method)

		customFuncMap := make(map[string]any, 0)
		customFuncMap["convertImportGrpc"] = s.convertImportGRPC
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

func (s *srcGenerator) convertRequest(method *domain.Method) string {
	rt := ""
	for _, payload := range method.Requests {
		if payload.Name != method.RequestName {
			continue
		}
		for _, field := range payload.Data {
			if strcase.ToLowerCamel(field.Name) == "defaultField" {
				continue
			}
			list := ""
			if field.IsArray {
				list = "List"
			}
			if !util.IsProtoType(field.Type) {
				previousFiledPath := "graphqlReq"
				rt += s.convertRequestByPayLoad(previousFiledPath, method, field)
			}

			rt += fmt.Sprintf(`	rt.set%s%s(grpc%s)`,
				strcase.ToCamel(field.Name),
				list,
				strcase.ToCamel(field.Name))
		}
	}
	return rt
}

func (s *srcGenerator) convertRequestByPayLoad(previousFiledPath string, method *domain.Method, rootField *domain.PayloadData) string {
	rt := ""
	constVar := strcase.ToLowerCamel(fmt.Sprintf(`grpc%s`, strcase.ToCamel(rootField.Name)))
	rt += fmt.Sprintf(`const %s = rt.get%s()
	`,
		constVar,
		strcase.ToCamel(rootField.Name),
	)
	for _, payload := range method.Requests {
		if payload.Name != rootField.Type {
			continue
		}
		for _, field := range payload.Data {
			if strcase.ToLowerCamel(field.Name) == "defaultField" {
				continue
			}
			list := ""
			if field.IsArray {
				list = "List"
			}
			if util.IsProtoType(field.Type) {
				rt += fmt.Sprintf(`	%s.set%s%s(%s?.%s?.%s)
	`,
					constVar,
					strcase.ToCamel(field.Name),
					list,
					strcase.ToLowerCamel(previousFiledPath),
					strcase.ToLowerCamel(rootField.Name),
					strcase.ToLowerCamel(field.Name),
				)
			}
		}
	}
	return rt
}

func (s *srcGenerator) convertResponse(method *domain.Method) string {
	rt := ""
	for _, payload := range method.Responses {
		if payload.Name != method.ResponseName {
			continue
		}
		for _, field := range payload.Data {
			if strcase.ToLowerCamel(field.Name) == "defaultField" {
				continue
			}
			if !util.IsProtoType(field.Type) {
				previousFiledPath := "grpcRes"
				rt += s.convertResponseByPayLoad(previousFiledPath, method, field)
			}
		}
	}
	return rt
}

func (s *srcGenerator) convertResponseByPayLoad(previousFiledPath string, method *domain.Method, rootField *domain.PayloadData) string {
	rt := ""
	for _, payload := range method.Responses {
		if payload.Name != rootField.Type {
			continue
		}
		rt += fmt.Sprintf(`			%s: {
`, rootField.Name)
		for _, field := range payload.Data {
			if strcase.ToLowerCamel(field.Name) == "defaultField" {
				continue
			}
			list := ""
			if field.IsArray {
				list = "List"
			}
			if util.IsProtoType(field.Type) {
				rt += fmt.Sprintf(`				%s: %s?.%s?.%s%s,
`,
					strcase.ToLowerCamel(field.Name),
					strcase.ToLowerCamel(previousFiledPath),
					strcase.ToLowerCamel(rootField.Name),
					strcase.ToLowerCamel(field.Name),
					list,
				)
			}
		}
		rt += `			},`
	}
	return rt
}

func (s *srcGenerator) convertImportGRPC(grpc string) string {
	//rt := strcase.ToCamel(grpc)
	rt := strings.TrimSuffix(grpc, "Type")
	rt = strings.TrimSuffix(rt, "Input")
	return rt
}
