package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthenticateHandler(c *gin.Context) {
	response := map[string]string{
		"message": "Authentication successful!",
	}

	c.JSON(http.StatusOK, response)

}
