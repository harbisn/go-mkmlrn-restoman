package routers

import (
	"github.com/gorilla/mux"
	"github.com/harbisn/go-mkmlrn-restoman/pkg/api/handlers"
)

var RestaurantMenuRoutes = func(router *mux.Router) {
	router.HandleFunc("/restoman/menu/", handlers.CreateMenu).Methods("POST")
	router.HandleFunc("/restoman/menu/", handlers.GetAllMenu).Methods("GET")
	router.HandleFunc("/restoman/menu/{menuCode}", handlers.GetMenuByCode).Methods("GET")
	router.HandleFunc("/restoman/menu/{menuCode}", handlers.UpdateMenu).Methods("PATCH")
}
