package controller

import (
	response_handler "calendar_automation/internal/response"
	internal "calendar_automation/internal/token"
	"calendar_automation/middleware"
	"calendar_automation/models"
	"calendar_automation/pkg/database"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func LoginUserHandler(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	var user models.User
	db := database.DB

	if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		response_handler.Error(c, http.StatusUnauthorized, "User not found")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		response_handler.Error(c, http.StatusUnauthorized, "Incorrect password")

		return
	}

	_, accessToken, err := internal.CreateToken(user.Email, 24*15*time.Hour)

	if err != nil {
		response_handler.Error(c, http.StatusUnauthorized, "Unable to generate token")
		return
	}
	ipData, exists := c.Get("ipData")
	if !exists {
		response_handler.Error(c, http.StatusBadRequest, "Unable to save session")
		return
	}
	session := models.UserSession{
		UserID:     user.ID,
		City:       ipData.(middleware.IPData).City,
		Country:    ipData.(middleware.IPData).Country,
		RegionName: ipData.(middleware.IPData).RegionName,
		Zip:        ipData.(middleware.IPData).Zip,
		Timezone:   ipData.(middleware.IPData).Timezone,
		ISP:        ipData.(middleware.IPData).ISP,
		Query:      ipData.(middleware.IPData).Query,
	}
	if err := db.Create(&session).Error; err != nil {
		response_handler.Error(c, http.StatusBadRequest, "Unable to save session")
		return
	}

	response_handler.Success(c, gin.H{"access_token": accessToken})

}
