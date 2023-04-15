package main

import (
	"fmt"
	"log"
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

	/* Usando Template.Must para tratar erros */

	tmpl := template.Must(template.New("Template Course").Parse("Course: {{.Name}} - Hours: {{.Hours}}"))

	for _, value := range courses {
		err := tmpl.Execute(os.Stdout, value)
		fmt.Printf("\n")
		if err != nil {
			log.Fatal("[ERROR] Executing Template: ", err)
		}
	}
}
