package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/isoment/bookstore-manager/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	fmt.Printf("Starting server on port 9010\n")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
