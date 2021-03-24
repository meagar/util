package util

import (
	"log"
	"os"
)

// Env loads an environment variable, and panics if it is unset or blank
func Env(name string) string {
	value, ok := os.LookupEnv(name)

	if !ok || value == "" {
		log.Panicf("ENV var is blank or not set: %s", name)
	}

	return value
}
