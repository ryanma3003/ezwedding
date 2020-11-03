package controllers

import (
	"net/http"
	
	"../models"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetPerson(c *gin.Context) {
	var (
		person models.Person
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&person).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetPersons(C *gin.Context) {
	var (
		persons []models.Person
		result  gin.H
	)

	idb.DB.Find(&persons)
	if len(persons) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": persons,
			"count":  len(persons),
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreatePerson(c *gin.Context) {
	var (
		person models.Person
		result gin.H
	)
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	person.firstName = first_name
	person.lastName = last_name
	idb.DB.Create(&person)
	result = gin.H{
		"result": person,
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdatePerson(c *gin.Context) {
	id := c.Query("id")
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	var (
		person    models.Person
		newPerson models.Person
		result    gin.H
	)

	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "Data not Found",
		}
	}
	newPerson.firstName = first_name
	newPerson.lastName = last_name

	err = idb.DB.Model(&person).Updates(newPerson).Error
	if err != nil {
		result = gin.H {
			"result": "Update Failed",
		}
	} else {
		result = gin.H {
			"result" = "Successfully Updated data",
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeletePerson(c *gin.Context) {
	var (
		person models.Person
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H {
			"result": "Data not Found",
		}
	}
	err = idb.DB.Delete(&person).Error
	if err != nil {
		result = gin.H {
			"result": "Delete Failed",
		}
	} else {
		result = gin.H {
			"result": "Data Successfully Deleted",
		}
	}

	c.JSON(http.StatusOK, result)
}
