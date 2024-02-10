package handlers

import (
	"github.com/harbisn/go-mkmlrn-restoman/pkg/models"
	"net/http"
)

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	CreateRoom := &models.Room{}
	ParseJSONRequestBody(r, CreateRoom)
	userId := r.Header.Get("x-user-id")
	CreateRoom.CreatedBy = userId
	CreateRoom.UpdatedBy = userId
	room := CreateRoom.CreateRoom()
	SendJSONResponse(w, http.StatusCreated, room)
}

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	params := []string{"status", "capacity", "highestPricePerHour", "lowestPricePerHour"}
	offset, size, order, filters := GetFilterAndPagination(r, params)
	rooms, _ := models.GetAllRooms(offset, size, order, filters)
	pageRooms := Paginate(rooms, len(rooms), size, offset, order, filters)
	SendJSONResponse(w, http.StatusOK, pageRooms)
}

func GetRoomByID(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIDFromRequestToUint64(r, "id")
	ValidateInternalError(w, err)
	roomDetails := models.GetRoomByID(id)
	SendJSONResponse(w, http.StatusOK, roomDetails)
}

func UpdateRoom(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIDFromRequestToUint64(r, "id")
	ValidateInternalError(w, err)
	roomDetails := models.GetRoomByID(id)
	UpdateRoom := &models.Room{}
	ParseJSONRequestBody(r, UpdateRoom)
	if UpdateRoom.Status != "" {
		roomDetails.Status = UpdateRoom.Status
	}
	if UpdateRoom.Type != "" {
		roomDetails.Type = UpdateRoom.Type
	}
	if UpdateRoom.Capacity != roomDetails.Capacity {
		roomDetails.Capacity = UpdateRoom.Capacity
	}
	if UpdateRoom.PricePerHour != roomDetails.PricePerHour {
		roomDetails.PricePerHour = UpdateRoom.PricePerHour
	}
	roomDetails.UpdatedBy = r.Header.Get("x-user-id")
	m := roomDetails.UpdateRoom()
	SendJSONResponse(w, http.StatusOK, m)
}

func DeleteRoom(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIDFromRequestToUint64(r, "id")
	ValidateInternalError(w, err)
	room := models.DeleteRoom(id)
	SendJSONResponse(w, http.StatusOK, room)
}
