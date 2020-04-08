package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)


type Customer struct {
	Id          int    `json:"id"`
	FirstName   string `json:"firstName"`
	MiddleName  string `json:"middleName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
}
type Response struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    []Customer `json:"data"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "********"
	dbname   = "customer_db"
)

var connStr = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func (Customer) TableName() string {
	return "customer"
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	var response Response
	w.Header().Set("Content-Type", "application/json")
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		response.Code = 1
		response.Message = "Database connection failed!"
	}
	decoder := json.NewDecoder(r.Body)
	var customer Customer
	decoder.Decode(&customer)
	db.Create(&customer)
	db.Close()
	response.Code = 0
	response.Message = "Customer created!"
	json.NewEncoder(w).Encode(response)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	var response Response
	w.Header().Set("Content-Type", "application/json")
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		response.Code = 1
		response.Message = "Database connection failed!"
	}
	decoder := json.NewDecoder(r.Body)
	var customer Customer
	decoder.Decode(&customer)
	db.Save(&customer)
	db.Close()
	response.Code = 0
	response.Message = "Customer updated!"
	json.NewEncoder(w).Encode(response)
}

func getAllCustomer(w http.ResponseWriter, r *http.Request) {
	var response Response
	w.Header().Set("Content-Type", "application/json")
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		response.Code = 1
		response.Message = "Database connection failed!"
	}
	var list []Customer
	db = db.Limit(5)
	db.Find(&list)
	db.Close()
	response.Code = 0
	response.Message = "Success!"
	response.Data=list
	json.NewEncoder(w).Encode(response)
}

func getCustomerById(w http.ResponseWriter, r *http.Request) {
	var response Response
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		response.Code = 1
		response.Message = "Database connection failed!"
	}
	db = db.Where("id=?", id)
	var customers []Customer
	db = db.Limit(5)
	db.Find(&customers)
	db.Close()
	response.Code = 0
	response.Message = "Success!"
	response.Data=customers
	json.NewEncoder(w).Encode(response)
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	var response Response
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		response.Code = 1
		response.Message = "Failed to connect to databases"
	}
	db = db.Where("id=?", id)
	var customer Customer
	customer.Id = int(id)
	db.Delete(&customer)
	db.Close()
	response.Code = 0
	response.Message = "Customer deleted!"
	json.NewEncoder(w).Encode(response)
}
