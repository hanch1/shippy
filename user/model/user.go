package model

import "gorm.io/gorm"

type User struct {
	_ struct{}
	gorm.Model
	Id string
	Name string
	Company string
	Email string
	Password string
}
