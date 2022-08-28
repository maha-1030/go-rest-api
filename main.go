package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/maha-1030/go-rest-api/http/customer"
	customer_service "github.com/maha-1030/go-rest-api/service/customer"
	customer_store "github.com/maha-1030/go-rest-api/store/customer"
	"github.com/maha-1030/go-rest-api/store/database"
)

func main() {
	db, err := database.GetDB()
	if err != nil {
		fmt.Println("Failed to connect to the db, err: ", err)

		return
	}

	customerStore := customer_store.NewCustomer(db)
	customerService := customer_service.NewCustomer(customerStore)
	cust := customer.NewCustomer(customerService)

	router := mux.NewRouter().Headers("Content-Type", "application/json").Subrouter()
	router.HandleFunc("/customers", cust.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/customer", cust.Create).Methods(http.MethodPost)
	router.HandleFunc("/customer/{id}", cust.Update).Methods(http.MethodPut)
	router.HandleFunc("/customer/{id}", cust.Get).Methods(http.MethodGet)
	router.HandleFunc("/customer/{id}", cust.Delete).Methods(http.MethodDelete)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	fmt.Println(port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
