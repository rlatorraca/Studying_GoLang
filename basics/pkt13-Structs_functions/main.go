package main

import "fmt"

// GoLang NAO eh OO
// GoLang se assemela as OO languages ao usar Structs

// STRUCTS
// NAO é possivel HERANCA
// É possivel COMPOSICAO de Structs

type Address struct {
	Street   string
	number   int8
	city     string
	province string
	country  string
}

type SIN struct {
	number int8
	year   uint16
}

type Client struct {
	name   string
	age    int
	active bool
	Address
	sin SIN
}

type Company struct {
	Name   string
	number uint16
}

func (client Client) inactiveClient() {
	client.active = false
	fmt.Printf("Client inactive | value: %t", client.active)
}

func main() {

	client01 := Client{
		name:   "Jorge",
		age:    22,
		active: true,
	}

	client01.city = "Rio"
	client01.Address.country = "Brasil"
	client01.sin.year = 2022

	fmt.Printf("Name: %s | Age: %d | Active: %t\n", client01.name, client01.age, client01.active)

	client01.inactiveClient()

	fmt.Printf("Name: %s | Age: %d | Active: %t", client01.name, client01.age, client01.active)

}
