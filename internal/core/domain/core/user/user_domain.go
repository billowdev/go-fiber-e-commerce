package domain

import (
	"time"

	"github.com/billowdev/go-fiber-e-commerce/internal/core/domain"
	"github.com/billowdev/go-fiber-e-commerce/pkg/uuidv7"
	"gorm.io/gorm"
)

type USER_STATUS string

const (
	USER_STATUS_ACTIVE   USER_STATUS = "active"
	USER_STATUS_INACTIVE USER_STATUS = "inactive"
	USER_STATUS_BANNED   USER_STATUS = "banned"
)

type User struct {
	domain.BaseModel
	FirstName        string      `json:"first_name" validate:"required,max=255" gorm:"size:255"`
	LastName         string      `json:"last_name" validate:"required,max=255"`
	RoleID           string      `json:"role_id"`
	Address          string      `json:"address" validate:"required,max=255" gorm:"size:255"`
	PhoneNumber      string      `json:"phone_number" validate:"required,max=50" gorm:"size:50"`
	Email            string      `json:"email" validate:"required,max=150" gorm:"size:50"`
	Username         string      `json:"username" validate:"required,max=255" gorm:"size:255"`
	Password         string      `json:"password"`
	Status           USER_STATUS `json:"status" validate:"required,max=50" gorm:"size:50"`
	LastLogin        time.Time   `json:"last_login"`
	TwoFactorEnabled bool        `json:"two_factor_enabled"`
	TwoFactorSecret  string      `json:"two_factor_secret"`
	CreatedByID      string      `json:"created_by_id"`
}

var TNUser = "users"

func (st *User) TableName() string {
	return TNUser
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
		return err
	}
	return nil
}

type UpdateUserDomain struct {
	domain.BaseModel
	FirstName        string    `json:"first_name" validate:"omitempty,max=255"`
	LastName         string    `json:"last_name"`
	RoleID           string    `json:"role_id"`
	Address          string    `json:"address"`
	PhoneNumber      string    `json:"phone_number"`
	Username         string    `json:"username"`
	Password         string    `json:"password"`
	Status           string    `json:"status"`
	LastLogin        time.Time `json:"last_login"`
	TwoFactorEnabled bool      `json:"two_factor_enabled"`
	TwoFactorSecret  string    `json:"two_factor_secret"`
	Email            string    `json:"email"`
	CreatedByID      string    `json:"created_by_id"`
}
