package main

import (
	"fmt"

	"github.com/maha-1030/go-rest-api/database"
)

func main() {
	_, err := database.GetDB()
	if err != nil {
		fmt.Println("Failed to connect to the db")

		return
	}
}
