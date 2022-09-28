package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/maha-1030/go-rest-api/http/account"
	"github.com/maha-1030/go-rest-api/http/address"
	"github.com/maha-1030/go-rest-api/http/customer"
	account_service "github.com/maha-1030/go-rest-api/service/account"
	address_service "github.com/maha-1030/go-rest-api/service/address"
	customer_service "github.com/maha-1030/go-rest-api/service/customer"
	account_store "github.com/maha-1030/go-rest-api/store/account"
	address_store "github.com/maha-1030/go-rest-api/store/address"
	customer_store "github.com/maha-1030/go-rest-api/store/customer"
	"github.com/maha-1030/go-rest-api/store/database"
)

const (
	DEFAULT_PORT = "9000"
	DEFAULT_HOST = "localhost"
)

func main() {
	db, err := database.GetDB()
	if err != nil {
		fmt.Println("Failed to connect to the db, err: ", err)

		return
	}

	customerStore := customer_store.NewCustomer(db)
	customerService := customer_service.NewCustomer(customerStore)
	customerHandlers := customer.NewCustomer(customerService)

	accountStore := account_store.NewAccount(db)
	accountService := account_service.NewAccount(accountStore, customerStore)
	accountHandlers := account.NewAccount(accountService)

	addressStore := address_store.NewAddress(db)
	addressService := address_service.NewAddress(addressStore, customerStore)
	addressHandlers := address.NewAddress(addressService)

	router := mux.NewRouter().Headers("Content-Type", "application/json").Subrouter()

	router.HandleFunc("/customer", customerHandlers.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/customer", customerHandlers.Create).Methods(http.MethodPost)
	router.HandleFunc("/customer/{id}", customerHandlers.Update).Methods(http.MethodPut)
	router.HandleFunc("/customer/{id}", customerHandlers.Get).Methods(http.MethodGet)
	router.HandleFunc("/customer/{id}", customerHandlers.Delete).Methods(http.MethodDelete)

	router.HandleFunc("/account", accountHandlers.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customerID}/account", accountHandlers.Create).Methods(http.MethodPost)
	router.HandleFunc("/customer/{customerID}/account/{id}", accountHandlers.UpdateBalance).Methods(http.MethodPut)
	router.HandleFunc("/customer/{customerID}/account/{id}", accountHandlers.Get).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customerID}/account/{id}", accountHandlers.Delete).Methods(http.MethodDelete)

	router.HandleFunc("/address", addressHandlers.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customerID}/address", addressHandlers.Create).Methods(http.MethodPost)
	router.HandleFunc("/customer/{customerID}/address/{id}", addressHandlers.Update).Methods(http.MethodPut)
	router.HandleFunc("/customer/{customerID}/address/{id}", addressHandlers.Get).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customerID}/address/{id}", addressHandlers.Delete).Methods(http.MethodDelete)

	port := os.Getenv("PORT")
	if port == "" {
		port = DEFAULT_PORT
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = DEFAULT_HOST
	}

	err = http.ListenAndServe(host+":"+port, router)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("Server API is running on port: ", port)
}
