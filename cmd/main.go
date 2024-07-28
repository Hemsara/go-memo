package main

import (
	"fmt"

	auth "calendar_automation/controllers/auth"
	calendar "calendar_automation/controllers/calendar"

	"calendar_automation/pkg/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.InitializeGoogle()
}

func main() {

	r := gin.Default()

	r.GET("/authenticate", func(c *gin.Context) {
		auth.AuthenticateHandler(c)
	})
	r.GET("/calendar/today", func(c *gin.Context) {
		calendar.TodaysCalendarHandler(c)
	})

	fmt.Println("Starting server on :8080...")
	r.Run(":8080")
}
