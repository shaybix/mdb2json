package main

import (
	"database/sql"
	"log"
	"os/exec"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// ...
// initDB ..
func initDB(dbFile string) (*sql.DB, error) {

	DB, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatalf("Error opening Database connection: %v", err)
		return nil, err
	}
	return DB, nil
}

// schema exports the mdb-schema in json format
func schema(filename string, db *sql.DB) error {

	out, err := exec.Command("mdb-schema", *dir+"/"+filename, "sqlite").Output()
	if err != nil {

		log.Print("Could not execute the command mdb-schema: ", err)

		return err

	}

	queries := strings.Split(string(out), ";")

	for _, query := range queries {

		_, err := db.Exec(query)

		if err != nil {
			log.Printf("Could not execute the query transaction: %v", err)
			return err

		}

	}

	return nil

}

func dumpToSQL(filename string, db *sql.DB) error {

	err := prepareEnv()
	if err != nil {

		log.Fatal(err)
		return err
	}

	// TODO: Dump the data to the Sqlite DB file

	out, err := exec.Command("mdb-tables", "-1", *dir+"/"+filename).Output()
	if err != nil {
		log.Fatalf("Can not execute command mdb-tables: %s", err)
		return err
	}

	tables := strings.Split(string(out), "\n")
	log.Printf("Starting Data Insertion in: %s", filename)

	for _, table := range tables {

		out, err := exec.Command("mdb-export", "-I", "sqlite", *dir+"/"+filename, table).Output()
		if err != nil {

			log.Fatalf("unable to export mdb as Sql queries: %v", err)
			return err

		}

		queries := strings.Split(string(out), "\n")

		for _, query := range queries {
			_, err := db.Exec(query)
			if err != nil {
				// Currently throwing unrecognized token error
				//

				log.Fatalf("Unable to execute query: %v", err)
				return err
			}
		}
	}

	return nil

}
