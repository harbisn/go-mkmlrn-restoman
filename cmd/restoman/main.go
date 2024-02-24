package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/harbisn/go-mkmlrn-restoman/internal/database"
	"github.com/harbisn/go-mkmlrn-restoman/internal/menu"
	"github.com/harbisn/go-mkmlrn-restoman/internal/reservation"
	"github.com/harbisn/go-mkmlrn-restoman/internal/room"
	"github.com/rs/zerolog"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	// connect to DB and set dependency
	db := database.Connect()
	newMenuRepository := menu.NewMenuRepository(db)
	newMenuService := menu.NewMenuService(newMenuRepository)
	newMenuHandler := menu.NewMenuHandler(newMenuService)

	newRoomRepository := room.NewRoomRepository(db)
	newRoomService := room.NewRoomService(newRoomRepository)
	newRoomHandler := room.NewRoomHandler(newRoomService)

	newReservationRepository := reservation.NewReservationRepository(db)
	newReservationService := reservation.NewReservationService(newReservationRepository, newRoomRepository, logger)
	newReservationHandler := reservation.NewReservationHandler(newReservationService)

	// set up router
	r := mux.NewRouter()
	http.Handle("/", r)

	r.HandleFunc("/restoman/menus", newMenuHandler.InsertMenuHandler).Methods("POST")
	r.HandleFunc("/restoman/menus", newMenuHandler.SelectMenuHandler).Methods("GET")
	r.HandleFunc("/restoman/menus/{id}", newMenuHandler.UpdateMenuHandler).Methods("PATCH")

	r.HandleFunc("/restoman/rooms", newRoomHandler.InsertRoomHandler).Methods("POST")
	r.HandleFunc("/restoman/rooms", newRoomHandler.SelectRoomHandler).Methods("GET")
	r.HandleFunc("/restoman/rooms/{id}", newRoomHandler.UpdateRoomHandler).Methods("PATCH")

	r.HandleFunc("/restoman/reservations", newReservationHandler.CreateReservationHandler).Methods("POST")

	fmt.Printf("Starting restoman server at port 8080\n")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
