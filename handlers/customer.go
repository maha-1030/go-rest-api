package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maha-1030/go-rest-api/database"
	"github.com/maha-1030/go-rest-api/models"
)

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	db, err := database.GetDB()
	if err != nil {
		fmt.Println("Failed to connect to the db, err: ", err)
		w.Write([]byte(err.Error()))

		return
	}

	var customers []models.Customer

	res := db.Find(&customers)
	if res.Error != nil {
		fmt.Println("Error while retrieving customers data, err: ", err)
		w.Write([]byte(err.Error()))

		return
	}

	bytes, err := json.Marshal(customers)
	if err != nil {
		fmt.Println("Error occured while marshaling, err: ", err)
		w.Write([]byte(err.Error()))

		return
	}

	w.Write(bytes)
}
