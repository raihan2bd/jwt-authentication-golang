package initializers

import (
	"github.com/raihan2bd/jwt-authentication-golang/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
