package guclient

import (
	"net/url"
	"strconv"
	"strings"
)

type Page struct {
	Total   int   `json:"total"`
	Page    int   `json:"page"`
	PerPage int   `json:"perPage"`
	Records []any `json:"records"`
}

const BaseUrl string = "https://api.godsunchained.com/v0/"

func Paginate(baseUrl string, page int, perPage int) string {

	if page > 0 || perPage > 0 {
		v := url.Values{}
		v.Set("page", strconv.Itoa(page))
		v.Set("perPage", strconv.Itoa(perPage))

		if strings.ContainsAny(baseUrl, "?") {
			return baseUrl + "?" + v.Encode()
		} else {
			return baseUrl + "&" + v.Encode()
		}
	}
	return baseUrl
}
