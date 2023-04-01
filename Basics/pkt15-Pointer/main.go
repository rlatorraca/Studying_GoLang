package main

import "fmt"

// POINTER

func main() {

	a := 10

	/* & : traz o endereo de memoria */
	/* * : traz o valor */

	// var01 *int = &var02 : var01(int) APONTA para o end de memoria (&var02)

	fmt.Printf("Valor de 'a': %d\n", a)                   // VALOR
	fmt.Printf("Endereco de memoria de '&a': %d\n\n", &a) // Endereco de memoria

	var ponteiro *int = &a

	fmt.Printf("Valor para 'ponteiro': %d\n", ponteiro)     // Endereco de memoria
	fmt.Printf("Valor para '&ponteiro': %d\n", &ponteiro)   // Endereco de memoria
	fmt.Printf("Valor para '*ponteiro': %d\n\n", *ponteiro) // Valor presente no end de memoria

	*ponteiro = 20 // Muda o valor para 20
	fmt.Printf("Valor para '*ponteiro': %d\n", *ponteiro)

	b := &a // b APONTA para A
	c := a  // c COPIA o VALOR de A
	fmt.Printf("Valor de 'b': %d\n", b)
	fmt.Printf("Valor de '*b': %d\n", *b)

	fmt.Printf("Valor de 'c': %d\n", c)
	fmt.Printf("Valor de '&c': %d\n\n", &c)

	*b = 100 // A posicao de memoria de B ( = a ) muda o valor para 100
	c = *b

	fmt.Printf("Valor de 'b': %d\n", b)
	fmt.Printf("Valor de '*b': %d\n", *b)

	fmt.Printf("Valor de 'c': %d\n", c)

}
