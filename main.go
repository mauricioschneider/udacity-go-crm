package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Customer struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Role      string `json:"role,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Contacted bool   `json:"contacted,omitempty"`
}

// Customer database
var customerList = []Customer{
	{
		ID:        "e1827a7d-1acd-46ef-9d92-2a4d78bd7669",
		Name:      "Clementina DuBuque",
		Role:      "CEO",
		Email:     "Rey.Padberg@karina.biz",
		Phone:     "(024)648-3804",
		Contacted: true,
	},
	{
		ID:        "e1827a7d-1acd-46ef-9d92-2a4d78bd7669",
		Name:      "Glenna Reichert",
		Role:      "Software Engineer",
		Email:     "Chaim_McDermott@dana.io",
		Phone:     "(775)976-6794",
		Contacted: false,
	},
	{
		ID:        uuid.New().String(),
		Name:      "Nicholas Runolfsdottir V",
		Role:      "Security Analyst",
		Email:     "Sherwood@rosamond.me",
		Phone:     "(586)493-6943",
		Contacted: true,
	},
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(customerList)
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userId := mux.Vars(r)["id"]

	var customer Customer

	for _, c := range customerList {
		if c.ID == userId {
			customer = c
			break
		}
	}

	if customer.ID == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(customer)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")
	fmt.Println("Listening on port 8000")
	http.ListenAndServe(":8000", router)
}
