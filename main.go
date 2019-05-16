package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Account struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
}

var accountMap map[string]Account

func init() {
	accountMap = make(map[string]Account)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/accounts", createAccountHandler).Methods("POST")
	r.HandleFunc("/accounts/{id}", getAccountHandler).Methods("GET")
	r.HandleFunc("/accounts/{id}", deleteAccountHandler).Methods("DELETE")

	fmt.Println("server is running at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func createAccountHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("create one account")

	var account Account
	json.NewDecoder(r.Body).Decode(&account)
	id := account.ID
	accountMap[id] = account

	log.Print("sucses create account ", account)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

func getAccountHandler(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]

	log.Print("get one account")

	account, key := accountMap[id]
	w.Header().Add("Content-Type", "application/json")

	if key {
		log.Print("sucsesfully retreived account ", account, " for account id :", id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(account)
	} else {
		log.Print("request account not found!!!")

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w)
	}
}

func deleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]

	log.Print("delete account")

	account, key := accountMap[id]
	w.Header().Delete("Content-Type", "application/json")

	if key {
		log.Print("sucsesfully delete account ", account, " for account id :", id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(account)
	} else {
		log.Print("failed to delete account")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w)
	}
}
