package domain

import (
	"time"

	"github.com/billowdev/go-fiber-e-commerce/internal/core/domain"
	"github.com/billowdev/go-fiber-e-commerce/pkg/uuidv7"
	"gorm.io/gorm"
)

type UserPermission struct {
	domain.BaseModel
	RoleID         uint      `json:"role_id" gorm:"not null"`                 // References the UserRole table to link permissions to specific roles
	Permission     string    `json:"permission" gorm:"size:100;not null"`     // Description of the permission (e.g., 'CREATE_PRODUCT', 'VIEW_ORDERS')
	PermissionType string    `json:"permission_type" gorm:"size:50;not null"` // Type of permission (e.g., 'Read', 'Write', 'Admin')
	EffectiveDate  time.Time `json:"effective_date"`                          // Timestamp when the permission became effective
	Description    string    `json:"description"`                             // Brief description of the permissions associated with this role
}

var TNUserPermission = "user_permissions"

// TableName sets the insert table name for UserPermission struct
func (UserPermission) TableName() string {
	return TNUserPermission
}

func (u *UserPermission) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
		return err
	}
	return nil
}
