package domain

import (
	"time"

	"github.com/billowdev/go-fiber-e-commerce/internal/core/domain"
	"github.com/billowdev/go-fiber-e-commerce/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderItem struct {
	domain.BaseModel
	OrderID    uuid.UUID      `json:"order_id" gorm:"not null"`                    // References the Order table to link the item to a specific order
	ProductID  uuid.UUID      `json:"product_id" gorm:"not null"`                  // References the Product table to identify the product being ordered
	Quantity   int            `json:"quantity" gorm:"not null"`                    // Quantity of the product ordered
	UnitPrice  float64        `json:"unit_price" gorm:"type:float;not null"`       // Price per unit of the product at the time of purchase
	TotalPrice float64        `json:"total_price" gorm:"type:float;not null"`      // Total price for this item, calculated as qty * unit_price
	CreatedAt  time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"` // Timestamp when the order item record was created
	UpdatedAt  time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"` // Timestamp when the order item record was last updated
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`                     // Timestamp for soft deletes
}

var TNOrderItem = "order_items"

// TableName sets the insert table name for OrderItem struct
func (OrderItem) TableName() string {
	return TNOrderItem
}

func (oi *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
	if oi.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
		return err
	}
	return nil
}
