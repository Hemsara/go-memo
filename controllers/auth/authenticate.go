package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

	response := map[string]string{
		"message": "Authentication successful!",
		"code":    code,
		"state":   state,
	}

	c.JSON(http.StatusOK, response)
}
