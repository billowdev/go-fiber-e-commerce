package domain

import (
	"github.com/billowdev/go-fiber-e-commerce/internal/core/domain"
	"github.com/billowdev/go-fiber-e-commerce/pkg/uuidv7"
	"gorm.io/gorm"
)

type UserRole struct {
	domain.BaseModel
	RoleName    string `json:"role_name" gorm:"size:50;not null"` // Name of the role (e.g., 'ADMIN', 'SELLER', 'CUSTOMER')
	Description string `json:"description" gorm:"size:255"`       // Description of the role and its purpose within the system
}

var TNUserRole = "user_roles"

// TableName sets the insert table name for UserRole struct
func (UserRole) TableName() string {
	return TNUserRole
}
func (u *UserRole) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
		return err
	}
	return nil
}
