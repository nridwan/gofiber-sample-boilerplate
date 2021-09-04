package configutil

import "os"

//Getenv is extension of os.Getenv
func Getenv(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
