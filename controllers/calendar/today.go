package controller

import (
	"calendar_automation/pkg/initializers"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// TodaysCalendarHandler handles requests for today's calendar events
func TodaysCalendarHandler(c *gin.Context) {

	t := time.Now().Format(time.RFC3339)

	events, err := initializers.GoogleService.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Unable to retrieve calendar events: %v", err),
		})
		return
	}

	if len(events.Items) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No upcoming events found.",
		})
	} else {

		var eventList []gin.H
		for _, item := range events.Items {
			date := item.Start.DateTime
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

			eventList = append(eventList, gin.H{
				"summary":   item.Summary,
				"date":      date,
				"meetLink":  meetLink,
				"attendees": attendees,
			})
		}

		c.JSON(http.StatusOK, eventList)
	}
}
