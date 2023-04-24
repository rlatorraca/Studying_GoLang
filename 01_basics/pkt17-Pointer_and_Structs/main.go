package main

import "fmt"

// STRUCT
type Client struct {
	name string
	age  uint16
}

func (c Client) walk() {
	c.name = "George Mitchell II"
	fmt.Printf("Client: %v | age: %d is walking ...", c.name, c.age)
}

func (c *Client) walkPointer() {
	c.name = "Jean Pointer II"
	fmt.Printf("Client: %v | age: %d is walking ...", c.name, c.age)
}

func main() {
	client01 := Client{
		name: "George",
		age:  36,
	}

	client01.walk()
	client01.name = "George Mitchell"
	fmt.Printf("\nClient: %v\n", client01.name)
	client01.walk()
	fmt.Printf("\nClient: %v\n", client01.name)

	println("===============================================")
	client01.walkPointer()
	client01.name = "Jean Pointer"
	fmt.Printf("\nClient: %v\n", client01.name)
	client01.walkPointer()
	fmt.Printf("\nClient: %v\n", client01.name)
}
