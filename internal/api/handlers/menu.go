package handlers

import (
	"github.com/harbisn/go-mkmlrn-restoman/internal/models"
	"github.com/harbisn/go-mkmlrn-restoman/internal/repository"
	"net/http"
)

type MenuHandler struct {
	menuRepository *repository.MenuRepository
}

func NewMenuHandler(r *repository.MenuRepository) *MenuHandler {
	return &MenuHandler{menuRepository: r}
}

func (h *MenuHandler) CreateMenu(w http.ResponseWriter, r *http.Request) {
	menu := &models.Menu{}
	ParseJSONRequestBody(r, menu)
	userId := r.Header.Get("x-user-id")
	menu.CreatedBy = userId
	menu.UpdatedBy = userId
	menu = h.menuRepository.Create(menu)
	SendJSONResponse(w, http.StatusCreated, menu)
}

func (h *MenuHandler) GetAllMenus(w http.ResponseWriter, r *http.Request) {
	params := []string{"status", "category", "highestPrice", "lowestPrice"}
	offset, size, order, filters := GetFilterAndPagination(r, params)
	var p = models.PageableDto{}
	p.Offset = offset
	p.Size = size
	p.Order = order
	p.Filter = filters
	menus, _ := h.menuRepository.GetAll(p)
	pageMenus := Paginate(menus, len(menus), size, offset, order, filters)
	SendJSONResponse(w, http.StatusOK, pageMenus)
}

func (h *MenuHandler) GetMenuByID(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIDFromRequestToUint64(r, "id")
	ValidateInternalError(w, err)
	menu := h.menuRepository.GetById(id)
	SendJSONResponse(w, http.StatusOK, menu)
}

func (h *MenuHandler) UpdateMenu(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIDFromRequestToUint64(r, "id")
	ValidateInternalError(w, err)
	menu := h.menuRepository.GetById(id)
	updateMenu := &models.Menu{}
	ParseJSONRequestBody(r, updateMenu)
	if updateMenu.Name != "" {
		menu.Name = updateMenu.Name
	}
	if updateMenu.Category != "" {
		menu.Category = updateMenu.Category
	}
	if updateMenu.Status != "" {
		menu.Status = updateMenu.Status
	}
	if updateMenu.Price != menu.Price {
		menu.Price = updateMenu.Price
	}
	menu.UpdatedBy = r.Header.Get("x-user-id")
	m := h.menuRepository.Update(updateMenu)
	SendJSONResponse(w, http.StatusOK, m)
}

func (h *MenuHandler) DeleteMenu(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIDFromRequestToUint64(r, "id")
	ValidateInternalError(w, err)
	menu := h.menuRepository.Delete(id)
	SendJSONResponse(w, http.StatusOK, menu)
}
