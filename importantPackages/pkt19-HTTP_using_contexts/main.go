package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

/*
CONTEXT é um pacote que permite que passamos suas informacoes para diferentes chamadas existentes no sistema, e que
permitem a parada em qualquer momento (para nao possamos perder tempo).
*/
func main() {

	// Cria um contexto vazio
	ctx := context.Background()

	// esse CONTEXTO sera cancelado passado 1 segundo ou quando chamada a funcao cancel (abaixo)
	ctx, cancel := context.WithTimeout(ctx, time.Second) // 1 Segundo ou time.Seoond
	//ctx, cancel := context.WithTimeout(ctx, time.Microsecond) // 1 microsegundo ou time.Microsecond

	defer cancel() // Roda o cancel quando o programa terminar

	// Cria uma request com o contexto criado
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.uol.com.br", nil)
	if err != nil {
		log.Fatal("[ERROR] Request Error.. ", err)
	}

	// Cria um cliente http default, faz a request com o contexto criado e joga para um response
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal("[ERROR] Response Error.. ", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal("[ERROR] Closing body ... ", err)
		}
	}(response.Body) // Fecha o body quando terminar

	// Imprime o response body
	body, err := io.ReadAll(response.Body)
	println(string(body))

}
