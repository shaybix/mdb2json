// Copyright 2016 Shaybix. All Rights Reserved.
// A simple tool that converts MDB files to JSON

package main

import (
	"io/ioutil"
	"log"
	"os/exec"
)

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

func main() {

	files := crawlDir("mdb")

	files = files[1:]

	var cmd *exec.Cmd

	for _, file := range files {

		if file == "" {
			continue
		}

		cmd = exec.Command("mdb-schema", "mdb/"+string(file))

		err := cmd.Run()

		if err != nil {
			log.Printf("Command finished with error: %v", err)
		} else {
			log.Printf("Command executed successfully! - File: %v", file)
		}

	}

}
