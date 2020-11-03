package models

import "github.com/jinzhu/gorm"

type Person struct {
	gorm.Model
	firstName string
	lastName  string
}
