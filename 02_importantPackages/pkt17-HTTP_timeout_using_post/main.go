package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

/* Manipulation Arquivos em GoLang*/
func main() {

	//callHTTP := http.Client{Timeout: 2 * time.Second} // 2 segundos é o tempo de espera
	clientHTTP := http.Client{Timeout: 2 * time.Microsecond} // 2 micro-segundos é o tempo de espera
	response, err := clientHTTP.Get("https://www.google.com")
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
