package domain

import (
	"time"

	"github.com/billowdev/go-fiber-e-commerce/internal/core/domain"
	"github.com/billowdev/go-fiber-e-commerce/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshToken struct {
	domain.BaseModel
	UserID    uuid.UUID      `json:"user_id" gorm:"not null"`                     // References the User table to link the refresh token to a specific user
	Token     string         `json:"token" gorm:"size:255;not null"`              // JWT refresh token string used to renew user authentication sessions
	ExpiresAt time.Time      `json:"expires_at"`                                  // Timestamp indicating when the refresh token will expire
	Used      bool           `json:"used" gorm:"default:false"`                   // Indicates whether the refresh token has been used to issue a new token
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"` // Timestamp when the refresh token record was created
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"` // Timestamp when the refresh token record was last updated
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`                     // Timestamp for soft deletes
}

var TNRefreshToken = "refresh_tokens"

// TableName sets the insert table name for RefreshToken struct
func (RefreshToken) TableName() string {
	return TNRefreshToken
}

func (u *RefreshToken) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
		return err
	}
	return nil
}
