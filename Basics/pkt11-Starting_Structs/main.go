package main

import "fmt"

// GoLang NAO eh OO
// GoLang se assemela as OO languages ao usar Structs

// STRUCTS
// NAO é possivel HERANCA
// É possivel COMPOSICAO de Structs
type Client struct {
	name   string
	age    int
	active bool
}

func main() {

	client01 := Client{
		name:   "Jorge",
		age:    22,
		active: true,
	}

	fmt.Printf("Name: %s | Age: %d | Active: %t\n", client01.name, client01.age, client01.active)

	client01.active = false

	fmt.Printf("Name: %s | Age: %d | Active: %t", client01.name, client01.age, client01.active)

}


