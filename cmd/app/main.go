package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"vl-template/app/config"
	"vl-template/app/copy"
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
		if err := generateSrc(domains); err != nil {
			panic(err)
		}
	}

	// Copy data
	copier := copy.NewGeneratedFileCopier(conf)
	if conf.CopyTo != "" {
		fmt.Println("copying...")
		generatedDir := "generated/"
		if err := copier.Copy(generatedDir, conf.CopyTo); err != nil {
			panic(err)
		}
	}
	if err := copyProtos(copier); err != nil {
		panic(err)
	}

	fmt.Println(".")

}

func generateSrc(req []*domain.Domain) error {
	srcGenerator := src.NewSrcGenerator(req)
	if err := srcGenerator.Generate(); err != nil {
		return err
	}
	return nil
}

func copyProtos(copier *copy.GeneratedFileCopier) error {
	copyMap := map[string]string{
		"vl-account":      "vl-bff-account",
		"vl-business":     "vl-bff-business",
		"vl-general-info": "vl-bff-general-info",
		"vl-notification": "vl-bff-notification",
		"vl-physical-obj": "vl-bff-physical-obj",
		"vl-procurement":  "vl-bff-procurement",
		"vl-production":   "vl-bff-production",
	}
	for from, to := range copyMap {
		if err := copier.CopyProto(from+"/proto", to+"/proto"); err != nil {
			return err
		}
		if err := copier.CopyProto("vl-backend-basement/proto", to+"/proto"); err != nil {
			return err
		}
	}
	return nil
}
