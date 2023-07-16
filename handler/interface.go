package handler

import (
	"net/http"
)

type RestHandler interface {
	GetCustomer(w http.ResponseWriter, r *http.Request)
	GetCustomerById(w http.ResponseWriter, r *http.Request)
}
