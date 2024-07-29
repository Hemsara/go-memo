package main

import (
	"fmt"

	auth "calendar_automation/controllers/auth"
	google "calendar_automation/controllers/auth/google"
	user "calendar_automation/controllers/auth/user"

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
	initializers.InitRedis()

}

func main() {

	r := gin.Default()

	// Authentication routes
	authRoutes := r.Group("/authenticate")
	{
		authRoutes.GET("/google", func(c *gin.Context) {
			google.AuthenticateHandler(c)
		})
		// This send request for google to grant access to api
		authRoutes.POST("/google/send-request", middleware.AuthenticationGuard, func(c *gin.Context) {
			google.SendRequestHandler(c)
		})
		authRoutes.POST("/login", func(c *gin.Context) {
			auth.LoginUserHandler(c)
		})
		authRoutes.POST("/register", func(c *gin.Context) {
			auth.RegisterUserHandler(c)
		})
	}

	// User routes
	userRoutes := r.Group("/user")
	{
		userRoutes.GET("/profile", middleware.AuthenticationGuard, func(c *gin.Context) {
			user.GetUserProfile(c)
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
