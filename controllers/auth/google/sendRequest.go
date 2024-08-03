package controller

import (
	response_handler "calendar_automation/internal/response"
	google_calendar "calendar_automation/pkg/google"
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

	config, err := google_calendar.GetGoogleConfig()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unable to read client secret file"})
		return

	}
	authURL := google_calendar.GetTokenFromWeb(config, tokenString)

	response_handler.Success(c, gin.H{"google": authURL})

}
