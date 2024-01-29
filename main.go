package main

import (
	"fmt"
	"strings"
)

//import "fmt"

func main() {
	line := `//import "fmt"`
	fmt.Println(strings.HasPrefix(strings.TrimSpace(line), "//"))
}
