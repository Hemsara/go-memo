package main

import (
	"fmt"
	"net/http"

	auth "calendar_automation/controllers/auth"
	calendar "calendar_automation/controllers/calendar"

	"calendar_automation/pkg/initializers"

	"github.com/go-chi/chi/v5"
)

func init() {
	initializers.InitializeGoogle()
}

func main() {
	r := chi.NewRouter()

	r.Get("/authenticate", auth.AuthenticateHandler)
	r.Get("/calendar/today", calendar.TodaysCalendarHandler)

	fmt.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", r)
}
