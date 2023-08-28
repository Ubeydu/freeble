package main

import (
	"log"
	"os"

	s5_dbase "example.com/hello"
)

const DB_NAME = "freeblebase.db"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("HELP INFO ...")
	}
	if os.Args[1] == "create" {
		err := s5_dbase.SetupDB(DB_NAME)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Printf("Database Created\n")
	}
}
