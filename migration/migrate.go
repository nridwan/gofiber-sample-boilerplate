package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/nridwan/config"
	"github.com/nridwan/config/configutil"
	"github.com/nridwan/sys/dbutil"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	count := flag.Int("count", 0, "Decide migration count")
	path := flag.String("path", "migration", "Decide migration path")
	force := flag.Int("force", 0, "Force migration")
	taskUp := flag.Bool("up", false, "Apply the migration")
	taskDown := flag.Bool("down", false, "Revert the migration")
	var driver database.Driver
	var err error
	var m *migrate.Migrate
	flag.Parse()
	if configutil.Getenv("PORT", "") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	config.LoadAllConfiguration()
	if dbutil.GetConnection("default") == "mysql" {
		driver, err = mysql.WithInstance(dbutil.Default(), &mysql.Config{})
	} else if dbutil.GetConnection("default") == "postgres" {
		driver, err = postgres.WithInstance(dbutil.Default(), &postgres.Config{})
	} else {
		return
	}
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	m, err = migrate.NewWithDatabaseInstance(
		"file://"+*path,
		dbutil.GetConnection("default"), driver)
	if err != nil {
		println(err.Error())
		return
	}
	if *force != 0 {
		m.Force(*force)
	} else if *taskDown {
		if *count == 0 {
			err = m.Steps(-1)
		} else {
			err = m.Steps(-*count)
		}
	} else if *taskUp {
		if *count == 0 {
			err = m.Up()
		} else {
			err = m.Steps(*count)
		}
	} else {
		version, _, err := m.Version()
		if err != nil {
			println(err.Error())
			return
		}
		println("Current migration version: ", version)
	}
	if err != nil {
		println(err.Error())
	}
}
