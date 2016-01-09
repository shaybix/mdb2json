// Copyright 2016 Shaybix. All Rights Reserved.
// A simple tool that converts MDB files to JSON

package main

import (
	"database/sql"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

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

	files = files[1:]

	for _, file := range files {

		if file == "" {
			continue
		}

		err := schema(file)

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
	list := make([]string, len(files))

	for _, f := range files {

		list = append(list, f.Name())
	}

	return list

}

// schema exports the mdb-schema in json format
func schema(filename string) error {

	var cmd *exec.Cmd
	var err error

	cmd = exec.Command("mdb-schema", filename)

	err = cmd.Run()

	if err != nil {
		log.Printf("Command finished with error: %v", err)
		return err
	} else {

		log.Printf("Command executed successfully - File: %v", filename)

	}
	return err

}

// initSqlDb ...
func initSqlDb(dbFile string) (*sql.DB, error) {

	// ...

	db, err := sql.Open("go-sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
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

	//

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
