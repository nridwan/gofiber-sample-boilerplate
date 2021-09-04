package main

import (
	"flag"
	"log"

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
	taskUp := flag.Bool("up", false, "Apply the migration")
	taskDown := flag.Bool("down", false, "Revert the migration")
	flag.Parse()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.LoadAllConfiguration()
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
