package store

import "github.com/maha-1030/go-rest-api/models"

type Customer interface {
	Get() (customers []models.Customer, err error)
}
