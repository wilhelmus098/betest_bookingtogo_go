package models

import (
	"context"

	"gorm.io/datatypes"
)

type Customer struct {
	Id        int64 `gorm:"primaryKey" json:"id"`
	CountryId int64 `gorm:"type:int" json:"country_id"`
	Country   Country
	Name      string         `gorm:"type:string" json:"name"`
	Dob       datatypes.Date `json:"dob"`
	Phone     string         `gorm:"type:string" json:"phone"`
	Email     string         `gorm:"type:string" json:"email"`
	Families  []Family       `gorm:"foreignkey:customer_id"`
}

type CustomerRepository interface {
	FetchById(c context.Context, id string) ([]Customer, error)
}
