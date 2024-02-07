package routers

import (
	"github.com/gorilla/mux"
	"github.com/harbisn/go-mkmlrn-restoman/pkg/api/handlers"
)

var RestaurantMenuRoutes = func(router *mux.Router) {
	router.HandleFunc("/restoman/menu", handlers.CreateMenu).Methods("POST")
	router.HandleFunc("/restoman/menu", handlers.GetAllMenu).Methods("GET")
	router.HandleFunc("/restoman/menu/{id}", handlers.GetMenuById).Methods("GET")
	router.HandleFunc("/restoman/menu/{id}", handlers.UpdateMenu).Methods("PATCH")
	router.HandleFunc("/restoman/menu/{id}", handlers.DeleteMenu).Methods("DELETE")
}
