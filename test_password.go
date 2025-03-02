package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// 🔥 Your input password
	inputPassword := "sayyid123"

	// 🔥 Manually hash the password
	newHash, err := bcrypt.GenerateFromPassword([]byte(inputPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return
	}
	fmt.Println("🔹 Freshly Hashed Password:", string(newHash))

	// 🔥 Stored hash from MySQL (replace this with the actual hash from your DB)
	storedHash := "$2a$10$4mxZOq.nbFwDbGPyxVh9aujtRUgBFxctSTddFCE8aHO3jrZsd.Kr2"

	// 🔥 Compare stored hash with the input password
	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(inputPassword))
	if err != nil {
		fmt.Println("❌ Password does NOT match!")
	} else {
		fmt.Println("✅ Password matches!")
	}
}
