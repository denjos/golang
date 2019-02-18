package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Payload struct {
	Stuff Data
}
type Data struct {
	Fruit   Fruits
	Veggies Vegetables
}
type Fruits map[string]int
type Vegetables map[string]int

func main() {
	http.HandleFunc("/", serveRest)
	http.ListenAndServe("localhost:3000", nil)
}
func serveRest(W http.ResponseWriter, req *http.Request) {
	response, error := getJsonResponse()
	if error != nil {
		panic(error)
	}
	fmt.Fprintf(W, string(response))
}
func getJsonResponse() ([]byte, error) {
	fruits := make(map[string]int)
	fruits["Apples"] = 25
	fruits["Orange"] = 11

	vegetables := make(map[string]int)
	vegetables["Carrots"] = 21
	vegetables["Peppers"] = 0

	d := Data{fruits, vegetables}
	p := Payload{d}
	return json.MarshalIndent(p, "", "  ")
}
