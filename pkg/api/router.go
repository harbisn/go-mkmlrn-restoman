package api

import (
	"github.com/gorilla/mux"
	"github.com/harbisn/go-mkmlrn-restoman/pkg/api/handlers"
)

var Routes = func(router *mux.Router) {
	// Menu
	router.HandleFunc("/restoman/menus", handlers.CreateMenu).Methods("POST")
	router.HandleFunc("/restoman/menus", handlers.GetAllMenus).Methods("GET")
	router.HandleFunc("/restoman/menus/{id}", handlers.GetMenuByID).Methods("GET")
	router.HandleFunc("/restoman/menus/{id}", handlers.UpdateMenu).Methods("PUT")
	router.HandleFunc("/restoman/menus/{id}", handlers.DeleteMenu).Methods("DELETE")

	// Room
	router.HandleFunc("/restoman/rooms", handlers.CreateRoom).Methods("POST")
	router.HandleFunc("/restoman/rooms", handlers.GetAllRooms).Methods("GET")
	router.HandleFunc("/restoman/rooms/{id}", handlers.GetRoomByID).Methods("GET")
	router.HandleFunc("/restoman/rooms/{id}", handlers.UpdateRoom).Methods("PUT")
	router.HandleFunc("/restoman/rooms/{id}", handlers.DeleteRoom).Methods("DELETE")

	// Reservation
	router.HandleFunc("/restoman/reservations", handlers.CreateReservation).Methods("POST")
	router.HandleFunc("/restoman/reservations", handlers.GetAllReservations).Methods("GET")
	router.HandleFunc("/restoman/reservations/{id}", handlers.GetReservationByID).Methods("GET")
	router.HandleFunc("/restoman/reservations/{id}", handlers.UpdateReservation).Methods("PUT")
}
