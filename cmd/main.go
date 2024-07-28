package main

import (
	"fmt"

	auth "calendar_automation/controllers/auth"
	calendar "calendar_automation/controllers/calendar"
	"calendar_automation/middleware"

	"calendar_automation/pkg/database"
	"calendar_automation/pkg/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadENV()
	database.New()
	initializers.MakeMigrations()
}

func main() {

	r := gin.Default()

	// Authentication routes
	authRoutes := r.Group("/authenticate")
	{
		authRoutes.GET("/google", func(c *gin.Context) {
			auth.AuthenticateHandler(c)
		})
		authRoutes.POST("/login", func(c *gin.Context) {
			auth.LoginUserHandler(c)
		})
		authRoutes.POST("/register", func(c *gin.Context) {
			auth.RegisterUserHandler(c)
		})
	}

	// Calendar routes
	calendarRoutes := r.Group("/calendar")
	{
		calendarRoutes.GET("/today", middleware.AuthenticationGuard, func(c *gin.Context) {
			calendar.TodaysCalendarHandler(c)
		})
	}

	fmt.Println("Starting server on :8080...")
	r.Run(":8080")
}
