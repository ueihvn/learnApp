package main

import (
	"net/http"

	_ "github.com/ueihvn/learnApp/data"
)

func index(w http.ResponseWriter, r *http.Request) {

}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
