package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// Cria um contexto padrão em branco
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)

	defer cancel()

	bookingHotel(ctx, "Ottawa")
}

func bookingHotel(ctx context.Context, city string) {
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
