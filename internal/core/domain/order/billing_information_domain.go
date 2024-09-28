package domain

import (
	"time"

	"github.com/billowdev/go-fiber-e-commerce/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BillingInfo contains billing details associated with user orders.
type BillingInfo struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();" json:"id"` // Unique identifier for each billing information record
	UserID    uuid.UUID      `json:"user_id" gorm:"not null"`                         // References the User table to link the billing info to a specific user
	OrderID   uuid.UUID      `json:"order_id" gorm:"not null"`                        // References the Order table to link the billing info to a specific order
	Address   string         `json:"address" gorm:"size:255;not null"`                // Billing address for the order
	Phone     string         `json:"phone" gorm:"size:10;not null"`                   // Phone number associated with the billing information
	Email     string         `json:"email" gorm:"size:100;not null"`                  // Email address associated with the billing information
	Method    string         `json:"method" gorm:"size:50;not null"`                  // Method used for billing (e.g., 'Credit Card', 'Bank Transfer')
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`     // Timestamp when the billing information record was created
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`     // Timestamp when the billing information record was last updated
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`                         // Timestamp for soft deletes
}

var TNBillingInfo = "billing_infos"

// TableName sets the insert table name for BillingInfo struct
func (BillingInfo) TableName() string {
	return TNBillingInfo
}

func (o *BillingInfo) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
		return err
	}
	return nil
}
