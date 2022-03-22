package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetCategories(siteID string) (Categories, error) {
	response := http.Get(siteID)
	bytes := ioutil.ReadAll(response.bytes) // completar

	var cats Categories
	err := json.Unmarshal(bytes, &cats) // completar

	return cats, nil
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type Categories []Category

// type json1 struct {

// }
// type json2 struct {

// }

func main() {
	url := "https://api.mercadolibre.com/sites/MLA/categories"
	cats, err := GetCategories(url)
	if err != nil {
		// handle error
	}
	fmt.Println(" -> Las categorias son: \n", cats)

}
