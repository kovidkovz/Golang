package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

var key = "cpscertgentopsecrethashSaltstringthing"

func createHash(s []byte) string {
	salted := append(s, []byte(key)...)               // Add salt
	hash := sha256.Sum256(salted)                     // SHA-256 hash
	return hex.EncodeToString(hash[:])                // Hex string
}

func main() {
	inputString := "SWMTKN-1-4f8z5y7x3w2v1u0t9s8r6q5p4n3m2l1k-j8h7g6f5d4s3a2b1c0"

	byteInput := []byte(inputString)                  // Convert to []byte
	hashedString := createHash(byteInput)             // Generate hash

	fmt.Println("Final Hashed Hex String:")
	fmt.Println(hashedString)
}
