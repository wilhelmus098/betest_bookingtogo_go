package handler

import (
	"net/http"

	"github.com/wilhelmus098/betest_bookingtogo_go/app/usecase/crud_customer"
	"github.com/wilhelmus098/betest_bookingtogo_go/helper"
)

var responseJson = helper.ResponseJson
var responseError = helper.ResponseError

type restHandler struct {
	customerUsecase crud_customer.UseCase
}

func NewHandler(customerUsecase crud_customer.UseCase) RestHandler {
	return &restHandler{customerUsecase: customerUsecase}
}

func (h *restHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	data, err := h.customerUsecase.GetAll()
	if err == nil {
		responseJson(w, http.StatusOK, data)
		// c.JSON(http.StatusOK, SuccessResponse{Data: data})
	} else {
		responseError(w, http.StatusBadRequest, err.Error())
		// c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
}

func (h *restHandler) GetCustomerById(w http.ResponseWriter, r *http.Request) {
	var id = 1
	data, err := h.customerUsecase.Get(id)
	if err == nil {
		// c.JSON(http.StatusOK, SuccessResponse{Data: data})
		responseJson(w, http.StatusOK, data)
	} else {
		responseError(w, http.StatusBadRequest, err.Error())
	}
}
