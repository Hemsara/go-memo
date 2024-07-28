package controller

import (
	internal "calendar_automation/internal/token"
	"calendar_automation/models"
	"calendar_automation/pkg/database"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

func AuthenticateHandler(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")

	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization code not found"})
		return
	}

	if state == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "State parameter not found"})
		return
	}

	fmt.Printf("Authorization code: %s\n", code)
	fmt.Printf("State: %s\n", state)

	isValid, _, err, clm := internal.Validate(state)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Error validating state: " + err.Error(),
		})
		return
	}

	if !isValid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid state",
		})
		return
	}

	uid, ok := clm["email"].(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid email claim in token",
		})
		return
	}

	// TODO : Organize code

	b, err := os.ReadFile("../credentials.json")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to read credentials file",
		})
		return
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to parse credentials file",
		})
		return
	}

	// Exchange the authorization code for an access token
	tok, err := config.Exchange(context.TODO(), code)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to retrieve token from web: " + err.Error(),
		})
		return
	}

	var user models.User

	// Find the user in the database
	if err := database.DB.Where("Email = ?", uid).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "User not found",
		})
		return
	}

	// Update the user's tokens in the database
	user.AccessToken = tok.AccessToken
	user.RefreshToken = tok.RefreshToken

	if err := database.DB.Save(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to update user tokens",
		})
		return
	}

	response := map[string]string{
		"message": "Authentication successful!",
		"code":    code,
		"state":   state,
	}

	c.JSON(http.StatusOK, response)
}
