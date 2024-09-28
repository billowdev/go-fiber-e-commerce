package domain

import (
	"time"

	"github.com/billowdev/go-fiber-e-commerce/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ShippingInfo contains shipping details for orders.
type ShippingInfo struct {
	ID                    uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();" json:"id"` // Unique identifier for each shipping record
	OrderID               uuid.UUID      `json:"order_id" gorm:"not null"`                        // References the Order table to link the shipping info to a specific order
	Address               string         `json:"address" gorm:"size:255;not null"`                // Address where the order will be shipped
	Method                string         `json:"method" gorm:"size:50;not null"`                  // Method of shipping (e.g., 'Standard', 'Express')
	ShippingCost          float64        `json:"shipping_cost" gorm:"not null"`                   // Cost of shipping the order
	TrackingNumber        string         `json:"tracking_number"`                                 // Tracking number for the shipment
	ShippedAt             time.Time      `json:"shipped_at"`                                      // Timestamp when the order was shipped
	DeliveredAt           time.Time      `json:"delivered_at"`                                    // Timestamp when the order was delivered
	EstimatedDeliveryDate time.Time      `json:"estimated_delivery_date"`                         // Estimated delivery date for the order
	CreatedAt             time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`     // Timestamp when the shipping information record was created
	UpdatedAt             time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`     // Timestamp when the shipping information record was last updated
	DeletedAt             gorm.DeletedAt `gorm:"index" json:"deleted_at"`                         // Timestamp for soft deletes
}

var TNShippingInfo = "shipping_infos"

// TableName sets the insert table name for ShippingInfo struct
func (ShippingInfo) TableName() string {
	return TNShippingInfo
}

func (o *ShippingInfo) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
		return err
	}
	return nil
}
