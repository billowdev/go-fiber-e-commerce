package database

import (
	"github.com/billowdev/go-fiber-e-commerce/pkg/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase() (*gorm.DB, error) {

	dsn := configs.DB_URL
	loggerDBLevel := logger.Silent
	if configs.APP_DEBUG_MODE {
		loggerDBLevel = logger.Info
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		// PreferSimpleProtocol: DB_DRY_RUN,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(loggerDBLevel), // or logger.Silent if you don't want logs
		// Logger: logger.Default.LogMode(logger.Info), // or logger.Silent if you don't want logs
	})

	if err != nil {
		return nil, err // instead of panic, return the error
	}
	// !: ENABLE MIGRATIONS DB
	if configs.DB_RUN_MIGRATION {
		if err := RunMigrations(db); err != nil {
			return nil, err
		}
	}

	if configs.DB_RUN_SEEDER {
		// Reset new seed for testing purposes
		if err := resetSeeder(db); err != nil {
			return nil, err
		}
		RunSeeds(db)
	}
	return db, nil
}
