package config

import "github.com/nridwan/config/dbconfig"

//LoadAllConfiguration is extension of os.Getenv
func LoadAllConfiguration() {
	dbconfig.LoadConfiguration()
}
