package main

import "fmt"

// GoLang NAO eh OO
// GoLang se assemela as OO languages ao usar Structs

// INTERFACES
// Apenas possivel passar METDOOS e NAO pode ter atributtos

type Person interface {
	Desactive()
}

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
	year   int16
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
	number int32
}

func (client Client) Desactive() {
	client.active = false
	fmt.Printf("Client inactive | value: %t", client.active)
}

func (company Company) Desactive() {

}

func Inactive(person Person) {
	person.Desactive()
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

	Inactive(client01) // client01m is a PERSON too

	fmt.Printf("Name: %s | Age: %d | Active: %t", client01.name, client01.age, client01.active)

	myCompany := Company{
		Name:   "MyCompanny",
		number: 678342376,
	}

	Inactive(myCompany) // myCompanny is a PERSON too [implementa Desactive]

}
