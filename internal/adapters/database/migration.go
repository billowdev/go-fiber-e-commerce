package database

import (
	userDomain "github.com/billowdev/go-fiber-e-commerce/internal/core/domain/core/user"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error
		if err != nil {
			return err
		}

		err = tx.AutoMigrate(
			&userDomain.User{},
		)
		if err != nil {
			return err
		}

		return err
	})

	return err
}
