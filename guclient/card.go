package guclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Total   int   `json:"total"`
	Page    int   `json:"page"`
	PerPage int   `json:"perPage"`
	Records []any `json:"records"`
}

type Card struct {
	User   string `json:"user"`
	Id     Id     `json:"id"`
	Proto  int    `json:"proto"`
	Purity int    `json:"purity"`
}

type Id struct {
	Id    int64 `json:"Int64"`
	Valid bool  `json:"Valid"`
}

func GetCards() (Page, error) {
	response, err := http.Get("https://api.godsunchained.com/v0/card")

	if err != nil {
		fmt.Print(err.Error())
		return Page{}, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return Page{}, err
	}

	var responseObject Page
	json.Unmarshal(responseData, &responseObject)

	return responseObject, nil
}
