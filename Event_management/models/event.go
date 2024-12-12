package models

import (
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES(?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	e.ID = id
	return nil
}

func GetAllEvents(userId int64) ([]Event, error) {

	query := "SELECT * FROM events WHERE user_id = ?"

	rows, err := db.DB.Query(query, userId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []Event{}

	for rows.Next() {
		var event Event
		err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	var event Event

	query := "SELECT * FROM events WHERE id = ?"
	err := db.DB.QueryRow(query, id).Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return &Event{}, err
	}

	return &event, nil
}

func (e Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)

	return err
}

func (e Event) Delete() error {
	query := "DELETE FROM events where id = ?"

	_, err := db.DB.Exec(query, e.ID)

	return err
}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES(? , ?)"

	_, err := db.DB.Exec(query, e.ID, userId)

	return err
}

func (e Event) Cancel(userId int64) error {

	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"

	_, err := db.DB.Exec(query, e.ID, userId)

	return err
}
