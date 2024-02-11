package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/harbisn/go-mkmlrn-restoman/pkg/models"
	"io"
	"net/http"
	"strconv"
	"unicode"
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

func ParseJSONRequestBody(r *http.Request, data interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal(body, data); err != nil {
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
			filters[ConvertParamJsonToDB(param)] = value
		}
	}

	return offset, size, order, filters
}

func ConvertParamJsonToDB(s string) string {
	var res []rune
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				res = append(res, '_')
			}
			res = append(res, unicode.ToLower(r))
		} else {
			res = append(res, r)
		}
	}
	return string(res)
}

func Paginate(content interface{}, totalElement, size, offset int, order string,
	filter map[string]interface{}) models.PaginationResponse {
	return models.PaginationResponse{
		Content:      content,
		TotalElement: totalElement,
		Page:         (offset / size) + 1,
		Size:         size,
		Order:        order,
		Filter:       filter,
	}
}
