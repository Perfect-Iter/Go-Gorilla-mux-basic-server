package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Item struct {
	UID   string `json:"uid"`
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Price int    `json:"price"`
}

var inventory []Item

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint reached for Homepage")
}

func getInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Function Called: getInventory")

	json.NewEncoder(w).Encode(inventory)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item

	_ = json.NewDecoder(r.Body).Decode(&item)

	inventory = append(inventory, item)

	json.NewEncoder(w).Encode(item)

}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.vars(r)

	deleteItemAtUid(params["uid"])

	json.NewEncoder(w).Encode(inventory)
}

func deleteItemAtUid(uid string) {
	for index, item := range inventory {
		if item.UID == uid {
			inventory = append(inventory[:index], inventory[index+1:]...)
			break
		}
	}
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/inventory", getInventory).Methods("GET")
	router.HandleFunc("/inventory", createItem).Methods("POST")
	router.HandleFunc("/inventory/{uid}", deleteItem).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	inventory = append(inventory, Item{
		UID:   "0",
		Name:  "Food",
		Desc:  "Some delicious Food",
		Price: 300,
	})

	inventory = append(inventory, Item{
		UID:   "1",
		Name:  "sword",
		Desc:  "Bronze sword",
		Price: 1200,
	})
	handleRequests()
}
