package models

import "time"

// all the logic that deals with storing event data in a database later and that's related with fetching data and so on.

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events []Event = []Event{}

func (e Event) Save() {
	events = append(events, e)
	// later store into database
}

func GetAllEvents() []Event {
	return events
}
