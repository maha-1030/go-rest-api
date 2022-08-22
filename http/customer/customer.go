package customer

import (
	"encoding/json"
	"fmt"
	"net/http"

	http1 "github.com/maha-1030/go-rest-api/http"
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
