package pagination

import (
	"github.com/go-pg/pg/v10/orm"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type PageableDto struct {
	Content         interface{}            `json:"content"`
	TotalElement    int                    `json:"totalElement"`
	NumberOfElement int                    `json:"NumberOfElement"`
	Page            int                    `json:"page"`
	Size            int                    `json:"size"`
	Offset          int                    `json:"offset"`
	Order           string                 `json:"order"`
	Filter          map[string]interface{} `json:"filter"`
}

func GetFilterAndPagination(r *http.Request, params []string) PageableDto {
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

	return PageableDto{Page: page, Offset: offset, Size: size, Order: order, Filter: filters}
}

func SetFilterAndPagination(query *orm.Query, pageable PageableDto) *orm.Query {
	query = query.Limit(pageable.Size).Offset(pageable.Offset)
	if pageable.Order != "" {
		orders := strings.Split(pageable.Order, ",")
		for _, order := range orders {
			o := strings.Split(order, " ")
			query.Order(ConvertParamJsonToDB(o[0]) + " " + o[1])
		}
	}
	pattern := regexp.MustCompile(`^(lowest|highest)`)
	for key, value := range pageable.Filter {
		if strings.Contains(key, "highest") {
			key = pattern.ReplaceAllString(key, "")
			query = query.Where(ConvertParamJsonToDB(key)+" <= ?", value)
		} else if strings.Contains(key, "lowest") {
			key = pattern.ReplaceAllString(key, "")
			query = query.Where(ConvertParamJsonToDB(key)+" >= ?", value)
		} else {
			query = query.Where(ConvertParamJsonToDB(key)+" = ?", value)
		}
	}
	return query
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

func Paginate(content interface{}, pageableDto PageableDto) PageableDto {
	return PageableDto{
		Content:         content,
		TotalElement:    pageableDto.TotalElement,
		NumberOfElement: pageableDto.NumberOfElement,
		Page:            pageableDto.Page,
		Size:            pageableDto.Size,
		Offset:          pageableDto.Offset,
		Order:           pageableDto.Order,
		Filter:          pageableDto.Filter,
	}
}
