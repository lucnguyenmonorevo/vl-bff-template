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
	"vl-template/app/generator/util"
)

type DomainProtoReader struct {
	DomainsMap map[string][]*Domain
}

func NewDomainProtoReader() *DomainProtoReader {
	return &DomainProtoReader{
		DomainsMap: map[string][]*Domain{},
	}
}

func (r *DomainProtoReader) GetDomains(protoPath string) error {
	dirs, err := os.ReadDir(protoPath)
	if err != nil {
		return err
	}
	for _, dir := range dirs {
		dirName := dir.Name()
		servicePath := path.Join(protoPath, dirName)
		domains, err := r.getDomainsByService(servicePath)
		if err != nil {
			return err
		}
		r.DomainsMap[dirName] = domains
	}
	r.cleanData()
	return nil
}

func (r *DomainProtoReader) getDomainsByService(servicePath string) ([]*Domain, error) {
	domains := make([]*Domain, 0)
	domainPaths, err := r.getDomainPaths(servicePath)
	if err != nil {
		return nil, err
	}
	for _, domainPath := range domainPaths {
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
	relatedPaths, err := r.getImportPath(filePath)

	// Iterate through each line
	for _, line := range lines {
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

				requests, err := r.getRelatedPayload(relatedPaths, requestName)
				if err != nil {
					return nil, err
				}
				responses, err := r.getRelatedPayload(relatedPaths, responseName)
				if err != nil {
					return nil, err
				}
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
	}, nil
}

func getLastSplit(str string, split string) string {
	arr := strings.Split(str, split)
	return arr[len(arr)-1]
}

func (r *DomainProtoReader) getImportPath(filePath string) ([]string, error) {
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
		// Check if the line contains a message definition
		messageMatch := importRegex.FindStringSubmatch(line)
		if len(messageMatch) > 0 {
			rt = append(rt, "app/domain/"+messageMatch[1])
		}
	}
	return rt, nil
}

func (r *DomainProtoReader) getRelatedPayload(paths []string, payloadName string) ([]*Payload, error) {
	payloads := make([]*Payload, 0)
	for _, path := range paths {
		importPaths, err := r.getImportPath(path)
		if err != nil {
			importPaths = []string{}
		}
		importPaths = append(importPaths, path)
		tmpPayloads, err := r.getGrpc(path, payloadName)
		if err != nil {
			return nil, err
		}
		payloads = append(payloads, tmpPayloads...)

		for _, payload := range payloads {
			for _, field := range payload.Data {
				if !util.IsProtoType(field.Type) {
					addedPayloads, err := r.getRelatedPayload(importPaths, field.Type)
					if err != nil {
						return nil, err
					}
					payloads = append(payloads, addedPayloads...)
				}
			}
		}
	}
	return payloads, nil
}

func (r *DomainProtoReader) getGrpc(filePath string, payloadName string) ([]*Payload, error) {
	payloads := make([]*Payload, 0)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	// Define regular expressions for extracting message names and fields
	messageNameRegex := regexp.MustCompile(`message (\w+) {`)
	fieldRegex := regexp.MustCompile(`(\w+)\s+(\w+)\s*=\s*(\d+);`)

	// Split the file content by lines
	lines := strings.Split(string(content), "\n")

	// Variables to keep track of the current message
	var currentMessageName string
	var inMessage bool

	payload := &Payload{
		Data: []*PayloadData{},
	}
	// Iterate through each line
	for _, line := range lines {
		// Check if the line contains a message definition
		messageMatch := messageNameRegex.FindStringSubmatch(line)
		if len(messageMatch) > 0 {
			currentMessageName = messageMatch[1]
			if currentMessageName != payloadName {
				continue
			}
			inMessage = true
			payload.PayloadName = currentMessageName
			continue
		}

		// Check if the line contains a field definition
		if inMessage {
			fieldMatch := fieldRegex.FindStringSubmatch(line)
			if len(fieldMatch) > 0 {
				fieldType := fieldMatch[1]
				fieldName := fieldMatch[2]
				if fieldName == "base" {
					continue
				}
				payload.Data = append(payload.Data, &PayloadData{
					Name:      fieldName,
					Type:      fieldType,
					IsNotNull: false,
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
	}
	return payloads, nil
}

func getServiceFromPath(path string) (string, string) {
	serviceName := ""
	aggregateName := ""
	arr := strings.Split(path, "/")
	for i := range arr {
		switch arr[i] {
		case "proto":
			serviceName = arr[i+1]
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
	payloadMapRequest := map[string]bool{}
	payloadMapResponse := map[string]bool{}
	//aggregatePayloadMap := map[string]map[string]bool{}
	for _, domains := range r.DomainsMap {
		for _, domain := range domains {
			for _, queryMethod := range domain.QueryMethods {
				indices := make([]int, 0)
				for i, request := range queryMethod.Requests {
					if payloadMapRequest[request.PayloadName] {
						indices = append(indices, i)
						continue
					}
					payloadMapRequest[request.PayloadName] = true
				}
				if len(indices) > 0 {
					queryMethod.Requests = deleteElements(queryMethod.Requests, indices)
				}

				indices = make([]int, 0)
				for i, response := range queryMethod.Responses {
					if payloadMapResponse[response.PayloadName] {
						indices = append(indices, i)
						continue
					}
					payloadMapResponse[response.PayloadName] = true
				}
				if len(indices) > 0 {
					queryMethod.Responses = deleteElements(queryMethod.Responses, indices)
				}
			}

			for _, mutationMethod := range domain.MutationMethods {
				indices := make([]int, 0)
				for i, request := range mutationMethod.Requests {
					if payloadMapRequest[request.PayloadName] {
						indices = append(indices, i)
						continue
					}
					payloadMapRequest[request.PayloadName] = true
				}
				if len(indices) > 0 {
					mutationMethod.Requests = deleteElements(mutationMethod.Requests, indices)
				}
				indices = make([]int, 0)
				for i, response := range mutationMethod.Responses {
					if payloadMapResponse[response.PayloadName] {
						indices = append(indices, i)
						continue
					}
					payloadMapResponse[response.PayloadName] = true
				}

				if len(indices) > 0 {
					mutationMethod.Responses = deleteElements(mutationMethod.Responses, indices)
				}
			}
		}

	}
}
func deleteElements(slice []*Payload, indices []int) []*Payload {
	// Sort indices in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(indices)))

	for _, index := range indices {
		slice = append(slice[:index], slice[index+1:]...)
	}

	return slice
}
