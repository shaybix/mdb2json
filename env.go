package main

import (
	"log"
	"os"
)

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
