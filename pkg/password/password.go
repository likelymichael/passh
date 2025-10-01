package password

import (
	"crypto/rand"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/mclacore/passh/pkg/config"
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

func MasterPasswordTimeout() {
	timeVal, timeValErr := config.LoadConfigValue("auth", "timeout")
	if timeValErr != nil {
		log.Printf("Error loading timeout value: %v", timeValErr)
		os.Exit(2)
	}

	if timeVal == "" {
		config.SaveConfigValue("auth", "timeout", "900")
	}

	timeout, timeoutErr := strconv.Atoi(timeVal)
	if timeoutErr != nil {
		log.Printf("Error converting timeout string to int: %v", timeoutErr)
		os.Exit(2)
	}

	time.Sleep(time.Duration(timeout) * time.Second)
	config.SaveConfigValue("auth", "temp_pass", "")
}
