package customercontroller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wilhelmus098/betest_bookingtogo_go/helper"
	"github.com/wilhelmus098/betest_bookingtogo_go/models"
	"gorm.io/gorm"
)

var responseJson = helper.ResponseJson
var responseError = helper.ResponseError

func Index(w http.ResponseWriter, r *http.Request) {
	var customers []models.Customer

	if err := models.DB.Preload("Country").Preload("Families").Find(&customers).Error; err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}
	responseJson(w, http.StatusOK, customers)
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var customer models.Customer
	if err := models.DB.Preload("Country").Preload("Families").First(&customer, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			responseError(w, http.StatusNotFound, "Product not found")
			return
		default:
			responseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	responseJson(w, http.StatusOK, customer)
}
