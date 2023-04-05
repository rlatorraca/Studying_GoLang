package main

import (
	"fmt"
	"log"
	"os"
)

/* Manipulando Arquivos em GoLang*/
func main() {

	data := make([]byte, 100)

	// Criando arquivo
	f, err := os.Create("HandlingFiles.txt")
	if err != nil {
		panic(err)
	}

	f, err = os.Open("HandlingFiles.txt") // For read access.
	if err != nil {
		log.Fatal(fmt.Printf("[ERROR] Opening... : %v\n", err))
	}

	count, err := os.ReadFile("HandlingFiles.txt")
	if err != nil {
		log.Fatal(fmt.Printf("[ERROR] Reading... : %v\n", err))

	}
	fmt.Printf("read %b bytes: %q\n", count, data)

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
