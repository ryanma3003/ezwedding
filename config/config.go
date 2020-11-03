package config

import (
	"../models"
	"github.com/jinzhu/gorm"
)

// DBInit Connection to DB
func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/godb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to Connect to Database")
	}

	db.AutoMigrate(models.Person{})
	return db
}
