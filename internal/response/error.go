package response_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var statusMessages = map[int]string{
	http.StatusBadRequest:          "Bad Request",
	http.StatusUnauthorized:        "Unauthorized",
	http.StatusForbidden:           "Forbidden",
	http.StatusNotFound:            "Not Found",
	http.StatusInternalServerError: "Internal Server Error",
}

func Error(c *gin.Context, statusCode int, errorMsg string) {
	message, exists := statusMessages[statusCode]
	if !exists {
		message = "An error occurred"
	}

	response := gin.H{
		"error":   message,
		"message": errorMsg,
	}

	c.JSON(statusCode, response)
}
