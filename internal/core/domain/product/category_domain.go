package domain

import (
	"time"

	"github.com/billowdev/go-fiber-e-commerce/internal/core/domain"
	"github.com/billowdev/go-fiber-e-commerce/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	domain.BaseModel
	Name      string         `json:"name" gorm:"size:100;not null"`               // Name of the category
	CreatedBy uuid.UUID      `json:"created_by" gorm:"not null"`                  // References the User table to track who created the category
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"` // Timestamp when the category was created
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"` // Timestamp when the category was last updated
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`                     // Timestamp for soft deletes
}

var TNCategory = "categories"

// TableName sets the insert table name for Category struct
func (Category) TableName() string {
	return TNCategory
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
		return err
	}
	return nil
}
