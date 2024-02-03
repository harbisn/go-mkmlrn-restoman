package route

import (
	"github.com/gorilla/mux"
	"github.com/harbisn/go-mkmlrn-restoman/src/controller"
)

var RestaurantMenuRoutes = func(router *mux.Router) {
	router.HandleFunc("/restoman/menu/", controller.CreateMenu).Methods("POST")
	router.HandleFunc("/restoman/menu/", controller.GetAllMenu).Methods("GET")
	router.HandleFunc("/restoman/menu/{menuCode}", controller.GetMenuByCode).Methods("GET")
	router.HandleFunc("/restoman/menu/{menuCode}", controller.UpdateMenu).Methods("PATCH")
}
