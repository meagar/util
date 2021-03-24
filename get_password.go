package util

import (
	"fmt"
	"syscall"

	"golang.org/x/term"
)

// GetPassword prompts the user for a password without echoing as the user types
// Panics on error.
func GetPassword(prompt string) (password string) {
	fmt.Print(prompt)
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))

	if err != nil {
		panic(err)
	}
	fmt.Println("")

	password = string(bytePassword)
	return
}
