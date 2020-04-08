package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/customer/create", createCustomer).Methods("POST")
	router.HandleFunc("/api/v1/customer/getById/{id}", getCustomerById).Methods("GET")
	router.HandleFunc("/api/v1/customer/getAll", getAllCustomer).Methods("GET")
	router.HandleFunc("/api/v1/customer/update", updateCustomer).Methods("PUT")
	router.HandleFunc("/api/v1/customer/delete/{id}", deleteCustomer).Methods("DELETE")
	http.ListenAndServe(":9999", router)
}
