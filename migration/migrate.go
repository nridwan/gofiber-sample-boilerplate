package main

import (
	"flag"
	"log"
	"os"

	"github.com/nridwan/config"
	"github.com/nridwan/sys/dbutil"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	count := flag.Int("count", 0, "Decide migration count")
	path := flag.String("path", "migration", "Decide migration path")
	force := flag.Int("force", 0, "Force migration")
	flag.Parse()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.LoadAllConfiguration()
	args := os.Args
	task := "up"
	if len(args) > 1 && (args[1] == "up" || args[1] == "down" || args[1] == "version") {
		task = args[1]
	}
	driver, err := mysql.WithInstance(dbutil.Default(), &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+*path,
		"mysql", driver)
	if err != nil {
		println(err.Error())
		return
	}
	if *force != 0 {
		m.Force(*force)
	} else if task == "up" {
		if *count == 0 {
			err = m.Up()
		} else {
			err = m.Steps(*count)
		}
	} else if task == "down" {
		if *count == 0 {
			err = m.Steps(-1)
		} else {
			err = m.Steps(-*count)
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
