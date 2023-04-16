package main

import (
	"log"
	"net/http"
)

/* Manipulation Arquivos em GoLang*/
func main() {
	fileServer := http.FileServer(http.Dir("./ImportantPackages/pkt09-File_Server/public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/site", func(write http.ResponseWriter, request *http.Request) {
		_, err := write.Write([]byte("Site Running - RLSP"))
		if err != nil {
			write.WriteHeader(http.StatusInternalServerError) /* 500 */
			return
		}
	})
	log.Fatal(http.ListenAndServe(":8181", mux))
}
