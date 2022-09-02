package customer

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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

func (c *customer) GetAll(w http.ResponseWriter, r *http.Request) {
	customers, err := c.cs.GetAll()
	if err != nil {
		http1.RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	http1.RespondWithJson(w, http.StatusOK, customers)
}

func (c *customer) Create(w http.ResponseWriter, r *http.Request) {
	var customerRequest models.Customer

	if err := json.NewDecoder(r.Body).Decode(&customerRequest); err != nil {
		fmt.Println("Error while decoding the request body in customer create request, err: ", err)

		http1.RespondWithError(w, http.StatusBadRequest, "unable to decode request body into customer")

		return
	}

	newCustomer, err := c.cs.Create(&customerRequest)
	if err != nil {
		http1.RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	http1.RespondWithJson(w, http.StatusOK, newCustomer)
}

func (c *customer) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		fmt.Println("Missing path param 'id' in the customer update request")

		http1.RespondWithError(w, http.StatusBadRequest, "missing param id")

		return
	}

	var customerRequest models.Customer

	if err := json.NewDecoder(r.Body).Decode(&customerRequest); err != nil {
		fmt.Println("Error while decoding the request body in customer update request, err: ", err)

		http1.RespondWithError(w, http.StatusBadRequest, "unable to decode request body into customer")
	}

	updatedCustomer, err := c.cs.Update(id, &customerRequest)
	if err != nil {
		http1.RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	http1.RespondWithJson(w, http.StatusOK, updatedCustomer)
}

func (c *customer) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		fmt.Println("Missing path param 'id' in the customer get request")

		http1.RespondWithError(w, http.StatusBadRequest, "missing param id")

		return
	}

	cust, err := c.cs.Get(id)
	if err != nil {
		http1.RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	http1.RespondWithJson(w, http.StatusOK, cust)
}

func (c *customer) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		fmt.Println("Missing path param 'id' in the customer delete request")

		http1.RespondWithError(w, http.StatusBadRequest, "missing param id")

		return
	}

	if err := c.cs.Delete(id); err != nil {
		http1.RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	http1.RespondWithJson(w, http.StatusOK, map[string]string{"status": "success"})
}
