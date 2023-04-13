package main

import (
	"fmt"
	"log"
	"net/http"
)

type Blog struct {
	title    string
	language string
	date     string
}

/* Manipulation Arquivos em GoLang*/
func main() {

	// Cria um canal multiplexador PROPRIO para conectar as rotas/handlers (url's) no Web Server
	mux := http.NewServeMux()
	mux.HandleFunc("/site1", func(response http.ResponseWriter, request *http.Request) {
		_, err := response.Write([]byte("Hello Italy"))
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError) /* 500 */
			return
		}
	})
	/*err := http.ListenAndServe(":8765", mux)
	if err != nil {
		return
	}*/

	//mux2 := http.NewServeMux()
	mux.HandleFunc("/site2", HomeHandler)
	mux.Handle(
		"/site3",
		Blog{title: "RESP Blog", language: "Italian", date: "01/02/2023"},
	)

	err := http.ListenAndServe(":8888", mux)
	if err != nil {
		log.Fatal(fmt.Printf("[ERROR] Mux3 HTTP server:8765 ... : %v\n", err))
		return
	}
}

func HomeHandler(response http.ResponseWriter, request *http.Request) {
	_, err := response.Write([]byte("Hello Italy - Home Handler"))
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError) /* 500 */
		return
	}
}

/* Conecta a blog struct */
func (b Blog) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	_, err := writer.Write([]byte(b.title + "|" + b.language + "|" + b.date))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError) /* 500 */
		return
	}

}
