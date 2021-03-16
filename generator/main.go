package main

import (
	"bytes"
	"log"
	"os"
	"strings"
	"text/template"
)

var findSingleOperators = []string{"$eq", "$ne", "$gt", "$gte", "$lt", "$lte"}
var findSliceOperators = []string{"$in", "$nin"}
var builtins = []string{
	"bool",
	"float32",
	"float64",
	"int",
	"int16",
	"int32",
	"int64",
	"int8",
	"string",
	"uint",
	"uint16",
	"uint32",
	"uint64",
	"uint8",
}

func main() {

	filterFuncTemplate := `
// {{.Method}} ...	
func (f *FilterBuilder) {{.Method}}(field string, value {{.Type}}) *FilterBuilder {
	return f.addSelector(field, "{{.Operator}}", value)
}
	`

	tmpl, err := template.New("filter_func").Parse(filterFuncTemplate)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("filter.go")
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	for _, operator := range findSingleOperators {
		for _, t := range builtins {

			method := strings.Title(operator[1:]) + strings.Title(t)
			s := struct {
				Method, Type, Operator string
			}{
				Method:   method,
				Type:     t,
				Operator: operator,
			}

			tmpl.Execute(&buf, s)
		}
	}

	for _, operator := range findSliceOperators {
		for _, t := range builtins {

			method := strings.Title(operator[1:]) + strings.Title(t)
			s := struct {
				Method, Type, Operator string
			}{
				Method:   method,
				Type:     "[]" + t,
				Operator: operator,
			}

			tmpl.Execute(&buf, s)
		}
	}

	file.WriteString("// Code generated. DO NOT EDIT.\n")
	file.WriteString("package greenleaf\n\n")
	file.WriteString(buf.String())
}
