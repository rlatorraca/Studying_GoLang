package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Account struct {
	AccountNumber int     `json:"num"` // tag
	Balance       float64 `json:"bal" validate:"max=5000"`
}

/* Manipulando Arquivos em GoLang*/
func main() {

	// MANEIRA SIMPLES => Dentro de um variavel
	account := Account{AccountNumber: 12312, Balance: 11231.33}
	request, err := json.Marshal(account)
	if err != nil {
		log.Fatal("[ERROR] Reading JSON...")
	}

	fmt.Println(string(request))

	// RECEBE o valor e faz o ENCODER para gravar em outro lugar (arquivo, tela)
	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		log.Fatal("[ERROR] ENCODE JSON...")
	}

	// Transformar JSON em Struct
	jsonRaw := []byte(`{"num": 333, "bal": 34343.44}`)
	var account2 Account
	err = json.Unmarshal(jsonRaw, &account2)
	if err != nil {
		return
	}

	fmt.Printf("Account: %v | Balance: %v", account2.AccountNumber, account2.Balance)
}
