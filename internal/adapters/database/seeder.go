package database

import (
	"fmt"

	"github.com/billowdev/go-fiber-e-commerce/internal/adapters/database/seeders"
	userDomain "github.com/billowdev/go-fiber-e-commerce/internal/core/domain/core/user"
	"gorm.io/gorm"
)

func helperDeleteInfo(db *gorm.DB, table string) error {
	err := db.Exec(fmt.Sprintf("DELETE FROM %s", table)).Error
	if err != nil {
		return err
	}
	_ = db.Exec(fmt.Sprintf("SELECT setval('%s_id_seq', 1, false)", table)).Error
	return nil
}

func RunSeeds(db *gorm.DB) {
	seeders.SeedUser(db)

}

func resetSeeder(db *gorm.DB) error {
	if err := helperDeleteInfo(db, userDomain.TNUser); err != nil {
		return err
	}
	return nil
}
