package models

type Country struct {
	Id   int64  `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(50)" json:"name"`
	Code string `gorm:"type:varchar(2)" json:"code"`
}
