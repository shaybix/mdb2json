// Copyright 2016 Shaybix. All Rights Reserved.
// A simple tool that converts MDB files to JSON

package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	filename := "mdb/file.bok"

	cmd := exec.Command("mdb-schema", filename)

	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(output))

	log.Printf("Waiting for command to finish....")

	err = cmd.Run()

	if err != nil {
		log.Printf("Command finished with error: %v", err)
	} else {
		log.Printf("Command executed successfully!")
	}

}
