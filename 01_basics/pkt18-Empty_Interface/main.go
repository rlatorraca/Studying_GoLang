package main

import (
	"fmt"
)

/* Empty Inteface */
type empty interface{}

func show(e empty) string {
	if _, ok := e.(float64); ok {
		return fmt.Sprintf("Type: %T | Value: %.2f", e, e)
	}
	return fmt.Sprintf("Type: %T | Value: %v", e, e)
}

func main() {

	//var x interface{} = 100.00
	//var y interface{} = 200
	//var z interface{} = "Hello World"

	var x empty = 100.00
	var y empty = 200
	var z empty = "Hello World"

	fmt.Println(show(x))
	fmt.Println(show(y))
	fmt.Println(show(z))

}
