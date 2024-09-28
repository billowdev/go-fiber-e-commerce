package domain

import (
	"time"

	"github.com/billowdev/go-fiber-e-commerce/internal/core/domain"
	"github.com/billowdev/go-fiber-e-commerce/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductImage struct {
	domain.BaseModel
	ProductID uuid.UUID      `json:"product_id" gorm:"not null"`                  // References the Product table to link the image to a specific product
	ImageURL  string         `json:"image_url" gorm:"size:255;not null"`          // URL or path to the image file
	IsPrimary bool           `json:"is_primary" gorm:"default:false"`             // Indicates whether this image is the primary image for the product
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"` // Timestamp when the image record was created
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"` // Timestamp when the image record was last updated
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`                     // Timestamp for soft deletes
}

var TNProductImage = "product_images"

// TableName sets the insert table name for ProductImage struct
func (ProductImage) TableName() string {
	return TNProductImage
}

func (pi *ProductImage) BeforeCreate(tx *gorm.DB) (err error) {
	if pi.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
		return err
	}
	return nil
}
