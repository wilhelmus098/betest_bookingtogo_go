package crud_customer

import (
	"errors"
	"fmt"

	"github.com/wilhelmus098/betest_bookingtogo_go/app/repository"
	"github.com/wilhelmus098/betest_bookingtogo_go/models"
)

type usecase struct {
	customerRepo repository.CustomerRepository
}

func NewUseCase(customerRepo repository.CustomerRepository) UseCase {
	return &usecase{customerRepo: customerRepo}
}

func (uc *usecase) Get(id int) (*models.Customer, error) {
	data, err := uc.customerRepo.Get(id)
	if err != nil {
		if errors.Is(err, repository.ErrCustomerNotFound) {
			return nil, ErrCustomerNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return data, nil
}

func (uc *usecase) GetAll() ([]*models.Customer, error) {
	data, err := uc.customerRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return data, nil
}
