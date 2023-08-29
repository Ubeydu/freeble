package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

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
		fmt.Printf("User Added :  %s\n", os.Args[2])
	case "login":
		if len(os.Args) < 4 {
			log.Fatal("login <username> <password>")
		}
		n, err := s5_dbase.CheckLogin(db, os.Args[2], os.Args[3])
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("Login = %t\n", n)
	case "add_item":
		if len(os.Args) < 6 {
			log.Fatal("add_item <giver_id> <item_name> <description> <image_location>")
		}
		fdata, err := os.ReadFile(os.Args[5])
		if err != nil {
			log.Fatalf("Read Image Error : %w", err)
		}
		err = s5_dbase.AddItem(db, os.Args[2], os.Args[3], os.Args[4], fdata)
		if err != nil {
			log.Fatalf("Could not add Item: %w", err)
		}
		fmt.Printf("Item Added : %s\n", os.Args[3])
	case "get_image":
		if len(os.Args) < 3 {
			log.Fatal("get_image <image_id>")
		}
		item_id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal("image_id should be int")
		}
		err = s5_dbase.GetImage(db, item_id, os.Stdout)
		if err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Printf("TODO: Help Info ...")
	}
}
