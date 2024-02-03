package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/harbisn/go-mkmlrn-restoman/pkg/models"
	"github.com/harbisn/go-mkmlrn-restoman/pkg/util"
	"net/http"
)

func GetAllMenu(w http.ResponseWriter, r *http.Request) {
	menus := models.GetAllMenu()
	res, _ := json.Marshal(menus)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		return
	}
}

func GetMenuByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	menuCode := vars["menuCode"]
	menuDetails, _ := models.GetMenuByCode(menuCode)
	if menuDetails.Code != menuCode {
		errorMessage := fmt.Sprintf("Menu with code %v not found", menuCode)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return
	}
	res, _ := json.Marshal(menuDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		return
	}
}

func CreateMenu(w http.ResponseWriter, r *http.Request) {
	CreateMenu := &models.Menu{}
	util.ParseJSONRequestBody(r, CreateMenu)
	menuDetails, _ := models.GetMenuByCode(CreateMenu.Code)
	if menuDetails.Code == CreateMenu.Code {
		http.Error(w, "Menu Already Exists", http.StatusBadRequest)
		return
	}
	m := CreateMenu.CreateMenu()
	res, _ := json.Marshal(m)
	w.WriteHeader(http.StatusCreated)
	_, err := w.Write(res)
	if err != nil {
		return
	}
}

func UpdateMenu(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	menuCode := vars["menuCode"]
	menuDetails, _ := models.GetMenuByCode(menuCode)
	UpdateMenu := &models.Menu{}
	util.ParseJSONRequestBody(r, UpdateMenu)
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
	m := UpdateMenu.UpdateMenu(menuDetails.ID)
	res, _ := json.Marshal(m)
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		return
	}
}