package dbconfig

import (
	"github.com/nridwan/config/configutil"
	"github.com/nridwan/sys/dbutil"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

//LoadConfiguration to load DB config
func LoadConfiguration() {
	dbutil.AddConfig("default", dbutil.DbProfile{
		Connection: configutil.Getenv("DB_CONNECTION", ""),
		Host:       configutil.Getenv("DB_HOST", ""),
		Port:       configutil.Getenv("DB_PORT", ""),
		Database:   configutil.Getenv("DB_DATABASE", ""),
		Username:   configutil.Getenv("DB_USERNAME", ""),
		Password:   configutil.Getenv("DB_PASSWORD", ""),
		Locale:     configutil.Getenv("DB_LOCALE", "")})
	boil.SetDB(dbutil.Default())
}
