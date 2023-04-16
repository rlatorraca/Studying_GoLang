package main

import (
	"fmt"
	"os"
	"text/template"
)

type Course struct {
	Name  string
	Hours int
}

/* Manipulation Arquivos em GoLang*/
func main() {
	courses := []Course{
		{"GoLang Basic", 20},
		{"GoLang Intermediate", 30},
		{"GoLang Advanced", 40},
	}
	tmpl := template.New("CourseTemplate")
	tmpl, _ = tmpl.Parse("Course: {{.Name}} - Hours: {{.Hours}}")

	for _, value := range courses {
		err := tmpl.Execute(os.Stdout, value)
		fmt.Printf("\n")
		if err != nil {
			panic(err)
		}
	}
}
