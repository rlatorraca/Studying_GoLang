package main

import "fmt"

type inteiro int

var (
	var11 inteiro = 100
	var12 float32 = 1.33
)

func main() {
	mySlice := []uint{11, 22, 33, 44, 55, 66, 77, 88, 99}
	fmt.Printf("length: %d | cap: %d | value: %v\n", len(mySlice), cap(mySlice), mySlice)

	fmt.Printf("length: %d | cap: %d | value: %v\n", len(mySlice[:0]), cap(mySlice[:0]), mySlice[:0])

	fmt.Printf("length: %d | cap: %d | value: %v\n", len(mySlice[:4]), cap(mySlice[:4]), mySlice[:4])

	fmt.Printf("length: %d | cap: %d | value: %v\n", len(mySlice[2:]), cap(mySlice[2:]), mySlice[2:])

	mySlice = append(mySlice, 111)
	fmt.Printf("length: %d | cap: %d | value: %v\n", len(mySlice[2:]), cap(mySlice[2:]), mySlice[2:])

	mySlice = append(mySlice, 112)
	fmt.Printf("length: %d | cap: %d | value: %v\n", len(mySlice[2:]), cap(mySlice[2:]), mySlice[2:])
}
