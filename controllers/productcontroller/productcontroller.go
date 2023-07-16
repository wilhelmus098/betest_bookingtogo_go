package productcontroller

import (
	"encoding/json"
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
	var products []models.Product

	if err := models.DB.Find(&products).Error; err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}
	responseJson(w, http.StatusOK, products)
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var product models.Product
	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			responseError(w, http.StatusNotFound, "Product not found")
			return
		default:
			responseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	responseJson(w, http.StatusOK, product)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()
	if err := models.DB.Create(&product).Error; err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseJson(w, http.StatusOK, product)

}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var product models.Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()

	if models.DB.Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		responseError(w, http.StatusInternalServerError, "Id not found, failed to update")
		return
	}

	product.Id = id
	responseJson(w, http.StatusOK, product)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	input := map[string]string{"id": ""}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	var product models.Product
	if models.DB.Delete(&product, input["id"]).RowsAffected == 0 {
		responseError(w, http.StatusBadRequest, "Id not found, failed to delete")
		return
	}

	response := map[string]string{"message": "Product deleted"}
	responseJson(w, http.StatusOK, response)
}
