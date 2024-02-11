package api

import (
	"github.com/gorilla/mux"
)

var Routes = func(router *mux.Router) {
	// Menu
	//h := &handlers.MenuHandler{}
	//router.HandleFunc("/restoman/menus", h.GetAllMenus.CreateMenu).Methods("POST")
	//router.HandleFunc("/restoman/menus", h.GetAllMenus).Methods("GET")
	//router.HandleFunc("/restoman/menus/{id}", h.GetAllMenus.GetMenuByID).Methods("GET")
	//router.HandleFunc("/restoman/menus/{id}", h.GetAllMenus.UpdateMenu).Methods("PUT")
	//router.HandleFunc("/restoman/menus/{id}", h.GetAllMenus.DeleteMenu).Methods("DELETE")

	// Room
	//router.HandleFunc("/restoman/rooms", handlers.CreateRoom).Methods("POST")
	//router.HandleFunc("/restoman/rooms", handlers.GetAllRooms).Methods("GET")
	//router.HandleFunc("/restoman/rooms/{id}", handlers.GetRoomByID).Methods("GET")
	//router.HandleFunc("/restoman/rooms/{id}", handlers.UpdateRoom).Methods("PUT")
	//router.HandleFunc("/restoman/rooms/{id}", handlers.DeleteRoom).Methods("DELETE")

	//// Reservation
	//router.HandleFunc("/restoman/reservations", handlers.GetAllReservations).Methods("GET")
	//router.HandleFunc("/restoman/reservations/{id}", handlers.GetReservationByID).Methods("GET")
	//router.HandleFunc("/restoman/reservations/{id}", handlers.UpdateReservation).Methods("PUT")

	//// Reservation Room
	//router.HandleFunc("/restoman/reservation-rooms/make", handlers.MakeRoomReservation).Methods("POST")
	//router.HandleFunc("/restoman/reservation-rooms/add", handlers.AddMoreRooms).Methods("POST")
	//router.HandleFunc("/restoman/reservation-rooms", handlers.GetAllReservationRooms).Methods("GET")
}
