package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
	"vl-template/app/config"
	"vl-template/app/domain"
	"vl-template/app/generator/src"
	"vl-template/app/generator/util"
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

	// Copy data
	if conf.CopyTo != "" {
		fmt.Println("copying...")
		generatedDir := "generated/"
		copier := NewGeneratedFileCopier(conf)
		if err := copier.Copy(generatedDir, conf.CopyTo); err != nil {
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

type GeneratedFileCopier struct {
	conf *config.Config

	overwriteAllFiles    bool
	notOverwriteAllFiles bool
}

func NewGeneratedFileCopier(conf *config.Config) *GeneratedFileCopier {

	return &GeneratedFileCopier{
		conf: conf,
	}
}

func (r *GeneratedFileCopier) Copy(srcDir string, destDir string) error {
	r.overwriteAllFiles = false
	r.notOverwriteAllFiles = false

	if err := r.dirwalk(srcDir, srcDir, destDir); err != nil {
		return err
	}
	return nil
}

func (r *GeneratedFileCopier) dirwalk(srcRootDir string, srcDir string, destDir string) error {
	files, err := os.ReadDir(srcDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		fileName := file.Name()

		//if len(r.conf.CopyIncludeRepositories) != 0 {
		//	ignoreFilePathTemp := filepath.Join(srcDir, fileName)
		//	ignoreFilePath := strings.Replace(ignoreFilePathTemp, srcRootDir, "", 1)
		//	serviceName := strings.Split(ignoreFilePath, "/")[0]
		//	repositoryName := util.SnakeCaseToKebabCase(serviceName)
		//	if !slices.Contains(r.conf.CopyIncludeRepositories, repositoryName) {
		//		fmt.Println("- ignore directory:", ignoreFilePathTemp)
		//		continue
		//	}
		//}

		// log.Println("path:", fileName)
		if file.IsDir() {
			if err := r.dirwalk(
				srcRootDir,
				filepath.Join(srcDir, fileName),
				destDir,
			); err != nil {
				return err
			}
		} else {
			// log.Println("file", fileName)

			if !strings.Contains(fileName, ".ts") &&
				!strings.Contains(fileName, ".proto") &&
				!strings.Contains(fileName, ".graphql") &&
				!strings.Contains(fileName, ".yml") &&
				!strings.Contains(fileName, ".yaml") {
				continue
			}

			srcFilePath := filepath.Join(srcDir, fileName)

			destDirPart := strings.Replace(srcDir, srcRootDir, "", 1)
			destDirPath := filepath.Join(destDir, destDirPart)
			if err := util.Mkdirs(destDirPath); err != nil {
				return err
			}
			destFilePath := filepath.Join(destDirPath, fileName)

			if !r.overwriteAllFiles && r.conf.ConfirmOverwrite && util.ExistFile(destFilePath) {
				confirm := false

				if r.conf.ConfirmAllFiles {
					confirm = true
				} else {
					ignoreLaunchRegexp := regexp.MustCompile("app/aggregates/[^/]+/launch.go$")
					if ignoreLaunchRegexp.MatchString(srcFilePath) {
						confirm = true
					}
					// if strings.Contains(srcFilePath, "commands/migrations/mysql") {
					// 	confirm = true
					// }
					if strings.Contains(srcFilePath, "app/includes") {
						confirm = true
					}
				}

				if confirm {
					if r.notOverwriteAllFiles {
						fmt.Println("- ignore file:", srcFilePath)
						continue
					} else {
						fmt.Println("- exists file:", destFilePath)
						fmt.Println("  overwrite? (ALLYES/allno/Y/n)")
						scanner := bufio.NewScanner(os.Stdin)
						scanner.Scan()
						scannedText := scanner.Text()

						if scannedText == "ALLYES" {
							r.overwriteAllFiles = true
						} else if scannedText == "allno" {
							r.notOverwriteAllFiles = true
							fmt.Println("- ignore file:", srcFilePath)
							continue
						} else if scannedText != "Y" {
							fmt.Println("- ignore file:", srcFilePath)
							continue
						}
					}
				}
			}

			srcFile, err := os.Open(srcFilePath)
			if err != nil {
				return err
			}
			defer srcFile.Close()

			destFile, err := os.Create(destFilePath)
			if err != nil {
				return err
			}
			defer destFile.Close()

			_, err = io.Copy(destFile, srcFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
