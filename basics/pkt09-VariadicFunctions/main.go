package main

import (
	"fmt"
)

func main() {

	total := sum(11, 22, 33, 44, 55, 66, 77, 88, 99, 111, 222, 333, 444, 555)

	fmt.Printf("Total: %d", total)

}

func sum(numbers ...int) int {

	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
