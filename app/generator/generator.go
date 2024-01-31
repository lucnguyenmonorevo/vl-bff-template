package generator

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/iancoleman/strcase"
	"os"
	"path"
	"strings"
	"text/template"
	"vl-template/app/domain"
	"vl-template/app/generator/util"
)

type Generator struct {
	dm      *domain.Domain
	confMap map[string]any
	funcMap map[string]any
}

func NewGenerator(
	mdl *domain.Domain,
	customConfMap map[string]any,
	customFuncMap map[string]any,
) (*Generator, error) {
	confBytes, err := json.Marshal(mdl)
	if err != nil {
		return nil, fmt.Errorf("failed marshal json")
	}
	var confMap map[string]any
	if err := json.Unmarshal(confBytes, &confMap); err != nil {
		return nil, fmt.Errorf("failed unmarshal json")
	}

	// handle Columns
	//columns := confMap["columns"].([]any)
	//for i := 0; i < len(columns); i++ {
	//	column := columns[i].(map[string]any)
	//	length := column["length"].(float64)
	//	column["length"] = int(length)
	//	columns[i] = column
	//	confMap["columns"] = columns
	//}
	confMap = mergeMaps(confMap, customConfMap)
	// update funcMap
	funcMap := template.FuncMap{
		"ToGraphql":         util.ToGraphql,
		"ToLowerCamel":      strcase.ToLowerCamel,
		"ToCamel":           strcase.ToCamel,
		"ToSnake":           strcase.ToSnake,
		"ToKebab":           strcase.ToKebab,
		"GetUpperCaseChars": util.GetUpperCaseChars,
		"IsProtoType":       util.IsProtoType,
		"ToUpper":           strings.ToUpper,
		"ToUpperSnake":      util.ToUpperSnake,
	}
	funcMap = mergeMaps(funcMap, customFuncMap)

	return &Generator{
		dm:      mdl,
		confMap: confMap,
		funcMap: funcMap,
	}, nil
}

func mergeMaps(a, b map[string]any) map[string]any {
	out := make(map[string]any, len(a))
	for k, v := range a {
		out[k] = v
	}
	for k, v := range b {
		// If you use map[string]any, ok is always false here.
		// Because yaml.Unmarshal will give you map[any]any.
		if v, ok := v.(map[string]any); ok {
			if bv, ok := out[k]; ok {
				if bv, ok := bv.(map[string]any); ok {
					out[k] = mergeMaps(bv, v)
					continue
				}
			}
		}
		out[k] = v
	}
	return out
}

func (g *Generator) GenerateStringByString(text string) (string, error) {
	tpl := template.New("")
	tpl.Funcs(g.funcMap)
	_tpl, err := tpl.Parse(text)
	if err != nil {
		return "", err
	}

	resultBuf := new(bytes.Buffer)
	resultWriter := bufio.NewWriter(resultBuf)
	if err := _tpl.Execute(resultWriter, g.confMap); err != nil {
		return "", err
	}
	resultWriter.Flush()
	return resultBuf.String(), nil
}

func (g *Generator) GenerateStringByFile(dirPath string, inFileName string) (string, error) {
	inDirPath := path.Join("templates", dirPath)
	inFilePath := path.Join(inDirPath, inFileName)

	tpl := template.New(inFileName)
	tpl.Funcs(g.funcMap)
	_tpl, err := tpl.ParseFiles(inFilePath)
	if err != nil {
		return "", err
	}

	resultBuf := new(bytes.Buffer)
	resultWriter := bufio.NewWriter(resultBuf)
	if err := _tpl.Execute(resultWriter, g.confMap); err != nil {
		return "", err
	}
	resultWriter.Flush()
	return resultBuf.String(), nil
}

func (g *Generator) GenerateFileByString(dirPath string, outFileName string, text string, append bool) error {
	var err error
	outDirPath := path.Join("generated", dirPath)
	outDirPath, err = g.GenerateStringByString(outDirPath)
	if err != nil {
		return err
	}
	outFileName, err = g.GenerateStringByString(outFileName)
	if err != nil {
		return err
	}
	outFilePath := path.Join(outDirPath, outFileName)
	if util.ExistFile(outFilePath) {
		return fmt.Errorf("- already file: %s", outFilePath)
	}
	if err := util.Mkdirs(outDirPath); err != nil {
		return err
	}
	fmt.Println("- create file:", outFilePath)
	if err := util.WriteFile(outFilePath, []byte(text)); err != nil {
		return err
	}
	return nil
}

func (g *Generator) GenerateFileByFile(dirPath string, inFileName string, outFileName string, append bool) error {
	inDirPath := path.Join("templates", dirPath)
	inFilePath := path.Join(inDirPath, inFileName)

	var err error
	outDirPath := path.Join("generated", dirPath)
	outDirPath, err = g.GenerateStringByString(outDirPath)
	if err != nil {
		return err
	}
	outFileName, err = g.GenerateStringByString(outFileName)
	if err != nil {
		return err
	}
	outFilePath := path.Join(outDirPath, outFileName)
	if util.ExistFile(outFilePath) {
		return fmt.Errorf("- already file: %s", outFilePath)
	}

	tpl := template.New(inFileName)
	tpl.Funcs(g.funcMap)
	_tpl, err := tpl.ParseFiles(inFilePath)
	if err != nil {
		return err
	}
	if err := util.Mkdirs(outDirPath); err != nil {
		return err
	}
	outFile, err := os.Create(outFilePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	fmt.Println("- create file:", outFilePath)
	outFileWriter := bufio.NewWriter(outFile)
	if err := _tpl.Execute(outFileWriter, g.confMap); err != nil {
		return err
	}
	outFileWriter.Flush()
	return nil
}
