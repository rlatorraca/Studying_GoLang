package main

import (
	"errors"
	"fmt"
)

func main() {

	sum, err := sum(51, 4)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("Sum: %d\n", sum)

	fmt.Println(subtraction(54, 4))

}

func sum01(x int, y int) int {
	return x + y
}

func sum(x int, y int) (int, error) {
	if x > 50 {
		return 0, errors.New("Number is higher of 50")
	}
	return x + y, nil
}

func subtraction(x, y int) (int, bool) {
	if x > 50 {
		return x - y, true
	}
	return x - y, false
}
