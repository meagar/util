// util contains some random utilities I reach for in many one-off things I write
package util

import (
	"bufio"
	"log"
	"os"
)

// EachLine iterates over each line of a file, calling the fn callback
func EachLine(filename string, fn func(string)) {
	file, err := os.Open(filename)
	if err != nil {
		log.Panicf("EachLine: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fn(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Panic(err)
	}
}

// Env loads an environment variable, and panics if it is unset or blank
func Env(name string) string {
	value, ok := os.LookupEnv(name)

	if !ok || value == "" {
		log.Panicf("ENV var is blank or not set: %s", name)
	}

	return value
}
