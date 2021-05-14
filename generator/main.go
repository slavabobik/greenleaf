package main

import (
	"bytes"
	"log"
	"os"
	"strings"
	"text/template"
)

type operatorType int

const (
	scalar = operatorType(1)
	slice  = operatorType(2)
)

type operator struct {
	name  string
	value string
	typ   operatorType
}

var findOperators = []operator{
	{"Eq", "$eq", scalar},
	{"Ne", "$ne", scalar},
	{"Gt", "$gt", scalar},
	{"Gte", "$gte", scalar},
	{"Lt", "$lt", scalar},
	{"Lte", "$lte", scalar},
	{"In", "$in", slice},
	{"Nin", "$nin", slice},
}

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
	tmpl := template.Must(template.ParseFiles("./generator/filter.tmpl"))

	file, err := os.Create("filter.go")
	if err != nil {
		return err
	}

	var filterData []templateData
	for _, op := range findOperators {

		// generate methods for builtin types.
		for _, typ := range builtins {
			data := templateData{
				Method:   op.name + strings.Title(typ),
				Operator: op.value,
			}

			switch op.typ {
			case scalar:
				data.Type = typ
			case slice:
				data.Type = "[]" + typ
			}

			filterData = append(filterData, data)
		}

		// generate methods for time data type.
		data := templateData{
			Method:   op.name + "Time",
			Operator: op.value,
			Type:     "time.Time",
		}
		filterData = append(filterData, data)
	}

	tmpl.Execute(file, filterData)
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
		for _, typ := range builtins {
			method := op + strings.Title(typ)
			s := templateData{
				Method:   method,
				Type:     typ,
				Operator: operator,
			}

			tmpl.Execute(&buf, s)

			s.Type = "[]" + typ
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
