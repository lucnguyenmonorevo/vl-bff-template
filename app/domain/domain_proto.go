package domain

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"
	"vl-template/app/config"
	"vl-template/app/generator/util"
)

type DomainProtoReader struct {
	Config     *config.Config
	DomainsMap map[string][]*Domain
}

func NewDomainProtoReader() *DomainProtoReader {
	return &DomainProtoReader{
		DomainsMap: map[string][]*Domain{},
	}
}

func (r *DomainProtoReader) GetDomains(dirs map[string]string) error {
	for serviceName, dir := range dirs {
		domains, err := r.getDomainsByService(dir)
		if err != nil {
			return err
		}
		r.DomainsMap[serviceName] = domains
	}
	r.cleanData()
	r.updateDomain()
	return nil
}

func (r *DomainProtoReader) getDomainsByService(servicePath string) ([]*Domain, error) {
	domains := make([]*Domain, 0)
	domainPaths, err := r.getDomainPaths(servicePath)
	if err != nil {
		return nil, err
	}
	for _, domainPath := range domainPaths {
		if ok, err := r.isDomain(domainPath); !ok {
			continue
			if err != nil {
				fmt.Println("[Error] Err about domain: ", err)
			}
		}
		domain, err := r.getDomain(domainPath)
		if err != nil {
			return nil, err
		}
		domains = append(domains, domain)
	}
	return domains, nil
}

func (r *DomainProtoReader) getDomainPaths(dir string) ([]string, error) {
	rt := []string{}
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		fileName := file.Name()
		if file.IsDir() {
			/* to be updated */
			childDir := path.Join(dir, fileName)
			filePaths, err := r.getDomainPaths(childDir)
			if err != nil {
				return nil, err
			}
			rt = append(rt, filePaths...)
		} else {
			filePath := path.Join(dir, fileName)
			if strings.HasSuffix(filePath, "_service.proto") {
				rt = append(rt, filePath)
			}
		}
	}
	return rt, nil
}

func (r *DomainProtoReader) getDomain(filePath string) (*Domain, error) {
	queryMethods := make([]*Method, 0)
	mutationMethods := make([]*Method, 0)
	enums := make([]*Enum, 0)
	domainName := ""
	serviceName, aggregateName := getServiceFromPath(filePath)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	// Define regular expressions for extracting service and method details
	serviceRegex := regexp.MustCompile(`service (\w+) {`)
	methodRegex := regexp.MustCompile(`rpc (\w+)\(([^)]+)\) returns \(([^)]+)\);`)

	// Split the file content by lines
	lines := strings.Split(string(content), "\n")

	// Variables to keep track of the current service and method
	var currentServiceName string
	var inService bool
	relatedPaths, err := r.getImportPaths(filePath)

	// Iterate through each line
	for _, line := range lines {
		if r.isCommentLine(line) {
			continue
		}
		// Check if the line contains a service definition
		serviceMatch := serviceRegex.FindStringSubmatch(line)
		if len(serviceMatch) > 0 {
			currentServiceName = serviceMatch[1]
			inService = true
			domainName = strcase.ToSnake(strings.ReplaceAll(currentServiceName, "Service", ""))
			continue
		}

		// Check if the line contains a method definition
		if inService {
			methodMatch := methodRegex.FindStringSubmatch(line)
			if len(methodMatch) > 0 {
				methodName := methodMatch[1]
				requestType := strings.TrimSpace(methodMatch[2])
				requestName := getLastSplit(requestType, ".")
				responseType := strings.TrimSpace(methodMatch[3])
				responseName := getLastSplit(responseType, ".")

				requests, tmpEnums, err := r.getRelatedPayload(relatedPaths, requestName)
				if err != nil {
					return nil, err
				}
				r.updatePayloadType(requests, requestName, responseName)
				enums = append(enums, tmpEnums...)
				responses, tmpEnums, err := r.getRelatedPayload(relatedPaths, responseName)
				if err != nil {
					return nil, err
				}
				r.updatePayloadType(responses, requestName, responseName)
				enums = append(enums, tmpEnums...)
				if getMethodType(methodName) == "query" {
					queryMethods = append(queryMethods, &Method{
						Name:         strcase.ToSnake(methodName),
						RequestName:  requestName,
						ResponseName: responseName,
						Requests:     requests,
						Responses:    responses,
					})
				} else {
					mutationMethods = append(mutationMethods, &Method{
						Name:         strcase.ToSnake(methodName),
						RequestName:  requestName,
						ResponseName: responseName,
						Requests:     requests,
						Responses:    responses,
					})
				}
			}
		}

		// Check if the line ends the current service
		if strings.Contains(line, "}") && inService {
			inService = false
		}
	}
	return &Domain{
		ServiceName:     serviceName,
		AggregateName:   aggregateName,
		DomainName:      domainName,
		QueryMethods:    queryMethods,
		MutationMethods: mutationMethods,
		Enums:           enums,
	}, nil
}

