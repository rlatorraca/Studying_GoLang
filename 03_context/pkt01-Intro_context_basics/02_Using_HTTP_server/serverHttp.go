package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", handler)
	log.Println("[INFO] HTTP server start")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(fmt.Printf("[ERROR] HTTP ListenAndServe : %v\n", err))
		return
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("[RLSP - INFO] HTTP server handler")

	ctx := r.Context() // pega o contede do request
	log.Println("[RLSP - INFO] Request started")
	defer log.Println("[RLSP - INFO] Request done")

	select {
	case <-time.After(3 * time.Second):
		// Print on the server side (stdout)
		log.Println("[RLSP - INFO] Request totally processed successfully")

		// Print on the client side (browser)
		_, err := w.Write([]byte("Request processed successfully"))
		if err != nil {
			log.Fatal(fmt.Printf("[RLSP - ERROR] Printing on the Client side : %v\n", err))
			return
		}
	case <-ctx.Done():
		// Print on the server side (stdout)
		log.Println("[RLSP - INFO] Request canceled by Client", ctx.Err())
	}

}
