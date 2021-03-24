package util

import (
	"bufio"
	"log"
	"os"
)

// EachLine iterates over each line of a file, calling the fn callback
// Panics on any error. Don't use in production.
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
