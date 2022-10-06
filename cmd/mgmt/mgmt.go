//
// Command to use to manage the Concordia MariaDB database.
// Actions
//   export - export the database to a JSON file.
//   list   - list contents of the database
//
// Tables
//   yawl                 - current boat information.
//   yawl_image          - boat images.
//   yawl_link           - boat links.
//   yawl_owner_history  - boat owner history.

package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	cmd = "list" // export, or list

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
}

func main() {
	if len(os.Args) >= 2 {
		cmd = os.Args[1]
	}

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/concordia")
	if err != nil {
		fmt.Printf("Problem connecting to the database: %s\n", err)
	}
	defer db.Close()

	switch cmd {
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
		fmt.Printf("Unexpected command '%s'. Expecting export or list.\n", cmd)
		return
	}
}

func export(db *sql.DB) (err error) {
	return
}

func list(db *sql.DB) (err error) {
	results, err := db.Query(`
	SELECT boat_num, SUM(num_images) AS num_images, SUM(num_links) AS num_links, SUM(num_owners) AS num_owners
	FROM (
	SELECT y.boat_num, count(yi.id) AS num_images, 0 AS num_links, 0 AS num_owners
	FROM yawl y
	LEFT JOIN yawl_image yi ON y.boat_num = yi.boat_num
	GROUP BY y.boat_num
	UNION
	SELECT y.boat_num, 0 AS num_images, count(yl.id) AS num_links, 0 AS num_owners
	FROM yawl y
	LEFT JOIN yawl_link yl ON y.boat_num = yl.boat_num
	GROUP BY y.boat_num
	UNION
	SELECT y.boat_num, 0 AS num_images, 0 AS num_links, count(yoh.id) AS num_owners
	FROM yawl y
	LEFT JOIN yawl_owner_history yoh ON y.boat_num = yoh.boat_num
	GROUP BY y.boat_num
	) AS b GROUP BY boat_num
	`)
	if err != nil {
		return
	}
	fmt.Printf("  | Boat | Images | Links | Owners |\n")
	fmt.Printf("  | ---- | ------ | ----- | ------ |\n")
	for results.Next() {
		var boat_num, num_images, num_links, num_owners uint32

		err = results.Scan(&boat_num, &num_images, &num_links, &num_owners)
		if err != nil {
			return
		}
		fmt.Printf("  |  %3d |    %3d |   %3d |    %3d |\n", boat_num, num_images, num_links, num_owners)
	}
	return
}
