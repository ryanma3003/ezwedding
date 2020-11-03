package main

import (
	"./config"
	"./controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	InDB := &controllers.InDB{DB: db}

	router := gin.Default

	router.GET("/person/:id", InDB.GetPerson)
	router.GET("/persons", InDB.GetPersons)
	router.POST("/person", InDB.CreatePerson)
	router.PUT("/person", InDB.UpdatePerson)
	router.DELETE("/person/:id", InDB.DeletePerson)
	router.Run(":3000")
}
