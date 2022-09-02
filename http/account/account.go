package account

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	http1 "github.com/maha-1030/go-rest-api/http"
	"github.com/maha-1030/go-rest-api/models"
	"github.com/maha-1030/go-rest-api/service"
)

type account struct {
	as service.Account
}

func NewAccount(as service.Account) http1.Account {
	return &account{
		as: as,
	}
}

func (a *account) GetAll(w http.ResponseWriter, r *http.Request) {
	accounts, err := a.as.GetAll()
	if err != nil {
		http1.RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	http1.RespondWithJson(w, http.StatusOK, accounts)
}

func (a *account) Create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	customerID, ok := vars["customerID"]
	if !ok {
		fmt.Println("Missing path param 'customerID' in the account create request")

		http1.RespondWithError(w, http.StatusBadRequest, "missing path param customerID")

		return
	}

	var accountRequest models.Account

	if err := json.NewDecoder(r.Body).Decode(&accountRequest); err != nil {
		fmt.Println("Error while decoding the request body in account create request, err: ", err)

		http1.RespondWithError(w, http.StatusBadRequest, "unable to decode request body into account")

		return
	}

	newAccount, err := a.as.Create(customerID, &accountRequest)
	if err != nil {
		http1.RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	http1.RespondWithJson(w, http.StatusOK, newAccount)
}

func (a *account) UpdateBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["customerID"]
	if !ok {
		fmt.Println("Missing path param 'customerID' in the account update balance request")

		http1.RespondWithError(w, http.StatusBadRequest, "missing param id")

		return
	}

	var accountRequest models.Account

	if err := json.NewDecoder(r.Body).Decode(&accountRequest); err != nil {
		fmt.Println("Error while decoding the request body in account update balance request, err: ", err)

		http1.RespondWithError(w, http.StatusBadRequest, "unable to decode request body into account")
	}

	updatedAccount, err := a.as.UpdateBalance(id, &accountRequest)
	if err != nil {
		http1.RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	http1.RespondWithJson(w, http.StatusOK, updatedAccount)
}

func (a *account) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		fmt.Println("Missing path param 'id' in the account get request")

		http1.RespondWithError(w, http.StatusBadRequest, "missing path param id")

		return
	}

	customerID, ok := vars["customerID"]
	if !ok {
		fmt.Println("Missing path param 'customerID' in the account get request")

		http1.RespondWithError(w, http.StatusBadRequest, "missing path param customerID")

		return
	}

	acc, err := a.as.Get(id, customerID)
	if err != nil {
		http1.RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	http1.RespondWithJson(w, http.StatusOK, acc)
}

func (a *account) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		fmt.Println("Missing path param 'id' in the account delete request")

		http1.RespondWithError(w, http.StatusBadRequest, "missing path param id")

		return
	}

	customerID, ok := vars["customerID"]
	if !ok {
		fmt.Println("Missing path param 'customerID' in the account delete request")

		http1.RespondWithError(w, http.StatusBadRequest, "missing path param customerID")

		return
	}

	if err := a.as.Delete(id, customerID); err != nil {
		http1.RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	http1.RespondWithJson(w, http.StatusOK, map[string]string{"status": "success"})
}
