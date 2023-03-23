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
	fmt.Printf("Type of var11 is: %T\n", var11)
	fmt.Printf("Type of var11 is: %t\n", var11)
	fmt.Printf("Value of var11 is: %v\n", var11)
	fmt.Println("--------------------------------------")
	fmt.Printf("Type of var12 is: %T\n", var12)
	fmt.Printf("Type of var12 is: %t\n", var12)
	fmt.Printf("Value of var12 is: %v\n", var12)
}
