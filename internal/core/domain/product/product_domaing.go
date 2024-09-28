package domain

import (
	"time"

	"github.com/billowdev/go-fiber-e-commerce/internal/core/domain"
	"github.com/billowdev/go-fiber-e-commerce/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	domain.BaseModel
	Name       string         `json:"name" gorm:"size:150;not null"`               // Name of the product
	CategoryID uuid.UUID      `json:"category_id" gorm:"not null"`                 // References the Category table to classify the product
	CreatedBy  uuid.UUID      `json:"created_by" gorm:"not null"`                  // References the User table to track who created the product
	CreatedAt  time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"` // Timestamp when the product was created
	UpdatedAt  time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"` // Timestamp when the product was last updated
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`                     // Timestamp for soft deletes
}

var TNProduct = "products"

// TableName sets the insert table name for Product struct
func (Product) TableName() string {
	return TNProduct
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
		return err
	}
	return nil
}
