package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	hash := "$2a$10$wB5W3W.4pY4s4X/mY9J4c.6hX7bQpGg3fG.9J9.qXlU7Jq.j8Wl0q"
	password := "residen01"
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Printf("Err: %v\n", err)
}
