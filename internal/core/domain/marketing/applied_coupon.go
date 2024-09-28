package domain

import (
	"time"

	"github.com/billowdev/go-fiber-e-commerce/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AppliedCoupon struct {
	ID              uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();" json:"id"` // Unique identifier for each applied coupon record
	OrderID         uuid.UUID      `json:"order_id" gorm:"not null"`                        // References the Order table
	CouponID        uuid.UUID      `json:"coupon_id" gorm:"not null"`                       // References the Coupon table
	DiscountApplied float64        `json:"discount_applied" gorm:"not null"`                // Amount of discount applied
	Status          string         `json:"status" gorm:"size:50"`                           // Status of coupon application
	CreatedAt       time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`     // Created timestamp
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`                         // Soft delete timestamp
}

var TNAppliedCoupon = "applied_coupons"

// TableName sets the insert table name for AppliedCoupon struct
func (AppliedCoupon) TableName() string {
	return TNAppliedCoupon
}

func (o *AppliedCoupon) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
		return err
	}
	return nil
}
