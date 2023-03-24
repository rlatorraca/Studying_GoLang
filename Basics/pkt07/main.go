package main

import "fmt"

// MAPS (em GOLang) ~ HASHMAP Java

func main() {

	// Forma 01
	newSalaries := make(map[string]int)
	newSalaries["Item 01"] = 123
	fmt.Println("=======================================")
	fmt.Printf("%d\n", newSalaries["Item 01"])

	fmt.Println("=======================================")
	// Forma 02
	newestSalaries := map[string]int{}
	newestSalaries["Item 02"] = 321
	fmt.Printf("%d\n", newestSalaries["Item 02"])
	fmt.Println("=======================================")

	// Forma 03
	salaries := map[string]int{
		"John":    2500,
		"Michael": 24566,
		"George":  30000,
		"Marcus":  21500,
		"Jean":    26566,
		"Paul":    33000,
		"Lucas":   5500,
		"Peter":   28566,
		"Jason":   36000,
	}

	fmt.Printf("%d\n", salaries["Marcus"])

	for k, v := range salaries {
		fmt.Printf("Name: %s | Salary: %d\n", k, v)
	}
	fmt.Println("=======================================")
	delete(salaries, "Jean")

	for i, v := range salaries {
		fmt.Printf("Name: %s | Salary: %d\n", i, v)
	}

	total := 0.0
	count := 0
	size := len(salaries)
	fmt.Printf("Size: %d\n", size)
	for _, v := range salaries {
		total = total + float64(v)
		count++
		if count == size {
			fmt.Println("=======================================")
			fmt.Printf("All salaries sum: %.2f\n", total)
		}

	}
}
