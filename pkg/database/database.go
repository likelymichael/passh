package database

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/mclacore/passh/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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

	go databaseTimeout()

	return db, nil
}

func WizardPasswordSet(input string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s database=%s sslmode=disable",
		"localhost",
		"5432",
		"postgres")

	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		return nil, dbErr
	}

	pwQuery := fmt.Sprintf(`ALTER USER %q WITH PASSWORD '%s';`, "postgres", input)
	if pwErr := db.Exec(pwQuery).Error; pwErr != nil {
		return nil, pwErr
	}

	return db, nil
}

func databaseTimeout() {
	timeVal, timeValErr := config.LoadConfigValue("auth", "timeout")
	if timeValErr != nil {
		log.Print(timeValErr)
	}

	if timeVal == "" {
		config.SaveConfigValue("auth", "timeout", "900")
	}

	timeout, timeoutErr := strconv.Atoi(timeVal)
	if timeoutErr != nil {
		log.Print(timeoutErr)
	}

	time.Sleep(time.Duration(timeout) * time.Second)
	config.SaveConfigValue("auth", "temp_pass", "")
}
