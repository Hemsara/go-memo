package controller

import (
	response_handler "calendar_automation/internal/response"
	"calendar_automation/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserProfile struct {
	Email              string
	IsLinkedWithGoogle bool
	FullName           string
}

func GetUserProfile(c *gin.Context) {
	usr, _ := c.Get("usr")

	user, ok := usr.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to parse user"})
		return
	}

	isGoogleGranted := user.AccessToken != "" && user.RefreshToken != ""

	userProfile := UserProfile{
		FullName:           user.FullName,
		Email:              user.Email,
		IsLinkedWithGoogle: isGoogleGranted,
	}
	response_handler.Success(c, gin.H{"profile": userProfile})

}
