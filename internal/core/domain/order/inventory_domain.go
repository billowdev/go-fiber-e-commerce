package domain

import (
	"time"

	"github.com/billowdev/go-fiber-e-commerce/internal/core/domain"
	"github.com/billowdev/go-fiber-e-commerce/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Inventory struct {
	domain.BaseModel
	ProductID uuid.UUID      `json:"product_id" gorm:"not null"`                  // References the Product table to link the inventory record to a specific product
	Quantity  int            `json:"quantity" gorm:"not null"`                    // Quantity of the product available in inventory
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"` // Timestamp when the inventory item record was created
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"` // Timestamp when the inventory record was last updated
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`                     // Timestamp for soft deletes
}

var TNInventory = "inventories"

// TableName sets the insert table name for Inventory struct
func (Inventory) TableName() string {
	return TNInventory
}

func (inv *Inventory) BeforeCreate(tx *gorm.DB) (err error) {
	if inv.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
		return err
	}
	return nil
}
