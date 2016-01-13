// Copyright 2016 Shaybix. All Rights Reserved.
// A simple tool that converts MDB files to JSON

package main

import (
	"flag"
	"log"
	"os"
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

	for _, file := range files {

		if file == "" {
			continue
		}

		newFile := strings.Split(file, ".")[0]

		newFile = newFile + ".sqlite"
		f, err := os.Create(newFile)

		defer f.Close()

		if err != nil {
			log.Fatalf("Could not create file: %v", err)
		}

	}

}
