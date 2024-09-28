package domain

import (
	"time"

	"github.com/billowdev/go-fiber-e-commerce/internal/core/domain"
	"github.com/billowdev/go-fiber-e-commerce/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	domain.BaseModel
	TotalPrice   float64        `json:"total_price" gorm:"type:float;not null"`      // Total price of the order
	Status       string         `json:"status" gorm:"size:50;not null"`              // Current status of the order (e.g., 'pending', 'shipped', 'delivered')
	OrderDate    time.Time      `json:"order_date" gorm:"default:CURRENT_TIMESTAMP"` // Timestamp when the order was placed
	DeliveryDate *time.Time     `json:"delivery_date"`                               // Expected delivery date of the order
	CreatedBy    uuid.UUID      `json:"created_by" gorm:"not null"`                  // References the User table to track who created the order
	CreatedAt    time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"` // Timestamp when the order record was created
	UpdatedAt    time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"` // Timestamp when the order record was last updated
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`                     // Timestamp for soft deletes
}

var TNOrder = "orders"

// TableName sets the insert table name for Order struct
func (Order) TableName() string {
	return TNOrder
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
		return err
	}
	return nil
}
