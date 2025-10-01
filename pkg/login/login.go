package login

import (
	"log"

	"github.com/mclacore/passh/pkg/collection"
	"github.com/mclacore/passh/pkg/store"
	"gorm.io/gorm"
)

// LoginItem type containing all data for a login item.
type LoginItem struct {
	gorm.Model
	ItemName     string `gorm:"not null;uniqueIndex:idx_item_collection"`
	Username     string
	Password     string
	URL          string
	CollectionID int                   `gorm:"not null;uniqueIndex:idx_item_collection"`
	Collection   collection.Collection `gorm:"foreignKey:CollectionID;constraint:OnDelete:CASCADE;"`
}

var loginItem LoginItem

// Create a login item.
func CreateLoginItem(db *gorm.DB, item LoginItem) error {
	automigrateDB()
	result := db.Create(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Fetch data for a login item.
func GetLoginItem(db *gorm.DB, itemName string, colId int) (*LoginItem, error) {
	result := db.Where("item_name = ? AND collection_id = ?", itemName, colId).Find(&loginItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &loginItem, nil
}

// Update data for a login item.
func UpdateLoginItem(db *gorm.DB, loginItem *LoginItem) error {
	result := db.Save(loginItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// List all login items in a collection.
func ListLoginItems(db *gorm.DB, colId int) (*[]LoginItem, error) {
	var loginItems []LoginItem

	// add listing by item-name
	result := db.Select("item_name").
		Where(&LoginItem{CollectionID: colId}).
		Order("item_name asc").
		Find(&loginItems)
	if result.Error != nil {
		return nil, result.Error
	}

	return &loginItems, nil
}

// Delete a login item.
func DeleteLoginItem(db *gorm.DB, itemName string, colId int) error {
	result := db.Where(&LoginItem{ItemName: itemName, CollectionID: colId}).Delete(&loginItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Assign a login item to a collection.
func AssignCollection(db *gorm.DB, itemName, colName string) error {
	var col collection.Collection

	colId := col.ID

	result := db.Where(&LoginItem{ItemName: itemName}).Set("collection_id = ?", colId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func automigrateDB() {
	db := store.DB()

	if err := db.AutoMigrate(&LoginItem{}); err != nil {
		log.Fatal(err)
	}
}
