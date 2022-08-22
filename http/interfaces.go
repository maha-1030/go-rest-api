package http

import "net/http"

type Customer interface {
	Get(w http.ResponseWriter, r *http.Request)
}
