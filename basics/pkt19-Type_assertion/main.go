package main

import (
	"fmt"
)

/* Type Assertion */

// Melhor usar generics dpo que interface vazia (empty interface)

func main() {

	var name interface{} = "John Marshall"
	var number interface{} = 100.00

	println(name)
	println(name.(string))

	result, ok := name.(int)
	fmt.Printf("Result: %v | Ok: %v\n", result, ok)

	println(number)
	println(number.(float64))

	result, ok = number.(int)
	fmt.Printf("Result: %v | Ok: %v", result, ok)

}
