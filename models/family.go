package models

type Family struct {
	Id         int64  `gorm:"primaryKey" json:"id"`
	CustomerID int64  `gorm:"type:int" json:"customer_id"`
	Relation   string `gorm:"type:string" json:"relation"`
	Name       string `gorm:"type:string" json:"name"`
	Dob        string `gorm:"type:string" json:"dob"`
}
