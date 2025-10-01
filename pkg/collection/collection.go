package collection

import (
	"log"

	"github.com/mclacore/passh/pkg/database"
	"gorm.io/gorm"
)

// Collection type containing a collection name.
type Collection struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"` // gorm:unique is not working? maybe add checking in create step
}

var collection Collection

// Creates a collection.
func CreateCollection(db *gorm.DB, col Collection) error {
	automigrateDB()
	result := db.Create(&col)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Fetches a collection using a collection name.
func GetCollectionByName(db *gorm.DB, colName string) (*Collection, error) {
	var col Collection
	result := db.Where("name = ?", colName).Find(&col)
	if result.Error != nil {
		return nil, result.Error
	}
	return &col, nil
}

// Fetches a collection using a collection ID.
func GetCollectionById(db *gorm.DB, colId int) (*Collection, error) {
	result := db.Where("id = ?", colId).Find(&collection)
	if result.Error != nil {
		return nil, result.Error
	}
	return &collection, nil
}

//TODO: need to build this into cmd/collection.go

// Update a collection nametes a collection using a collection name.
func UpdateCollection(db *gorm.DB, colName string) (*Collection, error) {
	result := db.Where("name = ?", colName).Find(&collection)
	if result.Error != nil {
		return nil, result.Error
	}
	return &collection, nil
}

// List all collections.
func ListCollections(db *gorm.DB) (*[]Collection, error) {
	var collections []Collection

	result := db.Select("name").
		Order("name asc").
		Find(&collections)
	if result.Error != nil {
		return nil, result.Error
	}
	return &collections, nil
}

// Delete a collection using a collection name.
func DeleteCollection(db *gorm.DB, colName string) error {
	result := db.Where("name = ?", colName).Delete(&collection)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func automigrateDB() {
	db, err := database.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&Collection{})
	if err != nil {
		log.Fatal(err)
	}
}
