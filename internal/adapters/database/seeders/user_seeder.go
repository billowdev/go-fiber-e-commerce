package seeders

import (
	"time"

	domain "github.com/billowdev/go-fiber-e-commerce/internal/core/domain"
	userDomain "github.com/billowdev/go-fiber-e-commerce/internal/core/domain/core/user"
	"github.com/billowdev/go-fiber-e-commerce/pkg/argon2id"
	"github.com/billowdev/go-fiber-e-commerce/pkg/uuidv7"
	"gorm.io/gorm"
)

func SeedPassword(password string) string {
	hash, _ := argon2id.CreateHash(password, argon2id.DefaultParams)
	return hash
}

var SEED_USER = []userDomain.User{
	{
		BaseModel: domain.BaseModel{
			ID:        uuidv7.MustParseToUUIDv7("019237ed-062f-7884-8787-89f0fa7c038f"),
			UpdatedAt: time.Now(),
		},
		FirstName:        "John",
		LastName:         "Doe",
		RoleID:           "admin",
		Address:          "123 Elm St, Springfield, IL",
		PhoneNumber:      "+1234567890",
		Email:            "john@example.com",
		Username:         "john",
		Password:         SeedPassword("@Test1234"),
		Status:           "active",
		LastLogin:        time.Now(),
		TwoFactorEnabled: false,
		TwoFactorSecret:  "",
		CreatedByID:      "",
	},
	{
		BaseModel: domain.BaseModel{
			UpdatedAt: time.Now(),
		},
		FirstName:        "Jane",
		LastName:         "Smith",
		RoleID:           "user",
		Address:          "456 Oak St, Springfield, IL",
		PhoneNumber:      "+1987654321",
		Email:            "jane@example.com",
		Username:         "jane",
		Password:         SeedPassword("@Test1234"),
		Status:           "active",
		LastLogin:        time.Now(),
		TwoFactorEnabled: true,
		TwoFactorSecret:  "abc123secret",
		CreatedByID:      "",
	},
	{
		BaseModel: domain.BaseModel{
			UpdatedAt: time.Now(),
		},
		FirstName:        "Mike",
		LastName:         "Johnson",
		RoleID:           "user",
		Address:          "789 Pine St, Springfield, IL",
		PhoneNumber:      "+1472583690",
		Email:            "mike@example.com",
		Username:         "mike",
		Password:         SeedPassword("@Test1234"),
		Status:           "inactive",
		LastLogin:        time.Now(),
		TwoFactorEnabled: false,
		TwoFactorSecret:  "",
		CreatedByID:      "",
	},
	{
		BaseModel: domain.BaseModel{
			UpdatedAt: time.Now(),
		},
		FirstName:        "Alice",
		LastName:         "Brown",
		RoleID:           "user",
		Address:          "101 Maple St, Springfield, IL",
		PhoneNumber:      "+3216549870",
		Email:            "alice@example.com",
		Username:         "alice",
		Password:         SeedPassword("@Test1234"),
		Status:           "active",
		LastLogin:        time.Now(),
		TwoFactorEnabled: true,
		TwoFactorSecret:  "def456secret",
		CreatedByID:      "",
	},
}

func SeedUser(db *gorm.DB) error {
	if err := BaseStructSeeder(db, &SEED_USER); err != nil {
		return err
	}
	return nil
}