func getLastSplit(str string, split string) string {
	arr := strings.Split(str, split)
	return arr[len(arr)-1]
}

func (r *DomainProtoReader) getImportPaths(filePath string) ([]string, error) {
	rt := make([]string, 0)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	// Define regular expressions for extracting message names and fields
	importRegex := regexp.MustCompile(`import "(.*?)";`)

	// Split the file content by lines
	lines := strings.Split(string(content), "\n")

	// Iterate through each line
	for _, line := range lines {
		if r.isCommentLine(line) {
			continue
		}
		// Check if the line contains a message definition
		messageMatch := importRegex.FindStringSubmatch(line)
		if len(messageMatch) > 0 {
			rt = append(rt, path.Join(strings.Split(filePath, "/")[0], messageMatch[1]))
		}
	}
	return rt, nil
}

func (r *DomainProtoReader) getRelatedPayload(paths []string, payloadName string) ([]*Payload, []*Enum, error) {
	payloads := make([]*Payload, 0)
	enums := make([]*Enum, 0)
	for _, path := range paths {
		if !r.isValidPath(path) {
			continue
		}
		importPaths, err := r.getImportPaths(path)
		if err != nil {
			importPaths = []string{}
		}
		importPaths = append(importPaths, path)
		tmpPayloads, tmpEnums, err := r.getGrpc(path, payloadName)
		if err != nil {
			return nil, nil, err
		}
		payloads = append(payloads, tmpPayloads...)
		enums = append(enums, tmpEnums...)

		for _, payload := range payloads {
			for _, field := range payload.Data {
				if !util.IsProtoType(field.Type) && field.Type != payloadName {
					addedPayloads, addEnums, err := r.getRelatedPayload(importPaths, field.Type)
					if err != nil {
						return nil, nil, err
					}
					payloads = append(payloads, addedPayloads...)
					enums = append(enums, addEnums...)
				}
			}
		}
	}
	return payloads, enums, nil
}

func (r *DomainProtoReader) getGrpc(filePath string, payloadName string) ([]*Payload, []*Enum, error) {
	payloads := make([]*Payload, 0)
	enums := make([]*Enum, 0)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, nil, err
	}

	// Define regular expressions for extracting message names and fields
	messageNameRegex := regexp.MustCompile(`message (\w+) {`)
	fieldRegex := regexp.MustCompile(`(\w+)\s+(\w+)\s*=\s*(\d+);`)

	// Define regular expressions for extracting enum names and fields
	enumNameRegex := regexp.MustCompile(`enum (\w+) {`)
	enumFieldRegex := regexp.MustCompile(`(\w+)\s*=\s*(\d+);`)

	// Split the file content by lines
	lines := strings.Split(string(content), "\n")

	// Variables to keep track of the current message
	var currentMessageName string
	var inMessage bool

	// Variables to keep track of the embedded message
	var currentEmbeddedMessageName string
	var inEmbeddedMessage bool

	// Variables to keep track of the current enum
	var currentEnumName string
	var inEnum bool

	payload := &Payload{
		Data: []*PayloadData{},
		Path: r.getImportPath(filePath),
	}

	embeddedPayload := &Payload{
		Data: []*PayloadData{},
		Path: r.getImportPath(filePath),
	}

	enum := &Enum{
		Values: []string{},
	}
	// Iterate through each line
	for _, line := range lines {
		if r.isCommentLine(line) {
			continue
		}
		/* Handle message */
		// Check if the line contains a message definition
		messageMatch := messageNameRegex.FindStringSubmatch(line)
		if len(messageMatch) > 0 {
			// check embedded message
			if inMessage {
				currentEmbeddedMessageName = messageMatch[1]
				inEmbeddedMessage = true
				embeddedPayload.Name = currentEmbeddedMessageName
				continue
			}
			currentMessageName = messageMatch[1]
			if currentMessageName != payloadName {
				continue
			}
			inMessage = true
			payload.Name = currentMessageName
			continue
		}
		if inEmbeddedMessage {
			fieldMatch := fieldRegex.FindStringSubmatch(line)
			if len(fieldMatch) > 0 {
				isArray := false
				fieldType := fieldMatch[1]
				fieldName := fieldMatch[2]
				if strings.Contains(line, "repeated") {
					isArray = true
				}
				if fieldName == "base" {
					continue
				}
				embeddedPayload.Data = append(embeddedPayload.Data, &PayloadData{
					Name:      fieldName,
					Type:      fieldType,
					IsNotNull: false,
					IsArray:   isArray,
				})
			}
			// Check if the line ends the current message
			if strings.Contains(line, "}") {
				inEmbeddedMessage = false
				if len(embeddedPayload.Data) == 0 {
					embeddedPayload.Data = append(embeddedPayload.Data, &PayloadData{
						Name:      "default_field",
						Type:      "string",
						IsNotNull: false,
					})
				}
				payloads = append(payloads, embeddedPayload)
				embeddedPayload = &Payload{
					Data: []*PayloadData{},
				}
			}
			// Skip when read embedded message

			continue
		}

		if inMessage {

			// Complete read embedded message

			fieldMatch := fieldRegex.FindStringSubmatch(line)
			if len(fieldMatch) > 0 {
				isArray := false
				fieldType := fieldMatch[1]
				fieldName := fieldMatch[2]
				if strings.Contains(line, "repeated") {
					isArray = true
				}
				if fieldName == "base" {
					continue
				}
				payload.Data = append(payload.Data, &PayloadData{
					Name:      fieldName,
					Type:      fieldType,
					IsNotNull: false,
					IsArray:   isArray,
				})
			}
		}

		// Check if the line ends the current message
		if strings.Contains(line, "}") && inMessage {
			inMessage = false
			if len(payload.Data) == 0 {
				payload.Data = append(payload.Data, &PayloadData{
					Name:      "default_field",
					Type:      "string",
					IsNotNull: false,
				})
			}
			payloads = append(payloads, payload)
			payload = &Payload{
				Data: []*PayloadData{},
			}
		}

		/* Handle enum */
		// Check if the line contains a message definition
		enumMatch := enumNameRegex.FindStringSubmatch(line)
		if len(enumMatch) > 0 {
			currentEnumName = enumMatch[1]
			if currentEnumName != payloadName {
				continue
			}
			inEnum = true
			enum.Name = currentEnumName
			continue
		}

		// Check if the line contains a field definition
		if inEnum {
			fieldMatch := enumFieldRegex.FindStringSubmatch(line)
			if len(fieldMatch) > 0 {
				fieldName := fieldMatch[1]
				enum.Values = append(enum.Values, fieldName)
			}
		}

		// Check if the line ends the current message
		if strings.Contains(line, "}") && inEnum {
			inEnum = false
			enums = append(enums, enum)
			enum = &Enum{
				Values: []string{},
			}
		}
	}
	return payloads, enums, nil
}

