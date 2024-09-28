package domain

import (
	"time"

	"github.com/billowdev/go-fiber-e-commerce/internal/core/domain"
	"github.com/billowdev/go-fiber-e-commerce/pkg/uuidv7"
	"gorm.io/gorm"
)

type OTP struct {
	domain.BaseModel
	UserID     uint           `json:"user_id" gorm:"not null"`                     // References the User table to link the OTP to a specific user
	TemporalID string         `json:"temporal_id" gorm:"size:255"`                 // Unique identifier from Temporal services for tracking purposes
	Otp        string         `json:"otp" gorm:"size:10;not null"`                 // The OTP value generated for authentication purposes
	ExpiresAt  time.Time      `json:"expires_at"`                                  // Timestamp indicating when the OTP will expire
	Used       bool           `json:"used" gorm:"default:false"`                   // Indicates whether the OTP has been used
	CreatedAt  time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"` // Timestamp when the OTP record was created
	UpdatedAt  time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"` // Timestamp when the OTP record was last updated
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`                     // Timestamp for soft deletes
}

var TNOTP = "otps"

// TableName sets the insert table name for OTP struct
func (OTP) TableName() string {
	return TNOTP
}

func (o *OTP) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
		return err
	}
	return nil
}
