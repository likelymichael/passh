package password

import (
	"crypto/rand"
	"log"
	"math/big"
)

func GeneratePassword(length int, lowercase, uppercase, numbers, special bool) string {
	charset := []byte("")
	lower := []byte("abcdefghijklmnopqrstuvwxyz")
	upper := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	digits := []byte("0123456789")
	specials := []byte("!@#$%^&*()_+~><")

	if !lowercase {
		charset = append(charset, lower...)
	}

	if uppercase {
		charset = append(charset, upper...)
	}

	if numbers {
		charset = append(charset, digits...)
	}

	if special {
		charset = append(charset, specials...)
	}

	if len(charset) == 0 {
		log.Fatal("no character classes selected for password generation")
	}

	password := make([]byte, length)
	for i := range password {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			log.Fatalf("failed to generate secure random number: %v", err)
		}
		password[i] = charset[n.Int64()]
	}

	return string(password)
}
