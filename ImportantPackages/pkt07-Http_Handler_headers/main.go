package main

import (
	"fmt"
	"log"
	"net/http"
)

/* Manipulando Arquivos em GoLang*/
func main() {

	http.HandleFunc("/", getCEPHandler) //
	err := http.ListenAndServe(":8900", nil)
	if err != nil {
		log.Fatal(fmt.Printf("[ERROR] Listing HTTP server:8900 ... : %v\n", err))
	} // Cria um servidor HTTP na porta :8900
}

func getCEPHandler(response http.ResponseWriter, request *http.Request) {

	if request.URL.Path != "/" {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	cepParameter := request.URL.Query().Get("cep")
	if cepParameter == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	_, err := response.Write([]byte("Hello World: getCEPHandler()"))
	if err != nil {
		log.Fatal(fmt.Printf("[ERROR] getCEPHandler function ... : %v\n", err))
	}
}
