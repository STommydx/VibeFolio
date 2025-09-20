package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

func main() {
	db, err := sql.Open("sqlite3", "./data/vibefolio.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Set the directory for migrations
	goose.SetBaseFS(nil)
	if err := goose.SetDialect("sqlite3"); err != nil {
		log.Fatal(err)
	}

	// Run migrations
	if len(os.Args) > 1 && os.Args[1] == "down" {
		fmt.Println("Rolling back migrations...")
		if err := goose.DownTo(db, "./migrations", 0); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Running migrations...")
		if err := goose.Up(db, "./migrations"); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Migrations completed successfully!")
}