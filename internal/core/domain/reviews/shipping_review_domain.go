package domain

import (
	"time"

	"github.com/billowdev/go-fiber-e-commerce/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShippingReview struct {
	ID           uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();" json:"id"` // Unique identifier for each shipping review
	OrderID      uuid.UUID      `json:"order_id" gorm:"not null"`                        // References the Order table
	UserID       uuid.UUID      `json:"user_id" gorm:"not null"`                         // References the User table
	Rating       int            `json:"rating" gorm:"not null"`                          // Rating for the shipping experience
	Comment      string         `json:"comment"`                                         // Comment about the shipping experience
	DeliveryTime time.Time      `json:"delivery_time"`                                   // Timestamp when the product was delivered
	Status       string         `json:"status" gorm:"size:50"`                           // Status of the shipping process
	CreatedAt    time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`     // Created timestamp
	UpdatedAt    time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`     // Updated timestamp
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`                         // Soft delete timestamp
}

var TNShippingReview = "shipping_reviews"

// TableName sets the insert table name for ShippingReview struct
func (ShippingReview) TableName() string {
	return TNShippingReview
}

func (o *ShippingReview) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
		return err
	}
	return nil
}
