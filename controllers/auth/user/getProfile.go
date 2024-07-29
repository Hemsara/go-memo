package controller

import (
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

	c.JSON(http.StatusOK, gin.H{"profile": userProfile})
}
