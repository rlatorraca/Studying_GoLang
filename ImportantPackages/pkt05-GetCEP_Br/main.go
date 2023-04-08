package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

/* Manipulando Arquivos em GoLang*/
func main() {

	// Armazena o que for digitado no prompt de comando ao fazor o go run main.go XXXX
	for _, cep := range os.Args[1:] {
		request, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "[ERROR] Requesting URL: %v\n", err)
		}

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "[ERROR] Closing Request: %v\n", err)
			}
		}(request.Body)

		response, err := io.ReadAll(request.Body)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "[ERROR] Reading response body: %v\n", err)
		}

		var data ViaCep

		// Traduz o JSON pego na internet para a STRUCT do Golang ViaCep
		err = json.Unmarshal(response, &data)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "[ERROR] Parsing response to Struct:  %v\n", err)
		}

		// Salvando dentro de um arquvivo
		file, err := os.Create("cepInfo.txt")

		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "[ERROR] Closing file 'cepInfo.txt':  %v\n", err)
			}
		}(file)

		// Escrevendo dentro do arquivo
		_, err = file.WriteString(fmt.Sprintf("CEP: %s, City: %v, UF: %v, Rue/Street: %v, Neighbourhood: %v",
			data.Cep, data.Localidade, data.Uf, data.Logradouro, data.Bairro))
		fmt.Printf("File saved successfully")
	}

}
