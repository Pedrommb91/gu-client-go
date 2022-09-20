package guclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

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

func GetCards(page int, perPage int) (Page, error) {
	perform := BaseUrl + "card"

	url := Paginate(perform, page, perPage)

	response, err := http.Get(url)

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
