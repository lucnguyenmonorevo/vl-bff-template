package copy

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"vl-template/app/generator/util"
)

func (r *GeneratedFileCopier) CopyProto(srcDir string, destDir string) error {
	r.overwriteAllFiles = false
	r.notOverwriteAllFiles = false

	if err := r.walkProtoDir(srcDir, srcDir, destDir); err != nil {
		return err
	}
	return nil
}

func (r *GeneratedFileCopier) walkProtoDir(srcRootDir string, srcDir string, destDir string) error {
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
			if err := r.walkProtoDir(
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

			// Replace the old string with the new string
			content, err := io.ReadAll(srcFile)
			if err != nil {
				return err
			}
			newContent := strings.Replace(string(content), `import "proto/`, `import "`, -1)

			destFile, err := os.Create(destFilePath)
			if err != nil {
				return err
			}
			defer destFile.Close()

			_, err = destFile.WriteString(newContent)
			if err != nil {
				return err
			}

		}
	}
	return nil
}
