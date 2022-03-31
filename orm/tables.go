package orm

import "github.com/jinzhu/gorm"

type Person struct {
	gorm.Model
	ID    int `gorm:"primaryKey"`
	Name  string
	Email string
	City  string
	Phone string
}
