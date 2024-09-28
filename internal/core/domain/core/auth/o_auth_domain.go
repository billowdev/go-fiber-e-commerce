package domain

import (
	"time"

	"github.com/billowdev/go-fiber-e-commerce/internal/core/domain"
	"github.com/billowdev/go-fiber-e-commerce/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OAuthProvider struct {
	domain.BaseModel
	ProviderName string         `json:"provider_name" gorm:"size:100;not null"`      // Name of the OAuth provider (e.g., 'Google', 'Facebook')
	ClientID     string         `json:"client_id" gorm:"size:255;not null"`          // Client ID provided by the OAuth provider
	ClientSecret string         `json:"client_secret" gorm:"size:255;not null"`      // Client Secret provided by the OAuth provider
	CreatedAt    time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"` // Timestamp when the record was created
	UpdatedAt    time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"` // Timestamp when the record was last updated
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`                     // Timestamp for soft deletes
}

var TNOAuthProvider = "oauth_providers"

// TableName sets the insert table name for OAuthProvider struct
func (OAuthProvider) TableName() string {
	return TNOAuthProvider
}

func (o *OAuthProvider) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
		return err
	}
	return nil
}

type OAuthAccessToken struct {
	domain.BaseModel
	UserID      uuid.UUID      `json:"user_id" gorm:"not null"`                     // References the User table to link the token to a specific user
	ProviderID  uuid.UUID      `json:"provider_id" gorm:"not null"`                 // References the OAuthProvider table to identify the provider of the token
	AccessToken string         `json:"access_token" gorm:"size:255;not null"`       // The OAuth access token
	TokenType   string         `json:"token_type" gorm:"size:50"`                   // Type of token (e.g., 'Bearer')
	ExpiresIn   int            `json:"expires_in"`                                  // Token expiration time in seconds
	CreatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"` // Timestamp when the access token was created
	UpdatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"` // Timestamp when the access token was last updated
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`                     // Timestamp for soft deletes
}

var TNOAuthAccessToken = "oauth_access_tokens"

// TableName sets the insert table name for OAuthAccessToken struct
func (OAuthAccessToken) TableName() string {
	return TNOAuthAccessToken
}

func (o *OAuthAccessToken) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
		return err
	}
	return nil
}

type OAuthRefreshToken struct {
	domain.BaseModel
	UserID       uuid.UUID      `json:"user_id" gorm:"not null"`                     // References the User table to link the token to a specific user
	ProviderID   uuid.UUID      `json:"provider_id" gorm:"not null"`                 // References the OAuthProvider table to identify the provider of the token
	RefreshToken string         `json:"refresh_token" gorm:"size:255;not null"`      // The OAuth refresh token
	CreatedAt    time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"` // Timestamp when the refresh token was created
	UpdatedAt    time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"` // Timestamp when the refresh token was last updated
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`                     // Timestamp for soft deletes
}

var TNOAuthRefreshToken = "oauth_refresh_tokens"

// TableName sets the insert table name for OAuthRefreshToken struct
func (OAuthRefreshToken) TableName() string {
	return TNOAuthRefreshToken
}

func (o *OAuthRefreshToken) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
		return err
	}
	return nil
}
