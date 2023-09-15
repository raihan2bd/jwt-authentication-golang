package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

func Login(c *gin.Context) {
	// Read email/pass from the body
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

	// Get the user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email or password",
		})

		return
	}

	// Compare the password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email or password",
		})

		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create token",
		})

		return
	}

	// set cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	// send it as a response
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged in",
		"token":   tokenString,
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	userData := user.(models.User)
	c.JSON(http.StatusOK, gin.H{
		"message": "I'm logged in",
		"user":    userData.ID,
	})
}
