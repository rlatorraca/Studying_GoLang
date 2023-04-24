package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// Cria um contexto padrão em branco
	ctx := context.Background() // Roda na THread principal

	/**
	WithCancel cria um contexto com um cancelador associado, idenpendente de tempo.
	WithDeadline cria um contexto com um deadline associado, quando um tempo eh atingido.
	WithTimeout cria um contexto com um timeout associado, como se fosse um contagem regressiva.
	WithValue cria um contexto com um valor associado, valor a ser atingido.
	*/
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)

	defer cancel()

	bookingHotel(ctx, "Ottawa")
}

func bookingHotel(ctx context.Context, city string) {

	/**
	Select funciona de forma assincrona (fica esperando até que um case seja atingido)
	*/
	select {
	// Caso o contexto tenha sido cancelado, retorna o erro
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached. Try again later.")
		return
	// caso o context atinja 5 segundos
	case <-time.After(time.Second * 5):
		fmt.Println("Hotel booked for", city)
	}
}
