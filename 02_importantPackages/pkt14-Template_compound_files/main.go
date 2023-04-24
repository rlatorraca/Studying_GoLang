package main

import (
	"log"
	"net/http"
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
		"./ImportantPackages/pkt14-Template_compound_files/public/header.html",
		"./ImportantPackages/pkt14-Template_compound_files/public/data.html",
		"./ImportantPackages/pkt14-Template_compound_files/public/footer.html",
	}

	/* Sera chamado data.html, que tem a composicao do arquivo html com header + footer */
	tmpl := template.Must(template.New("data.html").ParseFiles(templateFilePaths...))

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		err := tmpl.Execute(response, myCourses)
		if err != nil {
			log.Fatal("[ERROR] WebServer: ", err)
		}
	})

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("[ERROR] WebServer: ", err)
	}

}
