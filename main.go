// Copyright 2016 Shaybix. All Rights Reserved.
// A simple tool that converts MDB files to JSON

package main

import (
	"flag"
	"log"
	"os"
	"strings"
)

var format = flag.String("f", "json", "Choose either json or sql for format")
var dir = flag.String("d", "mdb", "Select the directory with the bok files")
var outputFile = flag.String("o", "db.sqlite", "Output file to dump your database content")

func init() {

	flag.Parse()

}

func main() {

	files := crawlDir("mdb")

	for _, file := range files {

		if file == "" {
			continue
		}

		newFile := strings.Split(file, ".")[0]

		dbFile := newFile + ".sqlite"
		f, err := os.Create(dbFile)

		defer f.Close()

		if err != nil {
			log.Fatalf("Could not create file: %v", err)
		}

		db, err := initDB(dbFile)

		if err != nil {
			log.Fatalf("could not initialise Sqlite database: %v", err)

		}

		err = schema(dbFile, db)

		if err != nil {
			log.Printf("Not able to set the schema: %v", err)
		}

		err = dumpToSQL(dbFile, db)

		if err != nil {
			log.Fatalf("Could not dump SQL to Database: %v", err)

		}

	}

}
