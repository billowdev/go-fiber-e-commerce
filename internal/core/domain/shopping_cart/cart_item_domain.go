package domain

import (
	"time"

	"github.com/billowdev/go-fiber-e-commerce/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartItem struct {
	ID              uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();" json:"id"` // Unique identifier for each cart item
	CartID          uuid.UUID      `json:"cart_id" gorm:"not null"`                         // References the Cart table to link the item to a specific cart
	ProductID       uuid.UUID      `json:"product_id" gorm:"not null"`                      // References the Product table to link the item to a specific product
	VariantID       uuid.UUID      `json:"variant_id"`                                      // References the ProductVariant table for product variations
	Quantity        int            `json:"quantity" gorm:"not null"`                        // Quantity of the product added to the cart
	UnitPrice       float64        `json:"unit_price" gorm:"not null"`                      // Price per unit of the product at the time of addition to the cart
	DiscountApplied float64        `json:"discount_applied"`                                // Discount amount applied to this item, if any
	CreatedAt       time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`     // Timestamp when the cart item was created
	UpdatedAt       time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`     // Timestamp when the cart item was last updated
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`                         // Timestamp for soft deletes
}

var TNCartItem = "cart_items"

// TableName sets the insert table name for CartItem struct
func (CartItem) TableName() string {
	return TNCartItem
}

func (o *CartItem) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
		return err
	}
	return nil
}