func getServiceFromPath(path string) (string, string) {
	serviceName := ""
	aggregateName := ""
	arr := strings.Split(path, "/")
	for i := range arr {
		switch arr[i] {
		case "proto":
			serviceName = strcase.ToKebab(arr[i+1])
		case "aggregates":
			aggregateName = arr[i+1]
		}
	}
	return serviceName, aggregateName
}

func getMethodType(methodName string) string {
	if strings.Contains(methodName, "List") || strings.Contains(methodName, "Get") {
		return "query"
	} else {
		return "mutation"
	}
}

func (r *DomainProtoReader) cleanData() {
	payloadRequestMap := map[string]bool{}
	payloadResponseMap := map[string]bool{}
	enumMap := map[string]bool{}
	//aggregatePayloadMap := map[string]map[string]bool{}
	for _, domains := range r.DomainsMap {
		for _, domain := range domains {
			serviceName := domain.ServiceName
			domainName := domain.DomainName
			for _, queryMethod := range domain.QueryMethods {
				indices := make([]int, 0)
				for i, request := range queryMethod.Requests {
					key := serviceName + domainName + request.Name
					if payloadRequestMap[key] {
						indices = append(indices, i)
						continue
					}
					payloadRequestMap[key] = true
				}
				if len(indices) > 0 {
					queryMethod.Requests = deletePayloadElements(queryMethod.Requests, indices)
				}

				indices = make([]int, 0)
				for i, response := range queryMethod.Responses {
					key := serviceName + domainName + response.Name
					if payloadResponseMap[key] {
						indices = append(indices, i)
						continue
					}
					payloadResponseMap[key] = true
				}
				if len(indices) > 0 {
					queryMethod.Responses = deletePayloadElements(queryMethod.Responses, indices)
				}
			}

			for _, mutationMethod := range domain.MutationMethods {
				indices := make([]int, 0)
				for i, request := range mutationMethod.Requests {
					key := serviceName + domainName + request.Name
					if payloadRequestMap[key] {
						indices = append(indices, i)
						continue
					}
					payloadRequestMap[key] = true
				}
				if len(indices) > 0 {
					mutationMethod.Requests = deletePayloadElements(mutationMethod.Requests, indices)
				}
				indices = make([]int, 0)
				for i, response := range mutationMethod.Responses {
					key := serviceName + domainName + response.Name
					if payloadResponseMap[key] {
						indices = append(indices, i)
						continue
					}
					payloadResponseMap[key] = true
				}

				if len(indices) > 0 {
					mutationMethod.Responses = deletePayloadElements(mutationMethod.Responses, indices)
				}
			}

			enumIndices := make([]int, 0)
			for i, enum := range domain.Enums {
				key := serviceName + enum.Name
				if enumMap[key] {
					enumIndices = append(enumIndices, i)
					continue
				}
				enumMap[key] = true
			}
			if len(enumIndices) > 0 {
				domain.Enums = deleteEnumElements(domain.Enums, enumIndices)
			}
		}

	}
}

