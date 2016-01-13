package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

// CrawlDir crawls a directory and return a slice of filenames
func crawlDir(dirname string) []string {
	// Crawls the directory and loads all the files

	allFiles, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Read files from directory")

	files := make([]string, len(allFiles))

	for _, f := range allFiles {

		files = append(files, f.Name())
	}

	return files

}
