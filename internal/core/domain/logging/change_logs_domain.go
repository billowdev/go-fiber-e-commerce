package domain

import (
	"time"

	"github.com/billowdev/go-fiber-e-commerce/internal/core/domain"
	"github.com/billowdev/go-fiber-e-commerce/pkg/uuidv7"
	"gorm.io/gorm"
)

// LogEntry represents a structured log entry for logging purposes.
type LogEntry struct {
	Timestamp time.Time              `json:"@timestamp"`        // Timestamp of the log entry
	Level     string                 `json:"level"`             // Severity level of the log (e.g., 'INFO', 'ERROR')
	Service   string                 `json:"service"`           // Service name where the log was generated
	Message   string                 `json:"message"`           // Log message content
	Context   map[string]interface{} `json:"context,omitempty"` // Additional contextual information
}

// ChangeLogs represents a record of changes made to various entities within the system.
type ChangeLogs struct {
	domain.BaseModel
	EntityType  string    `json:"entity_type" gorm:"size:255;not null"`       // Type of entity changed (e.g., 'User', 'Order', etc.)
	EntityID    uint      `json:"entity_id" gorm:"not null"`                  // ID of the modified record
	FieldName   string    `json:"field_name" gorm:"size:255;not null"`        // Name of the changed field
	OldValue    string    `json:"old_value" gorm:"type:text"`                 // Previous value of the field
	NewValue    string    `json:"new_value" gorm:"type:text"`                 // New value of the field
	Action      string    `json:"action" gorm:"size:50;not null"`             // Action performed (e.g., 'UPDATE', 'INSERT', 'DELETE')
	Timestamp   time.Time `json:"timestamp" gorm:"default:CURRENT_TIMESTAMP"` // When the change was recorded
	UserID      uint      `json:"user_id" gorm:"not null"`                    // User responsible for the change
	Description string    `json:"description" gorm:"type:text"`               // Optional context about the change
}

var TNChangeLogs = "change_logs"

// TableName sets the insert table name for ChangeLogs struct
func (ChangeLogs) TableName() string {
	return TNChangeLogs
}

// BeforeCreate generates a unique ID before a new ChangeLogs entry is created.
func (cl *ChangeLogs) BeforeCreate(tx *gorm.DB) (err error) {
	if cl.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
		return err
	}
	return nil
}

// NewChangeLog creates a new ChangeLogs instance with the current timestamp.
func NewChangeLog(entityType string, entityID uint, fieldName, oldValue, newValue, action string, userID uint, description string) *ChangeLogs {
	return &ChangeLogs{
		EntityType:  entityType,
		EntityID:    entityID,
		FieldName:   fieldName,
		OldValue:    oldValue,
		NewValue:    newValue,
		Action:      action,
		Timestamp:   time.Now(),
		UserID:      userID,
		Description: description,
	}
}
