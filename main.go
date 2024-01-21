package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	confMap := map[string]any{}
	confFunc := map[string]any{}
	confFunc["Func"] = func(r string) string {
		return fmt.Sprintf("func %s", r)
	}
	confMap["Count"] = "aa"
	confMap["Material"] = "aa"
	confMap["funcTest"] = "funcaa"
	tmpl, err := template.New("test").
		Funcs(confFunc).
		Parse(`{{ .Count }} items are made of {{ .Material }}
			Function: {{Func .funcTest}}
`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, confMap)
	if err != nil {
		panic(err)
	}
}
