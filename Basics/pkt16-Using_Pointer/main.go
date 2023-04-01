package main

import "fmt"

// VALORES MUTAVEIS => PONTEIRO
// VALORES IMUTAVEIS => Valor (e nao devemos usar PONTEIRO)

// SumByValue USING VALUES
func SumByValue(a, b int) int {
	return a + b
}

// SumByPointer USING POINTER
func SumByPointer(a, b *int) int {
	*a = *a + 12
	*b = *b + 3
	return *a + *b // retorna o valor da memoria
}

func main() {

	// Por valor
	var01 := 10
	var02 := 20
	fmt.Printf("Sum by value is %d\n", SumByValue(var01, var02))
	fmt.Printf("Var01 value is %d\n", var01)
	fmt.Printf("Var02 value is %d", var02)

	fmt.Println("\n===================================================")
	// Por referencia de memoria

	var03 := 30
	var04 := 40
	fmt.Printf("Sum by Pointer is %d\n", SumByPointer(&var03, &var04)) // Passa o ENDERECO DE MEMORIA e nao o valor
	fmt.Printf("Var03 value is %d\n", var03)
	fmt.Printf("Var04 value is %d", var04)
}
