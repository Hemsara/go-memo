package controller

import (
	google_calendar "calendar_automation/pkg/google"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
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
	b, err := os.ReadFile("../credentials.json")
	if err != nil {

		log.Fatal(err)

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unable to read client secret file"})
		return
	}
	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)

	authURL := google_calendar.GetTokenFromWeb(config, tokenString)

	response := map[string]string{
		"url": authURL,
	}

	c.JSON(http.StatusOK, response)
}
