package domain

import (
	"time"

	"github.com/billowdev/go-fiber-e-commerce/internal/core/domain"
	"github.com/billowdev/go-fiber-e-commerce/pkg/uuidv7"
	"gorm.io/gorm"
)

type Email struct {
	domain.BaseModel
	UserID         uint           `json:"user_id" gorm:"not null"`                     // References the User table to link the email to a specific user
	TemporalID     string         `json:"temporal_id" gorm:"size:255"`                 // Unique identifier from Temporal services for tracking purposes
	RecipientEmail string         `json:"recipient_email" gorm:"size:255;not null"`    // Email address of the recipient
	Subject        string         `json:"subject" gorm:"size:255;not null"`            // Subject line of the email
	Body           string         `json:"body" gorm:"type:text;not null"`              // Body content of the email
	Status         string         `json:"status" gorm:"size:50;not null"`              // Current status of the email (e.g., 'sent', 'failed')
	SentAt         time.Time      `json:"sent_at" gorm:"default:CURRENT_TIMESTAMP"`    // Timestamp when the email was sent
	FailedReason   string         `json:"failed_reason" gorm:"type:text"`              // Reason for failure if the email was not successfully sent
	CreatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"` // Timestamp when the email record was created
	UpdatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"` // Timestamp when the email record was last updated
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`                     // Timestamp for soft deletes
}

var TNEmail = "emails"

// TableName sets the insert table name for Email struct
func (Email) TableName() string {
	return TNEmail
}

func (e *Email) BeforeCreate(tx *gorm.DB) (err error) {
	if e.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
		return err
	}
	return nil
}
