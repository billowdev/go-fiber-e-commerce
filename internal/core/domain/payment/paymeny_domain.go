package domain

import (
	"time"

	"github.com/billowdev/go-fiber-e-commerce/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	ID            uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();" json:"id"` // Unique identifier for each payment record
	OrderID       uuid.UUID      `json:"order_id" gorm:"not null"`                        // References the Order table to link the payment to a specific order
	PaymentDate   time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"payment_date"`   // Timestamp when the payment was made
	PaymentMethod string         `json:"payment_method" gorm:"not null"`                  // Method used for payment (e.g., 'Credit Card', 'PayPal')
	Amount        float64        `json:"amount" gorm:"not null"`                          // Amount paid for the order
	Status        string         `json:"status" gorm:"size:50;not null"`                  // Current status of the payment (e.g., 'Pending', 'Completed', 'Failed')
	TransactionID string         `json:"transaction_id"`                                  // Transaction ID from the payment gateway
	CreatedAt     time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`     // Timestamp when the payment record was created
	UpdatedAt     time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`     // Timestamp when the payment record was last updated
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`                         // Timestamp for soft deletes
}

var TNPAYMENT = "payments"

// TableName sets the insert table name for Payment struct
func (Payment) TableName() string {
	return TNPAYMENT
}

func (o *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
		return err
	}
	return nil
}
