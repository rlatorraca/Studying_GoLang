package main

import (
	"fmt"
	"github.com/rlatorraca/golang/math"
)

func main() {
	s := math.Sum(10, 15)

	fmt.Printf("Result: %v", s)

	car := math.Car{Make: "Honda", Model: "Accord Touring"}
	fmt.Printf("Result: %v\n", s)
	fmt.Printf("Car: %v | %v\n", car.Make, car.Model)
	fmt.Printf("%v\n", math.A) // atributo criado dentro do arquivo math/math.go
	fmt.Printf("%v", car.Run())

}
