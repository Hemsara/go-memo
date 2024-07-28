package main

import (
	"calendar_automation/pkg/initializers"
	"fmt"
	"log"
	"time"
)

func init() {
	initializers.InitializeGoogle()
}
func main() {

	t := time.Now().Format(time.RFC3339)
	events, err := initializers.GoogleService.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}
	fmt.Println("Upcoming events:")
	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found.")
	} else {
		for _, item := range events.Items {
			date := item.Start.DateTime
			if date == "" {
				date = item.Start.Date
			}
			fmt.Printf("%v (%v)\n", item.Summary, date)
		}
	}
}
