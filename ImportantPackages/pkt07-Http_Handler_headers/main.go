package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

	http.HandleFunc("/", getCEPHandler)      //
	err := http.ListenAndServe(":8900", nil) // Multiplexer com o padrao do GoLang = nil (nao temos controler)
	if err != nil {
		log.Fatal(fmt.Printf("[ERROR] Listing HTTP server:8900 ... : %v\n", err))
	} // Cria um servidor HTTP na porta :8900
}

func getCEPHandler(response http.ResponseWriter, request *http.Request) {

	if request.URL.Path != "/" {
		response.WriteHeader(http.StatusNotFound) /* 404 */
		return
	}

	cepParameter := request.URL.Query().Get("cep")
	if cepParameter == "" {
		response.WriteHeader(http.StatusBadRequest) /* 400 */
		return
	}

	returnedCEP, err := getCEP(cepParameter)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError) /* 500 */
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	/* USING Marshal ~ NewEnconder
	jsonCEP, err := json.Marshal(returnedCEP)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError) / 500 /
		return
	}
	finalResponse, err := response.Write(jsonCEP)
	if err != nil {
		return
	}
	*/

	err = json.NewEncoder(response).Encode(returnedCEP)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError) /* 500 */
		return
	}

	/* PRINT on screen a text
	_, err := response.Write([]byte("Hello World: getCEPHandler()"))
	if err != nil {
		log.Fatal(fmt.Printf("[ERROR] getCEPHandler function ... : %v\n", err))
	}
	*/
}

func getCEP(cep string) (*ViaCep, error) {

	response, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(fmt.Printf("[ERROR] Closing body... : %v\n", err))
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var viaCEP ViaCep

	err = json.Unmarshal(body, &viaCEP)
	if err != nil {
		return nil, err
	}

	return &viaCEP, nil
}
