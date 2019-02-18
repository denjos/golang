package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty`
	FirstName string   `json:"firstname,omitempty`
	LastName  string   `json:"lastname,omitempty`
	Address   *Address `json:"address,omitempty`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPeople(W http.ResponseWriter, req *http.Request) {
	json.NewEncoder(W).Encode(people)
}
func GetPerson(W http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(W).Encode(item)
			return
		}
	}
	json.NewEncoder(W).Encode(&Person{})
}
func CreatePerson(W http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(W).Encode(people)
}
func DeletePerson(W http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(W).Encode(people)
}
func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", FirstName: "oscar", LastName: "herrera", Address: &Address{City: "Doubling", State: "California"}})
	people = append(people, Person{ID: "2", FirstName: "wendy", LastName: "astucuri", Address: &Address{City: "Peru", State: "Lima"}})
	//endpoints
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000", router))
}
