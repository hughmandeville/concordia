//
// Command to use to manage the Concordia SQLite database.
// Actions
//   create - create the database.
//   export - export the database to a JSON file.
//   list   - list contents of the database
//
// Tables
//   boat   - current boat information.
//   image  - boat images.
//   link   - boat links.
//   owner  - boat owner history.

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var (
	cmd       = "list" // create, delete, export, or list
	boatsFile = "../../json/boats.json"
	dbFile    = "../../db/concordia.db"
)

// store current boat information
type Boat struct {
	ID          uint32    `json:"id"`
	BoatNumber  uint32    `json:"boat_num"`
	Name        string    `json:"name"`
	Year        uint32    `json:"year"`
	Length      uint32    `json:"length"`
	BuildNumber string    `json:"build_num"`
	BoatURL     string    `json:"boat_url"`
	Owner       string    `json:"owner"`
	OwnerURL    string    `json:"owner_url"`
	Port        string    `json:"port"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Created     time.Time `json:"created"`
	Modified    time.Time `json:"modified"`
}

func main() {

	if len(os.Args) >= 2 {
		cmd = os.Args[1]
	}

	fmt.Printf("Concordia Database\n")
	fmt.Printf("------------------\n")

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		fmt.Printf("Problem opening the database (%s): %s", dbFile, err)
		return
	}

	switch cmd {
	case "create":
		fmt.Printf("Creating (%s)...\n", dbFile)
		err = create(db)
		if err != nil {
			fmt.Printf("Problem creating DB: %s\n", err)
			return
		}
		err = loadBoats(db)
		fmt.Printf("Loading boats...\n")
		if err != nil {
			fmt.Printf("Problem loading boats: %s\n", err)
			return
		}
	case "delete":
		fmt.Printf("Deleting (%s)...\n", dbFile)
		err = os.Remove(dbFile)
		if err != nil {
			fmt.Printf("Problem deleting DB: %s\n", err)
			return
		}
	case "export":
		err = export(db)
		if err != nil {
			fmt.Printf("Problem exporting DB: %s\n", err)
			return
		}
	case "list":
		err = list(db)
		if err != nil {
			fmt.Printf("Problem listing DB: %s\n", err)
			return
		}
	default:
		fmt.Printf("Unexpected command '%s'. Expecting create, delete, export, or list.\n", cmd)
		return
	}
}

// boat, image, link, owner
func create(db *sql.DB) (err error) {
	// create boat table
	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS boat (
	id          INTEGER             PRIMARY KEY  AUTOINCREMENT,
	boat_num    INTEGER   NOT NULL  UNIQUE,
	name	    TEXT      NOT NULL  DEFAULT '',
	year	    INTEGER   NOT NULL  DEFAULT 0,
	length      INTEGER   NOT NULL  DEFAULT 0,
	build_num   TEXT      NOT NULL  DEFAULT '',
	boat_url    TEXT      NOT NULL  DEFAULT '',
	owner       TEXT      NOT NULL  DEFAULT '',
	owner_url   TEXT      NOT NULL  DEFAULT '',
	port        TEXT      NOT NULL  DEFAULT '',
	latitude    REAL      NOT NULL  DEFAULT 0,
	longitude   REAL      NOT NULL  DEFAULT 0,
	created     INTEGER   NOT NULL  DEFAULT 0,
	modified    INTEGER   NOT NULL  DEFAULT 0)`)
	if err != nil {
		return
	}
	_, err = stmt.Exec()
	return
}

func export(db *sql.DB) (err error) {
	return
}

func list(db *sql.DB) (err error) {
	return
}

// Read boats data from JSON file and insert into DB.
func loadBoats(db *sql.DB) (err error) {
	data, err := os.ReadFile(boatsFile)
	if err != nil {
		return
	}
	var boats []Boat
	err = json.Unmarshal(data, &boats)
	if err != nil {
		return
	}

	for _, b := range boats {
		err = insertBoat(db, b)
		if err != nil {
			return
		}
	}
	return
}

// We are passing db reference connection from main to our method with other parameters
func insertBoat(db *sql.DB, b Boat) (err error) {
	q := `
	INSERT INTO boat (
		id,
		boat_num,
		name)
	VALUES (?, ?, ?)`
	stmt, err := db.Prepare(q)
	if err != nil {
		return
	}
	_, err = stmt.Exec(b.ID, b.BoatNumber, b.Name)
	return
}
