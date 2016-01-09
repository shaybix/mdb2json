// Copyright 2016 Shaybix. All Rights Reserved.
// A simple tool that converts MDB files to JSON

package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var format = flag.String("f", "json", "Choose either json or sql for format")
var dir = flag.String("d", "mdb", "Select the directory with the bok files")
var outputFile = flag.String("o", "db.sqlite", "Output file to dump your database content")

func init() {

	flag.Parse()

}

func main() {

	files := crawlDir("mdb")

	db, err := initSqlDb(*outputFile)
	if err != nil {

		log.Fatal("Could not initialise database: ", err)
	}

	for _, file := range files {

		if file == "" {
			continue
		}

		err := schema(file, db)

		if err != nil {
			log.Fatal(err)
		}
	}

}

// CrawlDir crawls a directory and return a slice of filenames
func crawlDir(dirname string) []string {

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Read files from directory")

	list := make([]string, len(files))

	for _, f := range files {

		list = append(list, f.Name())
	}

	return list

}

// schema exports the mdb-schema in json format
func schema(filename string, db *sql.DB) error {

	out, err := exec.Command("mdb-schema", *dir+"/"+filename, "mysql").Output()
	if err != nil {

		log.Fatal("Could not execute the command: ", err)

		return err

	}

	queries := strings.Split(string(out), ";")

	for _, query := range queries {

		_, err := db.Exec(query)

		if err != nil {
			log.Fatalf("Could not execute the query transaction: %v", err)
			return err

		}

	}

	defer db.Close()

	return err

}

// initSqlDb ...
func initSqlDb(dbFile string) (*sql.DB, error) {

	// ...

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal("Error opening the sqlite database file: ", err)
		return nil, err
	}

	err = db.Ping()

	if err != nil {

		log.Fatal("Could not ping the database: ", err)
		return nil, err

	}

	return db, nil
}

//dump dumps data into a Sql database and  returning an error
func dumpToSql(filename string) error {

	var err error

	err = prepareEnv()
	if err != nil {

		log.Fatal(err)
		return err
	}

	// TODO: Dump the data to the Sqlite DB file

	return nil

}

// prepareEnv prepares the environment variables
func prepareEnv() error {

	var err error

	err = os.Setenv("MDB_JET3_CHARSET", "cp1256")
	if err != nil {
		log.Fatal(err)
		return err

	}

	return nil

}
