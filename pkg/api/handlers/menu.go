package handlers

import (
	"github.com/harbisn/go-mkmlrn-restoman/pkg/models"
	"net/http"
)

func CreateMenu(w http.ResponseWriter, r *http.Request) {
	CreateMenu := &models.Menu{}
	ParseJSONRequestBody(r, CreateMenu)
	userId := r.Header.Get("x-user-id")
	CreateMenu.CreatedBy = userId
	CreateMenu.UpdatedBy = userId
	m := CreateMenu.CreateMenu()
	SendJSONResponse(w, http.StatusCreated, m)
}

func GetAllMenu(w http.ResponseWriter, r *http.Request) {
	menus := models.GetAllMenu()
	SendJSONResponse(w, http.StatusOK, menus)
}

func GetMenuById(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIDFromRequestToUint64(r, "id")
	ValidateInternalError(w, err)
	menuDetails, _ := models.GetMenuById(id)
	SendJSONResponse(w, http.StatusOK, menuDetails)
}

func UpdateMenu(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIDFromRequestToUint64(r, "id")
	ValidateInternalError(w, err)
	menuDetails, _ := models.GetMenuById(id)
	UpdateMenu := &models.Menu{}
	ParseJSONRequestBody(r, UpdateMenu)
	if UpdateMenu.Name != "" {
		menuDetails.Name = UpdateMenu.Name
	}
	if UpdateMenu.Category != "" {
		menuDetails.Category = UpdateMenu.Category
	}
	if UpdateMenu.Status != "" {
		menuDetails.Status = UpdateMenu.Status
	}
	if UpdateMenu.Price != menuDetails.Price {
		menuDetails.Price = UpdateMenu.Price
	}
	menuDetails.UpdatedBy = r.Header.Get("x-user-id")
	m := menuDetails.UpdateMenu(menuDetails.ID)
	SendJSONResponse(w, http.StatusOK, m)
}

func DeleteMenu(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIDFromRequestToUint64(r, "id")
	ValidateInternalError(w, err)
	menu, _ := models.DeleteMenu(id)
	SendJSONResponse(w, http.StatusOK, menu)
}
