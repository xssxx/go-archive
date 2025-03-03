package main

import (
	"fmt"
	"testptr/models"
	"testptr/services"
)

func main() {
	user1 := models.NewUser(1, "Alice", "password123", &services.Sha256Hasher{})
	user2 := models.NewUser(2, "Bob", "securepassword", &services.Sha512Hasher{})

	fmt.Printf("User1: %s -> %s\n", user1.Name, user1.HashPassword())
	fmt.Printf("User2: %s -> %s\n", user2.Name, user2.HashPassword())
}
