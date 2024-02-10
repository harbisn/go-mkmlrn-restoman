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
	ParseJSONRequestBody(r, CreateMenu)
	params := []string{"status", "category", "highestPrice", "lowestPrice"}
	offset, size, order, filters := GetFilterAndPagination(r, params)
	menus, _ := models.GetAllMenu(offset, size, order, filters)
	pageMenus := Paginate(menus, len(menus), size, offset, order, filters)
	SendJSONResponse(w, http.StatusOK, pageMenus)
}

func GetMenuById(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIDFromRequestToUint64(r, "id")
	ValidateInternalError(w, err)
	menuDetails := models.GetMenuById(id)
	SendJSONResponse(w, http.StatusOK, menuDetails)
}

func UpdateMenu(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIDFromRequestToUint64(r, "id")
	ValidateInternalError(w, err)
	menuDetails := models.GetMenuById(id)
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
	m := menuDetails.UpdateMenu()
	SendJSONResponse(w, http.StatusOK, m)
}

func DeleteMenu(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIDFromRequestToUint64(r, "id")
	ValidateInternalError(w, err)
	menu := models.DeleteMenu(id)
	SendJSONResponse(w, http.StatusOK, menu)
}
