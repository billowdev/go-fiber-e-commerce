package domain

import (
	"time"

	"github.com/billowdev/go-fiber-e-commerce/internal/core/domain"
	"github.com/billowdev/go-fiber-e-commerce/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserProfile struct {
	domain.BaseModel
	UserID           uuid.UUID    `json:"user_id" gorm:"not null"`         // References the User table, linking the profile to a specific user
	ProfilePicture   string       `json:"profile_picture" gorm:"size:255"` // URL or path to the user's profile picture
	Preferences      domain.JSONB `json:"preferences"`                     // JSON object storing user-specific preferences and settings
	DateOfBirth      *time.Time   `json:"date_of_birth"`                   // User's date of birth
	SocialMediaLinks domain.JSONB `json:"social_media_links"`              // JSON object storing links to user's social media profiles
}

var TNUserProfile = "user_profiles"

// TableName sets the insert table name for UserProfile struct
func (UserProfile) TableName() string {
	return TNUserProfile
}
func (u *UserProfile) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
		return err
	}
	return nil
}
