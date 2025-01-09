package models

import (
	"database/sql"
	"time"

	"github.com/backend/db"
)

// all the logic that deals with storing event data in a database later and that's related with fetching data and so on.

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() (Event, error) {
	query := `
		INSERT INTO events (name, description, location, dateTime, user_id) VALUES (?, ?, ?, ?, ?)
	`

	statement, prepareErr := db.DB.Prepare(query)

	if prepareErr != nil {
		return Event{}, prepareErr
	}

	result, execErr := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	defer statement.Close()

	if execErr != nil {
		return Event{}, execErr
	}

	id, idErr := result.LastInsertId()
	e.ID = id

	return *e, idErr
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	var events []Event = []Event{}

	defer rows.Close()

	for rows.Next() {
		var event Event
		scanERR := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if scanERR != nil {
			return nil, scanERR
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := `
 	SELECT * FROM events WHERE id = ?
 `
	row := db.DB.QueryRow(query, id)

	var event Event

	scanErr := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if scanErr != nil {
		return nil, scanErr
	}

	return &event, nil
}

func (e Event) UpdateEventByID() (sql.Result, error) {
	query := `
		UPDATE events
		SET name = ?,
				description = ?,
				location = ?,
				dateTime = ?
		WHERE id = ?
	`

	stmt, paraErr := db.DB.Prepare(query)

	if paraErr != nil {
		return nil, paraErr
	}

	result, execErr := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)

	defer stmt.Close()

	if execErr != nil {
		return nil, execErr
	}

	return result, nil

}

func (e Event) DeleteEventByID() (sql.Result, error) {
	query := "DELETE FROM events WHERE id = ?"

	stmt, paraErr := db.DB.Prepare(query)

	if paraErr != nil {
		return nil, paraErr
	}

	result, execErr := stmt.Exec(e.ID)

	defer stmt.Close()

	return result, execErr
}

func (e Event) Register(userId int64) error {
	query := `
		INSERT INTO registrations (event_id, user_id) VALUES (?, ?)
	`

	stmt, paraErr := db.DB.Prepare(query)

	if paraErr != nil {
		return paraErr
	}

	defer stmt.Close()

	_, execErr := stmt.Exec(e.ID, userId)

	return execErr

}

func (e Event) Cancel(userId int64) error {
	query := `
		DELETE FROM registrations WHERE event_id = ? AND user_id = ?
	`

	stmt, paraErr := db.DB.Prepare(query)

	if paraErr != nil {
		return paraErr
	}

	defer stmt.Close()

	_, execErr := stmt.Exec(e.ID, userId)

	return execErr

}
