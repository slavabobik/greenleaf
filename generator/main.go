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
var updateMixedOperators = []string{"$set"}

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

const greenleafDoc = "M"

type templateData struct {
	Method, Type, Operator string
}

func main() {
	err := generateFilterBuilder()
	if err != nil {
		log.Fatal(err)
	}

	err = generateUpdateBuilder()
	if err != nil {
		log.Fatal(err)
	}

}

func generateFilterBuilder() error {

	filterFuncTemplate := `
// {{.Method}} ...	
func (f *FilterBuilder) {{.Method}}(field string, value {{.Type}}) *FilterBuilder {
	return f.addSelector(field, "{{.Operator}}", value)
}
`

	tmpl, err := template.New("filter_func").Parse(filterFuncTemplate)
	if err != nil {
		return err
	}

	file, err := os.Create("filter.go")
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	for _, operator := range findSingleOperators {
		for _, t := range builtins {

			method := strings.Title(operator[1:]) + strings.Title(t)
			s := templateData{
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
			s := templateData{
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

	return nil
}

func generateUpdateBuilder() error {
	updateFuncTemplate := `
// {{.Method}} ...	
func (u *UpdateBuilder) {{.Method}}(field string, value {{.Type}}) *UpdateBuilder {
	return u.addOperator("{{.Operator}}", field, value)
}
`
	tmpl, err := template.New("update_func").Parse(updateFuncTemplate)
	if err != nil {
		return err
	}

	file, err := os.Create("update.go")
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	for _, operator := range updateMixedOperators {

		op := strings.Title(operator[1:])
		for _, t := range builtins {
			method := op + strings.Title(t)
			s := templateData{
				Method:   method,
				Type:     t,
				Operator: operator,
			}

			tmpl.Execute(&buf, s)

			s.Type = "[]" + t
			s.Method = method + "Slice"

			tmpl.Execute(&buf, s)
		}

		method := op + "Document"
		s := templateData{
			Method:   method,
			Type:     greenleafDoc,
			Operator: operator,
		}
		tmpl.Execute(&buf, s)

	}

	file.WriteString("// Code generated. DO NOT EDIT.\n")
	file.WriteString("package greenleaf\n\n")
	file.WriteString(buf.String())

	return nil
}
