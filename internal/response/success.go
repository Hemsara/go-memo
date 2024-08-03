package response_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data gin.H) {
	response := gin.H{
		"status":  "success",
		"message": "Request successful",
		"data":    data,
	}

	c.JSON(http.StatusOK, response)
}

func Created(c *gin.Context, data gin.H) {
	response := gin.H{
		"status":  "success",
		"message": "Resource created successfully",
		"data":    data,
	}

	c.JSON(http.StatusCreated, response)
}
