package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL)
	urlParams := mux.Vars(r)
	fmt.Println(urlParams)
	fmt.Println("quary params", r.URL.Query())

	switch r.Method {
	case http.MethodGet:
		fmt.Println("GET")
	case http.MethodPost:
		fmt.Println("POST")
	}

}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Middleware", r.URL)
			next.ServeHTTP(w, r)

		})
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/users/{id:[0-9]+}", Handler).Methods("GET", "POST")
	mux.Use(Middleware)

	http.ListenAndServe(":7777", mux)
}
