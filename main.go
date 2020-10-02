package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pressly/goose"

	"golang-starter/bootstrap"
	"golang-starter/db/migrations"
	"golang-starter/router"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
)

func main() {

	defer bootstrap.App.DB.Close()

	//Print Usage For This Program
	flags.Usage = usage
	flags.Parse(os.Args[1:])

	args := flags.Args()

	// if len(args) < 1 {
	// 	flags.Usage()
	// 	return
	// }

	dir := "db/migrations"

	//Run Program As Server
	if args[0] == "run" {
		fmt.Println("Golang Program Starter")

		log.Printf(" This Program Run In {ENV : %s}", bootstrap.App.ENV)

		e := router.New()

		e.Start(":8000")
		os.Exit(0)
	}

	//Run Seeder
	if args[0] == "seed" {
		log.Printf("ENV : %s", bootstrap.App.ENV)
		migrations.Seed()
		os.Exit(0)
	}

	// GOOSE For MIGRATION Package

	if len(args) > 1 && args[0] == "create" {
		if err := goose.Run("create", nil, dir, args[1:]...); err != nil {
			log.Fatalf("goose run: %v", err)
		}
		return
	}

	if len(args) < 3 {
		flags.Usage()
		return
	}

	if args[0] == "-h" || args[0] == "--help" {
		flags.Usage()
		return
	}

	driver, dbstring, command := args[0], args[1], args[2]

	switch driver {
	case "postgres", "mysql", "sqlite3", "redshift":
		if err := goose.SetDialect(driver); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("%q driver not supported\n", driver)
	}

	switch dbstring {
	case "":
		log.Fatalf("-dbstring=%q not supported\n", dbstring)
	default:
	}

	if driver == "redshift" {
		driver = "postgres"
	}

	db, err := sql.Open(driver, dbstring)
	if err != nil {
		log.Fatalf("-dbstring=%q: %v\n", dbstring, err)
	}

	arguments := []string{}
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}

	if err := goose.Run(command, db, dir, arguments...); err != nil {
		log.Fatalf("goose run: %v", err)
	}
}

// FOR GOOSE

func usage() {
	log.Print(usagePrefix)
	flags.PrintDefaults()
	log.Print(usageCommands)
}

var (
	usagePrefix = `
Usage for Running Server: 
	go run main.go run

Usage for Running as Worker: 
	go run main.go start_worker WORKERNAME QUEUENAME

	Examples:
		go run main.go start_worker worker_1 queue_to_listen

Usage for Migrate: 
	go run main.go [OPTIONS] DRIVER DBSTRING COMMAND

	Drivers:
		postgres
		mysql
		sqlite3
		redshift
	Examples:
		go run main.go sqlite3 ./foo.db status
		go run main.go sqlite3 ./foo.db create init sql
		go run main.go sqlite3 ./foo.db create add_some_column sql
		go run main.go sqlite3 ./foo.db create fetch_user_data go
		go run main.go sqlite3 ./foo.db up
		go run main.go postgres "user=postgres dbname=postgres sslmode=disable" status
		go run main.go mysql "user:password@/dbname?parseTime=true" status
		go run main.go redshift "postgres://user:password@qwerty.us-east-1.redshift.amazonaws.com:5439/db" status
	Options:
	`

	usageCommands = `
	Commands:
		run                  Running HTTP server
		up                   Migrate the DB to the most recent version available
		up-to VERSION        Migrate the DB to a specific VERSION
		down                 Roll back the version by 1
		down-to VERSION      Roll back to a specific VERSION
		redo                 Re-run the latest migration
		status               Dump the migration status for the current DB
		version              Print the current version of the database
		create NAME [sql|go] Creates new migration file with next version
`
)
