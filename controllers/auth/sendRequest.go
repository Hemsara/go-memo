package controller

import (
	google_calendar "calendar_automation/pkg/google"
	"calendar_automation/pkg/initializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendRequestHandler(c *gin.Context) {
	tk, exists := c.Get("tk")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token not found"})
		return
	}
	tokenString, ok := tk.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token is not a string"})
		return
	}

	config, err := initializers.GetGoogleConfig()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unable to read client secret file"})
		return

	}
	authURL := google_calendar.GetTokenFromWeb(config, tokenString)

	response := map[string]string{
		"url": authURL,
	}

	c.JSON(http.StatusOK, response)
}
