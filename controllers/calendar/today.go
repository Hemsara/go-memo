package controller

import (
	"calendar_automation/pkg/initializers"
	"fmt"
	"net/http"
	"time"
)

func TodaysCalendarHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC3339)
	events, err := initializers.GoogleService.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to retrieve calendar events: %v", err), http.StatusInternalServerError)
		return
	}

	if len(events.Items) == 0 {
		fmt.Fprintln(w, "No upcoming events found.")
	} else {
		for _, item := range events.Items {
			date := item.Start.DateTime
			if date == "" {
				date = item.Start.Date
			}
			fmt.Fprintf(w, "%v (%v)\n", item.Summary, date)
		}
	}
}
