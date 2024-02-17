package menu

import (
	"encoding/json"
	"github.com/harbisn/go-mkmlrn-restoman/internal/helper/handler"
	"github.com/harbisn/go-mkmlrn-restoman/internal/helper/pagination"
	"net/http"
)

type Handler struct {
	menuService *Service
}

func NewMenuHandler(s *Service) *Handler {
	return &Handler{menuService: s}
}

const BadRequestMessage = "Request can't be parsed."

func (h *Handler) InsertMenuHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var menu Menu
	err := decoder.Decode(&menu)
	if err != nil {
		handler.WriteFailResponse(w, http.StatusBadRequest, BadRequestMessage)
		return
	}
	userId := r.Header.Get("x-user-id")
	err = h.menuService.InsertMenu(&menu, userId)
	if err != nil {
		handler.WriteFailResponse(w, http.StatusInternalServerError, "Failed to insert new menu.")
		return
	}
	handler.WriteSuccessResponse(w, http.StatusCreated, menu, nil)
	return
}

func (h *Handler) SelectMenuHandler(w http.ResponseWriter, r *http.Request) {
	params := []string{"status", "category", "highestPrice", "lowestPrice"}
	pageableDto := pagination.GetFilterAndPagination(r, params)
	pageMenu, err := h.menuService.SelectMenu(pageableDto)
	if err != nil {
		handler.WriteFailResponse(w, http.StatusInternalServerError, "Failed to select menu.")
		return
	}
	handler.WriteSuccessResponse(w, http.StatusOK, pageMenu, nil)
}

func (h *Handler) UpdateMenuHandler(w http.ResponseWriter, r *http.Request) {
	var menu Menu
	if err := json.NewDecoder(r.Body).Decode(&menu); err != nil {
		handler.WriteFailResponse(w, http.StatusBadRequest, BadRequestMessage)
		return
	}
	id, err := handler.GetIdFromPath(r)
	if err != nil {
		handler.WriteFailResponse(w, http.StatusBadRequest, BadRequestMessage)
		return
	}
	userId := r.Header.Get("x-user-id")
	updatedMenu, err := h.menuService.UpdateMenu(id, &menu, userId)
	if err != nil {
		handler.WriteFailResponse(w, http.StatusInternalServerError, "Failed to update menu.")
		return
	}
	handler.WriteSuccessResponse(w, http.StatusCreated, updatedMenu, nil)
}
