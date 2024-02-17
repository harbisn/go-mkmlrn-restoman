package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/harbisn/go-mkmlrn-restoman/internal/database"
	"github.com/harbisn/go-mkmlrn-restoman/internal/menu"
	"github.com/harbisn/go-mkmlrn-restoman/internal/room"
	"log"
	"net/http"
)

func main() {
	// connect to DB and pass db connection
	db := database.Connect()
	newMenuRepository := menu.NewMenuRepository(db)
	newMenuService := menu.NewMenuService(newMenuRepository)
	newMenuHandler := menu.NewMenuHandler(newMenuService)

	room.NewRoomRepository(db)
	newRoomRepository := room.NewRoomRepository(db)
	newRoomService := room.NewRoomService(newRoomRepository)
	newRoomHandler := room.NewRoomHandler(newRoomService)

	// set up router
	r := mux.NewRouter()
	http.Handle("/", r)

	r.HandleFunc("/restoman/menus", newMenuHandler.InsertMenuHandler).Methods("POST")
	r.HandleFunc("/restoman/menus", newMenuHandler.SelectMenuHandler).Methods("GET")
	r.HandleFunc("/restoman/menus/{id}", newMenuHandler.UpdateMenuHandler).Methods("PATCH")

	r.HandleFunc("/restoman/rooms", newRoomHandler.InsertRoomHandler).Methods("POST")
	r.HandleFunc("/restoman/rooms", newRoomHandler.SelectRoomHandler).Methods("GET")
	r.HandleFunc("/restoman/rooms/{id}", newRoomHandler.UpdateRoomHandler).Methods("PATCH")

	fmt.Printf("Starting restoman server at port 8080\n")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
