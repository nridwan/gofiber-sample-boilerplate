package dbutil

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

//DbProfile = db configuration
type DbProfile struct {
	Connection string
	Host       string
	Port       string
	Database   string
	Username   string
	Password   string
	Locale     string
	DbUrl      string
}

var profiles = map[string]*sql.DB{}
var connections = map[string]string{}

func (prof *DbProfile) connect() (*sql.DB, error) {
	db, err := sql.Open(prof.Connection, prof.Username+":"+prof.Password+"@tcp("+prof.Host+":"+prof.Port+")/"+prof.Database)
	if err != nil {
		return nil, err
	}

	return db, nil
}

//Do connect db
func (prof *DbProfile) Do(onConnect func(*sql.DB), onError func(error)) {
	db, err := prof.connect()
	if err != nil {
		onError(err)
		return
	}
	defer db.Close()
	onConnect(db)
}

//AddConfig = add configuration
func AddConfig(profName string, config DbProfile) {
	if config.Connection == "mysql" {
		suffix := "?loc="
		if config.Locale == "" {
			suffix += "UTC"
		} else {
			suffix += url.QueryEscape(config.Locale)
		}
		suffix += "&parseTime=true&multiStatements=true"
		profiles[profName], _ = sql.Open(config.Connection, config.Username+":"+config.Password+"@tcp("+config.Host+":"+config.Port+")/"+config.Database+suffix)
	} else if config.Connection == "postgres" {
		var connInfo string
		if config.DbUrl != "" {
			connInfo = config.DbUrl
		} else {
			connInfo = fmt.Sprintf("host='%s' port=%s user='%s' "+
				"password='%s' dbname='%s' sslmode=disable",
				config.Host, config.Port, config.Username, config.Password, config.Database)
		}
		profiles[profName], _ = sql.Open(config.Connection, connInfo)
		err := profiles[profName].Ping()
		if err != nil {
			panic(err)
		}
	}
	connections[profName] = config.Connection
}

//RemoveConfig = remove configuration
func RemoveConfig(profName string) {
	delete(profiles, profName)
}

//Default : get default DB profile
func Default() *sql.DB {
	return profiles["default"]
}

//Get : get DB profile
func Get(profName string) *sql.DB {
	return profiles[profName]
}

func Migrate(profName string) {
	var driver database.Driver
	var err error
	var m *migrate.Migrate
	if connections[profName] == "mysql" {
		driver, err = mysql.WithInstance(Get(profName), &mysql.Config{})
	} else if connections[profName] == "postgres" {
		driver, err = postgres.WithInstance(Get(profName), &postgres.Config{})
	} else {
		return
	}
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	m, err = migrate.NewWithDatabaseInstance(
		"file://migration",
		connections[profName], driver)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	m.Up()
}
