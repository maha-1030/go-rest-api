package customer

import (
	"encoding/json"
	"fmt"
	"net/http"

	http1 "github.com/maha-1030/go-rest-api/http"
	"github.com/maha-1030/go-rest-api/models"
	"github.com/maha-1030/go-rest-api/service"
)

type customer struct {
	cs service.Customer
}

func NewCustomer(cs service.Customer) http1.Customer {
	return &customer{
		cs: cs,
	}
}

func (c *customer) Get(w http.ResponseWriter, r *http.Request) {
	customers, err := c.cs.Get()
	if err != nil {
		http1.RespondWithJson(w, http.StatusInternalServerError, err)

		return
	}

	http1.RespondWithJson(w, http.StatusOK, customers)
}

func (c *customer) Create(w http.ResponseWriter, r *http.Request) {
	var customerRequest models.Customer

	if err := json.NewDecoder(r.Body).Decode(&customerRequest); err != nil {
		fmt.Println("Error while decoding the request body into customer, err: ", err)

		http1.RespondWithJson(w, http.StatusBadRequest,
			map[string]string{"err": "unable to decode request body into customer"})

		return
	}

	newCustomer, err := c.cs.Create(&customerRequest)
	if err != nil {
		http1.RespondWithJson(w, http.StatusInternalServerError, err)

		return
	}

	http1.RespondWithJson(w, http.StatusOK, newCustomer)
}
