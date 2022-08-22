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

	router := mux.NewRouter()
	router.HandleFunc("/customers", cust.Get)
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
