package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/maha-1030/go-rest-api/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/customers", handlers.GetCustomers)
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	fmt.Println(port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
