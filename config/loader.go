package config

import (
	"github.com/nridwan/config/dbconfig"
	"github.com/nridwan/sys/jwtutil"
)

//LoadAllConfiguration is extension of os.Getenv
func LoadAllConfiguration() {
	dbconfig.LoadConfiguration()
	jwtutil.LoadConfiguration()
}
