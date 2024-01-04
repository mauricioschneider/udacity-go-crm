package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type Customer struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Role      string `json:"role,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Contacted bool   `json:"contacted,omitempty"`
}

// Customer database
var customerList = map[string]Customer{
	"e1827a7d-1acd-46ef-9d92-2a4d78bd7669": {
		ID:        "e1827a7d-1acd-46ef-9d92-2a4d78bd7669",
		Name:      "Clementina DuBuque",
		Role:      "CEO",
		Email:     "Rey.Padberg@karina.biz",
		Phone:     "(024)648-3804",
		Contacted: true,
	},
	"fb871ddf-ad69-40b9-966d-ab8e29504438": {
		ID:        "fb871ddf-ad69-40b9-966d-ab8e29504438",
		Name:      "Glenna Reichert",
		Role:      "Software Engineer",
		Email:     "Chaim_McDermott@dana.io",
		Phone:     "(775)976-6794",
		Contacted: false,
	},
	"17b0f3c0-7148-4e7a-8b91-71c22ca1105c": {
		ID:        "",
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

	var c Customer

	// This is more elegant, but project rubric requires use of io
	//err := json.NewDecoder(r.Body).Decode(&c)

	reqBody, _ := io.ReadAll(r.Body)
	json.Unmarshal(reqBody, &c)

	// No validation implementation specified in Rubric i.e. required fields
	// However, we can assume a contact should at least have a name

	if c.Name == "" {
		errorResponse := ErrorResponse{Error: "Invalid request data: Name is required"}
		jsonResonse, _ := json.Marshal(errorResponse)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResonse)
		return
	}

	c.ID = uuid.New().String()
	customerList[c.ID] = c

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userId := mux.Vars(r)["id"]

	if found, ok := customerList[userId]; ok {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(found)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userId := mux.Vars(r)["id"]

	if _, ok := customerList[userId]; ok {
		var c Customer
		reqBody, _ := io.ReadAll(r.Body)
		json.Unmarshal(reqBody, &c)
		if c.Name == "" {
			errorResponse := ErrorResponse{Error: "Invalid request data: Name is required"}
			jsonResonse, _ := json.Marshal(errorResponse)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(jsonResonse)
			return
		}
		c.ID = userId
		customerList[c.ID] = c
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(c)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userId := mux.Vars(r)["id"]

	if _, ok := customerList[userId]; ok {
		delete(customerList, userId)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("{}")
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
