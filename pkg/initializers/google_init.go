package initializers

import (
	"calendar_automation/models"
	google_calendar "calendar_automation/pkg/google"
	"context"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func GetGoogleConfig() (*oauth2.Config, error) {

	b, err := os.ReadFile("../credentials.json")
	if err != nil {
		return nil, fmt.Errorf("unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		return nil, fmt.Errorf("unable to parse client secret file to config: %v", err)
	}
	return config, nil
}

func SetupGoogle(user models.User) (*calendar.Service, error) {
	if user.AccessToken == "" || user.RefreshToken == "" {
		return nil, fmt.Errorf("please grant access to Google")
	}

	ctx := context.Background()

	config, err := GetGoogleConfig()
	if err != nil {
		return nil, fmt.Errorf("unable to parse client secret file to config: %v", err)
	}

	client, err := google_calendar.GetClientFromDB(user, config)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Calendar client: %v", err)
	}

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Calendar service: %v", err)
	}

	log.Print("🔥 Google service initialized for user")
	return srv, nil
}
