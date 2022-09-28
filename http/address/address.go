package address

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	http1 "github.com/maha-1030/go-rest-api/http"
	"github.com/maha-1030/go-rest-api/models"
	"github.com/maha-1030/go-rest-api/service"
)

type address struct {
	as service.Address
}

func NewAddress(as service.Address) http1.Address {
	return &address{
		as: as,
	}
}

func (a *address) GetAll(w http.ResponseWriter, r *http.Request) {
	addresses, err := a.as.GetAll()
	if err != nil {
		http1.RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	http1.RespondWithJson(w, http.StatusOK, addresses)
}

func (a *address) Create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	customerID, ok := vars["customerID"]
	if !ok {
		fmt.Println("Missing path param 'customerID' in the address create request")

		http1.RespondWithError(w, http.StatusBadRequest, "missing path param customerID")

		return
	}

	var addressRequest models.Address

	if err := json.NewDecoder(r.Body).Decode(&addressRequest); err != nil {
		fmt.Println("Error while decoding the request body in address create request, err: ", err)

		http1.RespondWithError(w, http.StatusBadRequest, "unable to decode request body into address")

		return
	}

	newAddress, err := a.as.Create(customerID, &addressRequest)
	if err != nil {
		http1.RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	http1.RespondWithJson(w, http.StatusOK, newAddress)
}

func (a *address) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		fmt.Println("Missing path param 'id' in the address get request")

		http1.RespondWithError(w, http.StatusBadRequest, "missing path param id")

		return
	}

	customerID, ok := vars["customerID"]
	if !ok {
		fmt.Println("Missing path param 'customerID' in the address update balance request")

		http1.RespondWithError(w, http.StatusBadRequest, "missing param id")

		return
	}

	var addressRequest models.Address

	if err := json.NewDecoder(r.Body).Decode(&addressRequest); err != nil {
		fmt.Println("Error while decoding the request body in address update balance request, err: ", err)

		http1.RespondWithError(w, http.StatusBadRequest, "unable to decode request body into address")
	}

	updatedAddress, err := a.as.Update(id, customerID, &addressRequest)
	if err != nil {
		http1.RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	http1.RespondWithJson(w, http.StatusOK, updatedAddress)
}

func (a *address) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		fmt.Println("Missing path param 'id' in the address get request")

		http1.RespondWithError(w, http.StatusBadRequest, "missing path param id")

		return
	}

	customerID, ok := vars["customerID"]
	if !ok {
		fmt.Println("Missing path param 'customerID' in the address get request")

		http1.RespondWithError(w, http.StatusBadRequest, "missing path param customerID")

		return
	}

	addr, err := a.as.Get(id, customerID)
	if err != nil {
		http1.RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	http1.RespondWithJson(w, http.StatusOK, addr)
}

func (a *address) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		fmt.Println("Missing path param 'id' in the address delete request")

		http1.RespondWithError(w, http.StatusBadRequest, "missing path param id")

		return
	}

	customerID, ok := vars["customerID"]
	if !ok {
		fmt.Println("Missing path param 'customerID' in the address delete request")

		http1.RespondWithError(w, http.StatusBadRequest, "missing path param customerID")

		return
	}

	if err := a.as.Delete(id, customerID); err != nil {
		http1.RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	http1.RespondWithJson(w, http.StatusOK, map[string]string{"status": "success"})
}
