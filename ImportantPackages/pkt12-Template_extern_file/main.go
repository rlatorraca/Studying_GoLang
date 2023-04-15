package main

import (
	"log"
	"os"
	"text/template"
)

type Course struct {
	Name  string
	Hours int
}

type Courses []Course

/* Manipulation Arquivos em GoLang*/
func main() {
	var myCourses Courses
	myCourses = []Course{
		{"GoLang Basic", 20},
		{"GoLang Intermediate", 30},
		{"GoLang Advanced", 40},
	}

	/* Usando Template.Must para tratar erros */
	/* Define a slice of template file paths */

	templateFilePaths := []string{
		"./ImportantPackages/pkt12-Template_extern_file/public/index.html",
	}
	tmpl := template.Must(template.New("index.html").ParseFiles(templateFilePaths...))

	err := tmpl.Execute(os.Stdout, myCourses)
	if err != nil {
		log.Fatal("[ERROR] Executing Template: ", err)
	}

}
