package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	"github.com/jdroguett/grest"
)

//City is my resource
type City struct {
	id   int
	name string
}

//CityController is my controller
type CityController struct {
	Db *sql.DB
}

//Index is the implementation of the action.
func (city *CityController) Index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "CityController index\n")
}

//Show is the implementation of the action.
func (city *CityController) Show(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, fmt.Sprintf("CityController show with id == %v \n", req.Form.Get("id")))
}

//Update is the implementation of the action.
func (city *CityController) Update(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, fmt.Sprintf("CityController update with id == %v \n", req.Form.Get("id")))
}

//Create is the implementation of the action.
func (city *CityController) Create(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "CityController create\n")
}

//Destroy is the implementation of the action.
func (city *CityController) Destroy(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, fmt.Sprintf("CityController destroy with id == %v \n", req.Form.Get("id")))
}

//Country is my resource
type Country struct {
	id   int
	name string
}

//CountryController is my controller
type CountryController struct {
	Db *sql.DB
}

//Index is the implementation of the action.
func (country *CountryController) Index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Country Controller index\n")
}

func main() {
	grest := grest.New()
	//nota: crea la función Resources debe recibir una interface, algo similar a https://golang.org/pkg/net/http/#Handler
	grest.Resources("/cities", &CityController{})
	grest.Resources("/countries", &CountryController{})

	http.ListenAndServe(":4000", nil)
}
