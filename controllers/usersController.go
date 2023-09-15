package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raihan2bd/jwt-authentication-golang/initializers"
	"github.com/raihan2bd/jwt-authentication-golang/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// Get the email/pass from req body
	var body struct {
		Email    string
		Password string
	}

	if (c.Bind(&body)) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read body",
		})

		return
	}

	// Hash the pass
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read body",
		})

		return
	}

	// create the user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create the user",
		})

		return
	}

	// Response
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully.",
	})
}
