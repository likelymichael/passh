package database

import (
	"fmt"

	"github.com/mclacore/passh/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connects to Passh database.
func ConnectToDB() (*gorm.DB, error) {
	user, userErr := config.LoadConfigValue("auth", "username")
	if userErr != nil {
		return nil, userErr
	}

	pass, passErr := config.LoadConfigValue("auth", "persist_pass")
	if pass == "" {
		pass, passErr = config.LoadConfigValue("auth", "temp_pass")
	}

	if passErr != nil {
		return nil, passErr
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s sslmode=disable",
		"localhost",
		"5432",
		user,
		pass,
		"postgres")

	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		return nil, dbErr
	}

	return db, nil
}
