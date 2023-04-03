package main

import "fmt"

/* For in Golang */
// Em GoLang temos apenas o FOR (loop)

func main() {

	// FOR => SIMPLES
	for i := 0; i < 10; i++ {
		fmt.Printf("Item %d = %d\n", i, i)
	}

	fmt.Printf("=======================================\n")
	numbers := []string{"one", "two", "three", "four", "five"}

	//FOR => Key, Value
	for k, v := range numbers {
		fmt.Printf("Item %d = %v\n", k, v)
	}

	fmt.Printf("=======================================\n")
	for k, _ := range numbers { // Blank identifier
		fmt.Printf("Blank identifier (just keys) %v\n", k)
	}
	for _, v := range numbers { // Blank identifier
		fmt.Printf("Blank identifier (just values) %v\n", v)
	}

	fmt.Printf("=======================================\n")
	// FAKE WHILE
	i := 0
	for i < 10 {
		fmt.Printf("i is %d\t", i)
		i++
	}

	x := 100
	// INFINTIE LOOP
	for {
		fmt.Printf("Loop: %d\t", x)
	}
}
