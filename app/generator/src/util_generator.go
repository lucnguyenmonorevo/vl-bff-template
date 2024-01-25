package src

import (
	"vl-template/app/domain"
	"vl-template/app/generator"
)

func (s *srcGenerator) generateUtils(domain []*domain.Domain) error {
	customConfMap := make(map[string]any)
	customFuncMap := make(map[string]any)
	gen, err := generator.NewGenerator(domain[0], customConfMap, customFuncMap)
	if err != nil {
		return err
	}
	dirPath := "vl-{{ ToSnake .service_name }}/src/utils"
	fileName := "const.ts"
	if err := gen.GenerateFileByFile(dirPath, fileName, fileName, false); err != nil {
		return err
	}
	dirPath = "vl-{{ ToSnake .service_name }}/src/utils"
	fileName = "errorHandler.ts"
	if err := gen.GenerateFileByFile(dirPath, fileName, fileName, false); err != nil {
		return err
	}
	dirPath = "vl-{{ ToSnake .service_name }}/src/utils"
	fileName = "logger.ts"
	if err := gen.GenerateFileByFile(dirPath, fileName, fileName, false); err != nil {
		return err
	}
	return nil
}