func deletePayloadElements(slice []*Payload, indices []int) []*Payload {
	// Sort indices in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(indices)))

	for _, index := range indices {
		slice = append(slice[:index], slice[index+1:]...)
	}

	return slice
}

func deleteEnumElements(slice []*Enum, indices []int) []*Enum {
	// Sort indices in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(indices)))

	for _, index := range indices {
		slice = append(slice[:index], slice[index+1:]...)
	}

	return slice
}

func (r *DomainProtoReader) isValidPath(path string) bool {
	if strings.Contains(path, "google/protobuf/timestamp.proto") {
		return false
	}
	if strings.Contains(path, "proto/common/grpc/grpc.proto") {
		return false
	}
	if strings.Contains(path, "common/messaging/messaging.proto") {
		return false
	}
	return true
}

func (r *DomainProtoReader) isDomain(path string) (bool, error) {
	// TODO: update ignore list
	ignorePaths := []string{
		"example",
	}
	for _, ignorePath := range ignorePaths {
		if strings.Contains(path, ignorePath) {
			return false, nil
		}
	}
	importPaths, err := r.getImportPaths(path)
	if err != nil {
		return false, err
	}
	for _, importPath := range importPaths {
		if strings.Contains(importPath, "/grpc/") {
			return true, nil
		}
	}
	return false, nil
}

func (r *DomainProtoReader) updateDomain() (bool, error) {
	requestType := "Input"
	responseType := "Type"
	for _, domains := range r.DomainsMap {
		for _, domain := range domains {
			for _, method := range domain.QueryMethods {
				requests := method.Requests
				responses := method.Responses
				for _, request := range requests {
					datas := request.Data
					for _, data := range datas {
						if !util.IsProtoType(data.Type) {
							data.Type = strcase.ToCamel(data.Type) + requestType
						}
					}
					request.Name = strcase.ToCamel(request.Name) + requestType
				}
				for _, response := range responses {
					datas := response.Data
					for _, data := range datas {
						if !util.IsProtoType(data.Type) {
							data.Type = strcase.ToCamel(data.Type) + responseType
						}
					}
					response.Name = strcase.ToCamel(response.Name) + responseType
				}
				method.RequestName = strcase.ToCamel(method.RequestName) + requestType
				method.ResponseName = strcase.ToCamel(method.ResponseName) + responseType
			}

			for _, method := range domain.MutationMethods {
				requests := method.Requests
				responses := method.Responses
				for _, request := range requests {
					datas := request.Data
					request.Name = strcase.ToCamel(request.Name) + requestType
					for _, data := range datas {
						if !util.IsProtoType(data.Type) {
							data.Type = strcase.ToCamel(data.Type) + requestType
						}
					}
				}
				for _, response := range responses {
					datas := response.Data
					response.Name = strcase.ToCamel(response.Name) + responseType
					for _, data := range datas {
						if !util.IsProtoType(data.Type) {
							data.Type = strcase.ToCamel(data.Type) + responseType
						}
					}
				}
				method.RequestName = strcase.ToCamel(method.RequestName) + requestType
				method.ResponseName = strcase.ToCamel(method.ResponseName) + responseType
			}

			//for _, enum := range domain.Enums {
			//	enum.Name = strcase.ToCamel(enum.Name)
			//}
		}
	}
	return false, nil
}

func (r *DomainProtoReader) isCommentLine(line string) bool {
	return strings.HasPrefix(strings.TrimSpace(line), "//")
}

func (r *DomainProtoReader) updatePayloadType(payloads []*Payload, requestName string, responseName string) {
	for _, payload := range payloads {
		switch payload.Name {
		case requestName:
			payload.Type = "request"
		case responseName:
			payload.Type = "response"
		default:
			payload.Type = ""
		}
	}
}

func (r *DomainProtoReader) getImportPath(pathName string) string {
	splits := strings.Split(pathName, "/proto/")
	if len(splits) != 2 {
		return ""
	}
	rt := strings.ReplaceAll(splits[1], ".proto", "")
	return rt
}
