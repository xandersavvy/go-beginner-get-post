package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Item struct {
	UID   int64   `json:"UID"`
	Name  string  `json:"Name"`
	Price float64 `json:"Price"`
}

var inventory []Item

func main() {
	inventory = append(inventory, Item{
		UID:   0,
		Name:  "Cheese",
		Price: 4.99,
	})

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var item Item
		err := json.NewDecoder(r.Body).Decode(&item) // Obtain item from request JSON
		if err != nil {
			log.Fatal(err)
			return
		}
		inventory = append(inventory, item)
		json.NewEncoder(w).Encode(item)
	}).Methods("POST")

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(inventory)
		fmt.Println(inventory)
	}).Methods("GET")

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
		return
	}
}
