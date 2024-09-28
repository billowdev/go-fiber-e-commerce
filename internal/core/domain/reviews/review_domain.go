package domain

import (
	"time"

	"github.com/billowdev/go-fiber-e-commerce/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();" json:"id"` // Unique identifier for each review
	ProductID uuid.UUID      `json:"product_id" gorm:"not null"`                      // References the Product table
	UserID    uuid.UUID      `json:"user_id" gorm:"not null"`                         // References the User table
	Rating    int            `json:"rating" gorm:"not null"`                          // Rating given by the user
	Comment   string         `json:"comment"`                                         // Comment left by the user
	Status    string         `json:"status" gorm:"size:50"`                           // Status of the review
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`     // Created timestamp
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`     // Updated timestamp
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`                         // Soft delete timestamp
}

var TNReview = "reviews"

// TableName sets the insert table name for Review struct
func (Review) TableName() string {
	return TNReview
}

func (o *Review) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
		return err
	}
	return nil
}
