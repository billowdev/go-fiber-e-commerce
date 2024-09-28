package domain

import (
	"time"

	"github.com/billowdev/go-fiber-e-commerce/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Coupon struct {
	ID            uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();" json:"id"` // Unique identifier for each coupon
	Code          string         `json:"code" gorm:"size:50;not null"`                    // Unique code for the coupon
	Description   string         `json:"description"`                                     // Description about the coupon
	DiscountType  string         `json:"discount_type" gorm:"size:50;not null"`           // Type of discount
	DiscountValue float64        `json:"discount_value" gorm:"not null"`                  // Value of the discount
	StartDate     time.Time      `json:"start_date" gorm:"not null"`                      // When the coupon becomes valid
	EndDate       time.Time      `json:"end_date" gorm:"not null"`                        // When the coupon expires
	CreatedBy     uuid.UUID      `json:"created_by" gorm:"not null"`                      // References the User table
	CreatedAt     time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`     // Created timestamp
	UpdatedAt     time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`     // Updated timestamp
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`                         // Soft delete timestamp
}

var TNCoupon = "coupons"

// TableName sets the insert table name for Coupon struct
func (Coupon) TableName() string {
	return TNCoupon
}

func (o *Coupon) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
		return err
	}
	return nil
}
