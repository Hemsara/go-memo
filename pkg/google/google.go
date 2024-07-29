package google_calendar

import (
	"calendar_automation/models"
	"calendar_automation/pkg/database"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func GetTokenFromWeb(config *oauth2.Config, token string) string {
	state := fmt.Sprintf(token)
	authURL := config.AuthCodeURL(state, oauth2.AccessTypeOffline)

	return authURL
}

func GetClientFromDB(user models.User, config *oauth2.Config) (*http.Client, error) {

	token := &oauth2.Token{
		AccessToken:  user.AccessToken,
		RefreshToken: user.RefreshToken,
		Expiry:       user.TokenExpiredAt,
	}

	if time.Now().After(token.Expiry.Add(-5 * time.Minute)) {

		tokenSource := config.TokenSource(context.Background(), token)
		newToken, err := tokenSource.Token()
		if err != nil {
			return nil, err
		}

		user.AccessToken = newToken.AccessToken
		user.RefreshToken = newToken.RefreshToken
		user.TokenExpiredAt = newToken.Expiry

		if err := database.DB.Save(&user).Error; err != nil {
			return nil, err
		}

		token = newToken
	}

	return config.Client(context.Background(), token), nil
}
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
		return nil, nil
	}

	ctx := context.Background()

	config, err := GetGoogleConfig()
	if err != nil {
		return nil, fmt.Errorf("unable to parse client secret file to config: %v", err)
	}

	client, err := GetClientFromDB(user, config)
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
