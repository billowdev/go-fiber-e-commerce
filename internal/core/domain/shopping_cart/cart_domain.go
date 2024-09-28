package domain

import (
	"time"

	"github.com/billowdev/go-fiber-e-commerce/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Cart represents a user's shopping cart.
type Cart struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();" json:"id"` // Unique identifier for each cart
	UserID    uuid.UUID      `json:"user_id" gorm:"not null"`                         // References the User table to link the cart to a specific user
	Status    string         `json:"status" gorm:"size:50;not null"`                  // Status of the cart (e.g., 'active', 'abandoned', 'completed')
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`     // Timestamp when the cart was created
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`     // Timestamp when the cart was last updated
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`                         // Timestamp for soft deletes
}

var TNCart = "carts"

// TableName sets the insert table name for Cart struct
func (Cart) TableName() string {
	return TNCart
}

func (o *Cart) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
		return err
	}
	return nil
}
