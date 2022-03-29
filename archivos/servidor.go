package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetCategories(siteID string) (Categories, error) {
	response, a := http.Get(siteID)
	if a != nil {
		fmt.Println("Error al conectar con el servidor", a)
		return Categories{}, a
	}

	bytes, b := ioutil.ReadAll(response.Body)
	if b != nil {
		fmt.Println("Error al leer la respuesta del servidor", b)
		return Categories{}, b
	}

	var cats Categories
	err := json.Unmarshal(bytes, &cats) // completar
	if err != nil {
		fmt.Println("Error al deserializar la respuesta del servidor", err)
		return Categories{}, err
	}

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
