package main

import "fmt"

// POINTER

func main() {

	a := 10

	fmt.Printf("Valor de 'a': %d\n", a)                   // VALOR
	fmt.Printf("Endereco de memoria de '&a': %d\n\n", &a) // Endereco de memoria

	var ponteiro *int = &a // * => APONTA para o end de memoria (&)

	fmt.Printf("Valor para 'ponteiro': %d\n", ponteiro)     // Endereco de memoria
	fmt.Printf("Valor para '&ponteiro': %d\n", &ponteiro)   // Endereco de memoria
	fmt.Printf("Valor para '*ponteiro': %d\n\n", *ponteiro) // Valor presente no end de memoria

	*ponteiro = 20
	fmt.Printf("Valor para '*ponteiro': %d\n", *ponteiro)

	b := &a
	c := a
	fmt.Printf("Valor de 'b': %d\n", b)
	fmt.Printf("Valor de '*b': %d\n", *b)

	fmt.Printf("Valor de 'c': %d\n", c)
	fmt.Printf("Valor de '&c': %d\n\n", &c)

	*b = 100
	c = *b

	fmt.Printf("Valor de 'b': %d\n", b)
	fmt.Printf("Valor de '*b': %d\n", *b)

	fmt.Printf("Valor de 'c': %d\n", c)

}
