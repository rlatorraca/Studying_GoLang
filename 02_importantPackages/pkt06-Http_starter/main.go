package main

import (
	"fmt"
	"log"
	"net/http"
)

/* Manipulando Arquivos em GoLang*/
func main() {

	http.HandleFunc("/", getCEP) //
	err := http.ListenAndServe(":8900", nil)
	if err != nil {
		log.Fatal(fmt.Printf("[ERROR] Listing HTTP server:8900 ... : %v\n", err))
	} // Cria um servidor HTTP na porta :8900
}

func getCEP(response http.ResponseWriter, request *http.Request) {
	_, err := response.Write([]byte("Hello World: getCEP()"))
	if err != nil {
		log.Fatal(fmt.Printf("[ERROR] getCEP function ... : %v\n", err))
	}
}
