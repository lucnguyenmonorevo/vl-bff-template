package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc/reflect/protoreflect"
	"google.golang.org/grpc/reflect/protoregistry"
)

func main() {
	// Load your Proto file and register it with the Protobuf registry
	protoFile := "path/to/your/proto/file.proto"
	fileDescriptor, err := loadFileDescriptor(protoFile)
	if err != nil {
		log.Fatal(err)
	}

	// Use protoreflect to inspect the file descriptor
	for _, service := range fileDescriptor.GetServices() {
		fmt.Printf("Service Name: %s\n", service.GetName())

		// Iterate over service methods
		serviceDesc := service.GetMethods()
		for _, methodDesc := range serviceDesc {
			fmt.Printf("  Method Name: %s\n", methodDesc.GetName())
		}

		// Iterate over service messages (requests and responses)
		messageTypes := service.GetMethodTypes()
		for _, messageType := range messageTypes {
			fmt.Printf("  Message Type: %s\n", messageType)
		}
	}
}

func loadFileDescriptor(protoFile string) (protoreflect.FileDescriptor, error) {
	// Load the Proto file descriptor using protoregistry
	var fileDescriptor protoreflect.FileDescriptor
	if err := protoregistry.GlobalFiles.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		if fd.Path() == protoFile {
			fileDescriptor = fd
			return false // Stop iterating once we find the file
		}
		return true
	}); err != nil {
		return nil, fmt.Errorf("failed to load proto file descriptor: %v", err)
	}

	if fileDescriptor == nil {
		return nil, fmt.Errorf("proto file not found: %s", protoFile)
	}

	return fileDescriptor, nil
}
