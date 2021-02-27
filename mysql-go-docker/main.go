package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// go get -u github.com/go-sql-driver/mysql if package is missing!

func main() {
	// Connect to the database with the name of the database container and it's login details.
	fmt.Println("Connecting to db")
	conn, err := sql.Open("mysql", "root:mypassword@tcp(db:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// MySQL server isn't fully active yet.
	// Block until connection is accepted. This is a docker problem with v3 & container doesn't start
	// up in time.
	for conn.Ping() != nil {
		fmt.Println("Attempting connection to db")
		time.Sleep(5 * time.Second)
	}
	fmt.Println("Connected")

	// Optional: Below is a quick demo.
	// drop table if it already exists
	fmt.Println("Dropping table")
	_, err = conn.Exec(`DROP TABLE IF EXISTS People;`)
	if err != nil {
		panic(err)
	}

	// create a new table
	fmt.Println("Creating table")
	_, err = conn.Exec(`
	CREATE TABLE People (
		ID int,
		LastName varchar(255),
		FirstName varchar(255),
		Address varchar(255),
		City varchar(255)
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Insert into the new table
	fmt.Println("Inserting person")
	_, err = conn.Exec("INSERT INTO People VALUES (1, 'test', 'test', 'test', 'test');")
	if err != nil {
		log.Fatal(err)
	}

	// Create struct to store data assuming non NULL values for testing purposes.
	var person struct {
		ID        int
		LastName  string
		FirstName string
		Address   string
		City      string
	}

	// Get all the users
	fmt.Println("Getting person")
	result, err := conn.Query("SELECT * FROM People;")
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()

	// Get the results and store them in person.
	if result.Next() {
		err = result.Scan(
			&person.ID,
			&person.LastName,
			&person.FirstName,
			&person.Address,
			&person.City,
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v", person)
	}

	// For testing only - sleep to keep container alive.
	// time.Sleep(1 * time.Minute)
}
