package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wilhelmus098/betest_bookingtogo_go/controllers/customercontroller"
	"github.com/wilhelmus098/betest_bookingtogo_go/controllers/productcontroller"
	"github.com/wilhelmus098/betest_bookingtogo_go/models"
)

func main() {
	r := mux.NewRouter()
	models.ConnectDatabase()

	r.HandleFunc("/api/products", productcontroller.Index).Methods("GET")
	r.HandleFunc("/api/products/{id}", productcontroller.Show).Methods("GET")
	r.HandleFunc("/api/products", productcontroller.Create).Methods("POST")
	r.HandleFunc("/api/products/{id}", productcontroller.Update).Methods("PUT")
	r.HandleFunc("/api/products", productcontroller.Delete).Methods("DELETE")

	r.HandleFunc("/api/customer", customercontroller.Index).Methods("GET")
	r.HandleFunc("/api/customer/{id}", customercontroller.Show).Methods("GET")

	http.ListenAndServe(":8080", r)
}
