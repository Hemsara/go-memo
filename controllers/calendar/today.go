package controller

import (
	response_handler "calendar_automation/internal/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/calendar/v3"
)

type dailyEvent struct {
	Attendees *[]string
	Color     *string
	Summary   *string
	Date      *time.Time
	MeetLink  *string
}

func TodaysCalendarHandler(c *gin.Context) {

	gs, exists := c.Get("gs")
	if !exists {
		response_handler.Error(c, http.StatusInternalServerError, "Google service not found")

		return
	}

	service, ok := gs.(*calendar.Service)
	if !ok {
		response_handler.Error(c, http.StatusInternalServerError, "Invalid Google service type")
		return
	}
	now := time.Now()

	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	tomorrowStart := todayStart.Add(24 * time.Hour)

	events, err := service.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(todayStart.Format(time.RFC3339)).TimeMax(tomorrowStart.Format(time.RFC3339)).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		response_handler.Error(c, http.StatusInternalServerError, "Unable to retrieve calendar events")
		return
	}

	if len(events.Items) == 0 {
		response_handler.Success(c, gin.H{"events": events})

	} else {

		var eventList []dailyEvent
		for _, item := range events.Items {
			date := item.Start.DateTime
			color := item.ColorId

			meetLink := ""
			if item.ConferenceData != nil && item.ConferenceData.EntryPoints != nil {
				for _, entryPoint := range item.ConferenceData.EntryPoints {
					if entryPoint.EntryPointType == "video" {
						meetLink = entryPoint.Uri
					}
				}
			}

			attendees := []string{}
			if item.Attendees != nil {
				for _, attendee := range item.Attendees {
					attendees = append(attendees, attendee.Email)
				}
			}

			if date == "" {
				date = item.Start.Date
			}

			t, err := time.Parse("2006-01-02T15:04:05-07:00", date)
			if err != nil {

				response_handler.Error(c, http.StatusBadRequest, "Something went wrong")

				return
			}
			eventList = append(eventList, dailyEvent{
				Attendees: &attendees,
				Color:     &color,
				Summary:   &item.Summary,
				Date:      &t,
				MeetLink:  &meetLink,
			})
		}
		response_handler.Success(c, gin.H{"events": eventList})

	}
}
