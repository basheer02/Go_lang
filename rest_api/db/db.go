package db

import (
	"database/sql" //using this

	_ "github.com/mattn/go-sqlite3" //this is a part(doesn't use directly), so it don't remove when it save it(_)
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite3", "api.db") //it creates a file api.db if it not exist

	if err != nil {
		panic(" Could not connect to database")
	}

	DB.SetMaxOpenConns(10) //how many connections can be open simultaneously
	DB.SetMaxIdleConns(5)  //how many connections keep open when no one using the connections
	createTable()
}

func createTable() {

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic(" Could not create users table")
	}

	createEventTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createEventTable)

	if err != nil {
		panic(" Could not create events table")
	}

	createRegistrationTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createRegistrationTable)

	if err != nil {
		panic(" Could not create registration table")
	}

}
