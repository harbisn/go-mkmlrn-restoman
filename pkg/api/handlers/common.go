package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

func ParseIDFromRequestToUint64(r *http.Request, key string) (uint64, error) {
	vars := mux.Vars(r)
	idStr, ok := vars[key]
	if !ok {
		return 0, http.ErrMissingFile
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func ParseJSONRequestBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal(body, x); err != nil {
			return
		}
	}
}

func SendJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	res, err := json.Marshal(data)
	ValidateInternalError(w, err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(res)
	if err != nil {
		return
	}
}

func ValidateInternalError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetFilterAndPagination(r *http.Request, params []string) (int, int, string, map[string]interface{}) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page == 0 {
		page = 1
	}
	size, _ := strconv.Atoi(r.URL.Query().Get("size"))
	if size == 0 {
		size = 10
	}
	offset := (page - 1) * size
	order := r.URL.Query().Get("order")

	filters := make(map[string]interface{})

	for _, param := range params {
		value := r.URL.Query().Get(param)
		if value != "" {
			filters[param] = value
		}
	}

	return offset, size, order, filters
}

type PaginationResponse struct {
	Content      interface{}            `json:"content"`
	TotalElement int                    `json:"totalElement"`
	Page         int                    `json:"page"`
	Size         int                    `json:"size"`
	Order        string                 `json:"order"`
	Filter       map[string]interface{} `json:"filter"`
}

func Paginate(content interface{}, totalElement, size, offset int, order string,
	filter map[string]interface{}) PaginationResponse {
	return PaginationResponse{
		Content:      content,
		TotalElement: totalElement,
		Page:         (offset / size) + 1,
		Size:         size,
		Order:        order,
		Filter:       filter,
	}
}
