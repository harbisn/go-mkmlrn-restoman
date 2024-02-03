package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/harbisn/go-mkmlrn-restoman/src/route"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	route.RestaurantMenuRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
