package main

// Importando um pacote
import (
	"fmt"
)

type inteiro int

var (
	var11 inteiro = 100
	var12 float32 = 1.33
)

func main() {
	var myArray [6]uint
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30
	myArray[3] = 40
	myArray[4] = 50
	myArray[5] = 60

	fmt.Println(len(myArray)) // qtas posicoes
	fmt.Println(myArray[1])
	fmt.Println(myArray[len(myArray)-1]) // ultima posicao

	for i, v := range myArray {
		fmt.Printf("Index value: %d | Value: %d\n", i, v)
	}

	var myArrayString [6]string
	myArrayString[0] = "Numero 01"
	myArrayString[1] = "Numero 02"
	myArrayString[2] = "Numero 03"
	myArrayString[3] = "Numero 04"
	myArrayString[4] = "Numero 05"
	myArrayString[5] = "Numero 06"

	for i, v := range myArrayString {
		fmt.Printf("Index value: %d | String value : %v\n", i, v)
	}
}
