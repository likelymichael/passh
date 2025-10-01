package store

import (
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	err  error
	once sync.Once
)

// Open sets up the global DB once. Safe to call multiple times.
func Open(path string) (*gorm.DB, error) {
	once.Do(func() {
		db, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	})
	return db, err
}

// DB returns the initialized handle. Panics if not opened.
func DB() *gorm.DB {
	if db == nil {
		panic("store.DB: not opened")
	}
	return db
}
