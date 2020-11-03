package main

import (
	"ez/config"
	"ez/controllers"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Credential var
type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	db := config.DBInit()
	InDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.POST("/login", loginHandler)
	router.GET("/person/:id", auth, InDB.GetPerson)
	router.GET("/persons", auth, InDB.GetPersons)
	router.POST("/person", auth, InDB.CreatePerson)
	router.PUT("/person", auth, InDB.UpdatePerson)
	router.DELETE("/person/:id", auth, InDB.DeletePerson)
	router.Run(":3000")
}

func loginHandler(c *gin.Context) {
	var user Credential
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Can't Bind Struct",
		})
	}
	if user.Username != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Wrong Username or Password",
		})
	} else {
		if user.Password != "ezadmin2020" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Wrong Password",
			})
		}
	}
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if token != nil && err == nil {
		fmt.Println("Token Verified")
	} else {
		result := gin.H{
			"message": "Not Authorized",
			"error":   err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}
