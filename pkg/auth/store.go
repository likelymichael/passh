package auth

import (
	"context"
	"errors"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

var (
	ErrNoAuthRecord       = errors.New("auth: not initialized")
	ErrAlreadyInitialized = errors.New("auth: already initialized")
)

type KDFParams struct {
	Algo      string `json:"algo"`
	Time      uint32 `json:"time"`
	MemoryKiB uint32 `json:"memory_kib"`
	Parallel  uint8  `json:"parallel"`
	KeyLen    uint32 `json:"key_len"`
}

// AuthRecord is a singleton row (id=1).
type AuthRecord struct {
	ID               uint   `gorm:"primaryKey;default:1"`
	Version          uint   `gorm:"not null;default:1"`
	Salt             []byte `gorm:"size:16"`
	KDFParams        datatypes.JSON
	WrapNonse        []byte `gorm:"size:24"`
	EncryptedDataKey []byte
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (AuthRecord) TableName() string {
	return "auth_metadata"
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&AuthRecord{})
}

// Exists checks if the AuthRecord exists.
func Exists(ctx context.Context, db *gorm.DB) (bool, error) {
	var n int64
	if err := db.WithContext(ctx).Model(&AuthRecord{}).Count(&n).Error; err != nil {
		return false, err
	}
	return n > 0, nil
}

// Get fetches the AuthRecord if found.
func Get(ctx context.Context, db *gorm.DB) (*AuthRecord, error) {
	var r AuthRecord
	if err := db.WithContext(ctx).First(&r, 1).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNoAuthRecord
		}
		return nil, err
	}
	return &r, nil
}

// Create inserts the singleton row during onboarding.
func Create(ctx context.Context, db *gorm.DB, r *AuthRecord) error {
	exists, err := Exists(ctx, db)
	if err != nil {
		return err
	}

	if exists {
		return ErrAlreadyInitialized
	}
	r.ID = 1
	return db.WithContext(ctx).Create(r).Error
}

// Update updates an existing AuthRecord.
func Update(ctx context.Context, db *gorm.DB, r *AuthRecord) error {
	if r.ID == 0 {
		r.ID = 1
	}
	return db.WithContext(ctx).Model(&AuthRecord{}).Where("id = 1").Updates(r).Error
}
