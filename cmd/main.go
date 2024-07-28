package main

import (
	"fmt"

	auth "calendar_automation/controllers/auth"
	calendar "calendar_automation/controllers/calendar"
	"calendar_automation/middleware"

	"calendar_automation/pkg/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadENV()

	initializers.MakeMigrations()

}

func main() {

	r := gin.Default()

	r.GET("/authenticate", func(c *gin.Context) {
		auth.AuthenticateHandler(c)
	})
	r.GET("/calendar/today", middleware.AuthenticationGuard, calendar.TodaysCalendarHandler)

	fmt.Println("Starting server on :8080...")
	r.Run(":8080")
}
