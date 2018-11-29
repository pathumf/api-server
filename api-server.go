package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pathumf/api-server/api"
)

//TODO import package
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo/", echo)
	http.HandleFunc("api/books", api.BooksHandlerFunc)
	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PATH")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

}

func echo(w http.ResponseWriter, r *http.Request) {
	m := r.URL.Query()["message"][0]
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, m)
}
