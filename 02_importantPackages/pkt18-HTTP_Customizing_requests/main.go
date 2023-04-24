package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

/* Manipulation Arquivos em GoLang*/
func main() {

	// Meu  HTTP Client
	clientHTTP := http.Client{Timeout: 2 * time.Second} // 2 segundos é o tempo de espera

	// Um Request Object
	request, err := http.NewRequest("GET", "http://www.cbc.ca", nil)
	if err != nil {
		log.Fatal("[ERROR] Getting Request Object.. ", err)
	}

	// Customizando o Reequest
	request.Header.Add("Accept", "application/json")
	request.Header.Set("Content-Language", "en-US")

	// CLiente HTTP, recebe a request modificada e o response
	response, err := clientHTTP.Do(request)
	if err != nil {
		log.Fatal("[ERROR] Sendin Http Request.. ", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal("[ERROR] Closing Body.. ", err)
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)

	// Lendo o Body
	println(string(body))
}
