package repository

import (
	"errors"

	"github.com/wilhelmus098/betest_bookingtogo_go/models"
)

var ErrCustomerNotFound = errors.New("Customer not found")

type CustomerRepository interface {
	Get(id int) (*models.Customer, error)
	GetAll() ([]*models.Customer, error)
}
