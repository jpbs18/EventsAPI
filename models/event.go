package models

import (
	"Events-API/db"
	"time"
)

type Event struct {
	ID          int64 		`json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserID 			int64				`json:"userId"`
}

func (e *Event) Save() error{
	query := `INSERT INTO events(name, description, location, dateTime, user_id) VALUES(?, ?, ?, ?, ?)`

	result, err := db.DB.Exec(query, e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id

	return err 
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(
			&event.ID, 
			&event.Name, 
			&event.Description, 
			&event.Location, 
			&event.DateTime, 
			&event.UserID,
		) 

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, rows.Err()
}

func GetEventById(id int64) (*Event, error){
	query := `SELECT * FROM events WHERE id = ?`

	row := db.DB.QueryRow(query, id)
	
	var event Event
	err := row.Scan(	
		&event.ID, 
		&event.Name, 
		&event.Description, 
		&event.Location, 
		&event.DateTime, 
		&event.UserID,
	)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e *Event) Update() error {
	query := `UPDATE events SET name = ?, description = ?, location = ?, dateTime = ? WHERE id = ?`

	_, err := db.DB.Exec(query, e.Name, e.Description, e.Location, e.DateTime, e.ID)
	return err
}

func (e *Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`
	_, err := db.DB.Exec(query, e.UserID)
	return err
}

func (e *Event) Register(userId int64) error {
	query := `INSERT INTO registrations(event_id, user_id) VALUES(?, ?)`
	_, err := db.DB.Exec(query, e.ID, userId)
	return err
}

func (e *Event) CancelRegistration(userId int64) error {
	query := `DELETE FROM registrations WHERE user_id = ? AND event_id = ?`
	_, err := db.DB.Exec(query, userId, e.ID)
	return err
}

func (e *Event) RegistrationExists(userId int64) (bool, error) {
	query := `SELECT COUNT(1) FROM registrations WHERE user_id = ? AND event_id = ?`
	var count int
	err := db.DB.QueryRow(query, userId, e.ID).Scan(&count)
	if err != nil {
		return false, err
	}
	
	return count > 0, nil
}