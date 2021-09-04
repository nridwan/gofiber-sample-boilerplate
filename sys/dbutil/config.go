package dbutil

import (
	"database/sql"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
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
}

var profiles = map[string]*sql.DB{}

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
	suffix := "?loc="
	if config.Locale == "" {
		suffix += "UTC"
	} else {
		suffix += url.QueryEscape(config.Locale)
	}
	if config.Connection == "mysql" {
		suffix += "&parseTime=true"
	}
	profiles[profName], _ = sql.Open(config.Connection, config.Username+":"+config.Password+"@tcp("+config.Host+":"+config.Port+")/"+config.Database+suffix)
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
