package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/calendar/v3"
)

func TodaysCalendarHandler(c *gin.Context) {

	gs, exists := c.Get("gs")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Google service not found"})
		return
	}

	// Assert the type of the service to the expected type
	service, ok := gs.(*calendar.Service)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Google service type"})
		return
	}
	t := time.Now().Format(time.RFC3339)

	events, err := service.Events.List("primary").ShowDeleted(false).
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
