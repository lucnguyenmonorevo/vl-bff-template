package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path"
	"syscall"
	"vl-template/app/config"
	"vl-template/app/domain"
	"vl-template/app/generator/graphql"
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
	specDir := "app/domain/spec"
	domainReader := NewDomainReader(conf)
	if err = domainReader.Read(specDir); err != nil {
		log.Fatalln("can't read domain reader")
	}

	for _, domains := range domainReader.domainsMap {
		if err := generateSchema(domains); err != nil {
			panic(err)
		}
	}

	fmt.Println(".")

}

type DomainReader struct {
	conf       *config.Config
	domainsMap map[string][]*domain.Domain
}

func NewDomainReader(conf *config.Config) *DomainReader {
	return &DomainReader{
		conf:       conf,
		domainsMap: map[string][]*domain.Domain{},
	}

}

func (r *DomainReader) Read(dir string) error {
	if err := r.walkDir(dir); err != nil {
		return err
	}
	return nil
}

func (r *DomainReader) walkDir(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		fileName := file.Name()
		if file.IsDir() {
			childDir := path.Join(dir, fileName)
			r.walkDir(childDir)
		} else {
			dm, err := domain.NewDomain(dir, fileName)
			if err != nil {
				return err
			}

			r.domainsMap[dm.ServiceName] = append(r.domainsMap[dm.ServiceName], dm)
		}
	}
	return nil
}

func generateSchema(req []*domain.Domain) error {
	schemaGenerator := graphql.NewSchemaGenerator(req)
	if err := schemaGenerator.Generate(); err != nil {
		return err
	}
	return nil
}
