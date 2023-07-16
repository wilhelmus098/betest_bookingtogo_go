package customerrepository2

import (
	"github.com/wilhelmus098/betest_bookingtogo_go/models"
	"gorm.io/datatypes"
)

type CustomerModel struct {
	Id    int64          `gorm:"primaryKey" json:"id"`
	Name  string         `gorm:"type:string" json:"name"`
	Dob   datatypes.Date `json:"dob"`
	Phone string         `gorm:"type:string" json:"phone"`
	Email string         `gorm:"type:string" json:"email"`
}

func (m *CustomerModel) ToCustomer() *models.Customer {
	return &models.Customer{
		Id:    m.Id,
		Name:  m.Name,
		Dob:   m.Dob,
		Phone: m.Phone,
		Email: m.Email,
	}
}
