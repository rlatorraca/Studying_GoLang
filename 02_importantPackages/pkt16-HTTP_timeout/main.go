package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"time"
)

/* Manipulation Arquivos em GoLang*/
func main() {

	clientHTTP := http.Client{Timeout: 2 * time.Second} // 2 segundos é o tempo de espera

	json := bytes.NewBuffer([]byte(`{"nome":"Lucas"}{"nome":"Joao"}{"nome":"Marcos"}`))

	/* Usando POST */
	response, err := clientHTTP.Post("https://www.uol.com.br", "application/json", json)
	if err != nil {
		log.Fatal("[ERROR] Timeout Exceeded ...", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal("[ERROR] Closing body ...", err)
		}
	}(response.Body) // fecha o body da resposta

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("[ERROR] Reading body ...", err)
	}
	println(response.Status)
	println(string(body))
}
