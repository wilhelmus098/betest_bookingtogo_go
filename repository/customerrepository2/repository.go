package customerrepository2

import (
	"github.com/wilhelmus098/betest_bookingtogo_go/helper"
	"gorm.io/gorm"
)

var responseJson = helper.ResponseJson
var responseError = helper.ResponseError

type repository struct {
	db        *gorm.DB
	tableName string
}

// func New(db *gorm.DB) irepository.CustomerRepository {
// 	return &repository{db, "customers"}
// }

// func (r *repository) Get(id int) (*models.Customer, error) {
// 	response := CustomerModel{}

// 	err := r.db.Table(r.tableName).Where("id = ?", id).First(&response).Error
// 	// responseJson(w, http.StatusOK, products)

// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, irepository.ErrVideoNotFound
// 		}
// 		return nil, err
// 	}
// 	return response, nil
// }
