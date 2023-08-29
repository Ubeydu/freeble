package main

import (
	"database/sql"
	"fmt"
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
		return
	}

	db, err := sql.Open("sqlite3", DB_NAME)
	if err != nil {
		log.Fatal(err.Error())
	}

	switch os.Args[1] {
	case "add_user":
		if len(os.Args) < 4 {
			log.Fatal("add_user <username> <password>")
		}
		err := s5_dbase.AddUser(db, os.Args[2], os.Args[3])
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("User Added :  %s", os.Args[2])
	case "login":
		if len(os.Args) < 4 {
			log.Fatal("login <username> <password>")
		}
		n, err := s5_dbase.CheckLogin(db, os.Args[2], os.Args[3])
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("Login = %t\n", n)
	default:
		fmt.Printf("TODO: Help Info ...")
	}
}
