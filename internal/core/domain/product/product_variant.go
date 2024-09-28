package domain

import (
	"time"

	"github.com/billowdev/go-fiber-e-commerce/internal/core/domain"
	"github.com/billowdev/go-fiber-e-commerce/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductVariant struct {
	domain.BaseModel
	ProductID      uuid.UUID      `json:"product_id" gorm:"not null"`                    // References the Product table to link the variant to a specific product
	VariantName    string         `json:"variant_name" gorm:"size:100;not null"`        // Name of the variant (e.g., 'Size', 'Color')
	VariantValue   string         `json:"variant_value" gorm:"size:100;not null"`       // Value of the variant (e.g., 'Large', 'Red')
	PriceAdjustment float64        `json:"price_adjustment" gorm:"type:float;default:0"` // Price adjustment based on the variant
	CreatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`   // Timestamp when the product variant record was created
	UpdatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`   // Timestamp when the product variant record was last updated
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`                       // Timestamp for soft deletes
}

var TNProductVariant = "product_variants"

// TableName sets the insert table name for ProductVariant struct
func (ProductVariant) TableName() string {
	return TNProductVariant
}

func (pv *ProductVariant) BeforeCreate(tx *gorm.DB) (err error) {
	if pv.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
		return err
	}
	return nil
}
