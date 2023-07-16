package crud_customer

import "github.com/wilhelmus098/betest_bookingtogo_go/models"

type UseCase interface {
	Get(id int) (*models.Customer, error)
	GetAll() ([]*models.Customer, error)
}
