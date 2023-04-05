package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/* Manipulando Arquivos em GoLang*/
func main() {

	// Criando arquivo
	f, err := os.Create("HandlingFiles02.txt")
	if err != nil {
		panic(err)
	}

	fileSize, err := f.Write([]byte("Hello, World")) // Escrever Bytes
	//fileSize, err := f.WriteString("Hello, World") // Escrever um STRING
	if err != nil {
		log.Fatal(fmt.Printf("[ERROR] Opening... : %v\n", err))
	}

	fmt.Printf("File Size: %d bytes\n", fileSize)

	err = f.Close()
	if err != nil {
		panic(err)
	}

	// abrir e ler arquivo
	file, err := os.ReadFile("HandlingFiles02.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))

	// Abrir, Ler em partes o arquivo (em Buffer)
	fileByPart, err := os.Open("HandlingFiles02.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(fileByPart)
	buffer := make([]byte, 1) // le em 100 em 100 bytes

	for {
		b, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:b]))
	}

	// Remove File
	err = os.Remove("HandlingFiles02.txt")
	if err != nil {
		panic(err)
	}
}
