package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"vl-template/app/config"
	"vl-template/app/domain"
	"vl-template/app/generator/src"
)

func main() {
	_, stop := signal.NotifyContext(context.Background(),
		syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGQUIT, os.Interrupt)
	defer func() {
		stop()
	}()

	fmt.Println("start")

	_, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("proto generating...")

	fmt.Println("cleaning...")
	if err := os.RemoveAll("generated/"); err != nil {
		panic(err)
	}

	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	domainReader := domain.NewDomainReader(conf)
	if err = domainReader.ReadFromProto(); err != nil {
		log.Fatalln("can't read domain reader")
	}

	for _, domains := range domainReader.DomainsMap {
		if err := generateSchema(domains); err != nil {
			panic(err)
		}
	}

	fmt.Println(".")

}

func generateSchema(req []*domain.Domain) error {
	schemaGenerator := src.NewSchemaGenerator(req)
	if err := schemaGenerator.Generate(); err != nil {
		return err
	}
	return nil
}
