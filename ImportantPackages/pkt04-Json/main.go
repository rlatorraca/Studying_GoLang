package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

/* Manipulando Arquivos em GoLang*/
func main() {

	request, err := http.Get("https://www.terra.com.br")

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(request.Body)

	if err != nil {
		log.Fatal(fmt.Printf("[ERROR] Getting site... : %v\n", err))
	}

	response, err := io.ReadAll(request.Body)
	if err != nil {
		log.Fatal(fmt.Printf("[ERROR] Reading Site Body... : %v\n", err))
	}

	fmt.Printf(string(response))

}
