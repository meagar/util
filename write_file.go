package util

import (
	"io"
	"os"
)

// WriteFile writes the given string to the given file.
// Panics on error.
func WriteFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		panic(err)
	}
	return file.Sync()
}
