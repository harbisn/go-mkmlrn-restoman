package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/harbisn/go-mkmlrn-restoman/internal/api/handlers"
	"github.com/harbisn/go-mkmlrn-restoman/internal/database"
	"github.com/harbisn/go-mkmlrn-restoman/internal/repository"
	"log"
	"net/http"
)

func main() {
	db := database.Connect()
	newMenuRepository := repository.NewMenuRepository(db)
	newRoomRepository := repository.NewRoomRepository(db)
	newReservationRepository := repository.NewReservationRepository(db)
	newReservationRoomRepository := repository.NewReservationRoomRepository(db)
	newMenuHandler := handlers.NewMenuHandler(newMenuRepository)
	newRoomHandler := handlers.NewRoomHandler(newRoomRepository)
	newReservationHandler := handlers.NewReservationHandler(newReservationRepository)
	newReservationRoomHandler := handlers.NewReservationRoomHandler(newReservationRepository, newReservationRoomRepository, newRoomRepository)

	r := mux.NewRouter()

	// Menu
	r.HandleFunc("/restoman/menus", newMenuHandler.CreateMenu).Methods("POST")
	r.HandleFunc("/restoman/menus", newMenuHandler.GetAllMenus).Methods("GET")
	r.HandleFunc("/restoman/menus/{id}", newMenuHandler.GetMenuByID).Methods("GET")
	r.HandleFunc("/restoman/menus/{id}", newMenuHandler.UpdateMenu).Methods("PATCH")
	r.HandleFunc("/restoman/menus/{id}", newMenuHandler.DeleteMenu).Methods("DELETE")

	// Room
	r.HandleFunc("/restoman/rooms", newRoomHandler.CreateRoom).Methods("POST")
	r.HandleFunc("/restoman/rooms", newRoomHandler.GetAllRooms).Methods("GET")
	r.HandleFunc("/restoman/rooms/{id}", newRoomHandler.GetRoomByID).Methods("GET")
	r.HandleFunc("/restoman/rooms/{id}", newRoomHandler.UpdateRoom).Methods("PATCH")
	r.HandleFunc("/restoman/rooms/{id}", newRoomHandler.DeleteRoom).Methods("DELETE")

	// Reservation
	r.HandleFunc("/restoman/reservations", newReservationHandler.GetAllReservations).Methods("GET")
	r.HandleFunc("/restoman/reservations/{id}", newReservationHandler.GetReservationByID).Methods("GET")
	r.HandleFunc("/restoman/reservations/{id}", newReservationHandler.UpdateReservation).Methods("PUT")

	// Reservation Room
	r.HandleFunc("/restoman/reservation-rooms/make", newReservationRoomHandler.MakeRoomReservation).Methods("POST")
	r.HandleFunc("/restoman/reservation-rooms/add", newReservationRoomHandler.AddMoreRooms).Methods("POST")
	r.HandleFunc("/restoman/reservation-rooms", newReservationRoomHandler.GetAllReservationRooms).Methods("GET")

	//api.Routes(r)
	http.Handle("/", r)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
