package main

import "fmt"

// STRUCT
type Account struct {
	balance float64
}

func (a *Account) lendSimulator(value float64) float64 {
	a.balance += value
	fmt.Printf("Value added: %.2f | New Balance: %.2f\n", value, a.balance)
	return a.balance
}

// NewAccount Tipo um CONSTRUCTOR in GoLang
func NewAccount() *Account {
	return &Account{balance: 0}
}

func main() {
	account01 := Account{
		balance: 150.00,
	}

	fmt.Printf("Balance: %.2f\n", account01.balance)
	println("===============================================")
	account01.lendSimulator(85)
	println("===============================================")
	account01.lendSimulator(23)

	println("+++++++++++++++++++++++++++++++++++++++++++++++")
	println("===============================================")
	println("+++++++++++++++++++++++++++++++++++++++++++++++")

	account02 := NewAccount() // Using a constructor
	fmt.Printf("Balance: %.2f\n", account02.balance)
	println("===============================================")
	account02.lendSimulator(85)
	println("===============================================")
	account02.lendSimulator(23)
}
