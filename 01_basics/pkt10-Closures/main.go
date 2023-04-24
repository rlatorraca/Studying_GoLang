package main

import (
	"fmt"
)

func main() {

	// CLOSURE => funcao dentro de uma outra funcaoi
	double := func() int {

		sumTmp := sum(11, 22, 33, 44, 55, 66, 77, 88, 99, 111, 222, 333, 444, 555)
		return sumTmp * 2
	}()

	fmt.Printf("Double: %d\n", double)

}

func sum(numbers ...int) int {

	sum := 0
	for _, number := range numbers {
		sum += number
		fmt.Println(number)
	}

	return sum
}
