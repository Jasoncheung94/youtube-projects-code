package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// Get this package if it's missing.
// go get -u github.com/lib/pq
// If you encounter problems like I did with a newer version of Go. Run the following:
// GO111MODULE="off" go get -u github.com/lib/pq
// Restart IDE

func main() {
	fmt.Println("connecting")
	// these details match the docker-compose.yml file.
	postgresInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"postgres", 5432, "user", "mypassword", "user")
	db, err := sql.Open("postgres", postgresInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	start := time.Now()
	for db.Ping() != nil {
		if start.After(start.Add(10 * time.Second)) {
			fmt.Println("failed to connect after 10 secs.")
			break
		}
	}
	fmt.Println("connected:", db.Ping() == nil)
	_, err = db.Exec(`DROP TABLE IF EXISTS COMPANY;`)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`CREATE TABLE COMPANY (ID INT PRIMARY KEY NOT NULL, NAME text);`)
	if err != nil {
		panic(err)
	}
	fmt.Println("table company is created")
}
