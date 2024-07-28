package middleware

import (
	internal "calendar_automation/internal/token"
	"calendar_automation/models"
	"calendar_automation/pkg/database"
	"calendar_automation/pkg/initializers"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthenticationGuard(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.AbortWithStatusJSON(401, gin.H{
			"error": "No authorization header",
		})
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(401, gin.H{
			"error": "Invalid token format",
		})
		return
	}

	token := headerParts[1]
	isValid, _, err, clm := internal.Validate(token)

	if !isValid {
		c.AbortWithStatusJSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}

	uid := clm["email"]

	var user models.User
	if err := database.DB.Where("Email = ?", uid).First(&user).Error; err != nil {

		c.AbortWithStatusJSON(401, gin.H{
			"error": "User not found",
		})
		return
	}
	service, err := initializers.SetupGoogle(user)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"error": err.Error(),
		})
	}
	c.Set("gs", service)
	c.Set("tk", token)

	c.Next()
}
