package database

import (
	"fmt"
	"github.com/tjaszai/go-ms-gateway/internal/config"
	"github.com/tjaszai/go-ms-gateway/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var entityList = []interface{}{
	&entity.Microservice{},
}

type DBManager struct {
	DB *gorm.DB
}

var Manager *DBManager

func (manager *DBManager) GetDB() *gorm.DB {
	return manager.DB
}

func (manager *DBManager) CheckConnection() error {
	sqlDB, err := manager.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get database connection: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("database is not reachable: %w", err)
	}

	return nil
}

func InitConnection(maxRetries int, delay time.Duration) {
	dsn := config.Config("DATABASE_DSN")
	if dsn == "" {
		log.Fatalf("Database DSN is required")
	}

	var db *gorm.DB
	var err error

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

		if err == nil {
			err := db.AutoMigrate(entityList...)
			if err != nil {
				log.Fatalf("Database migration failed: %s", err)
			}

			fmt.Println("Connected to database")
			Manager = &DBManager{DB: db}
			break
		}

		log.Printf("Database connection failed: %v. Retrying in %v...\n", err, delay)
		time.Sleep(delay)
	}
}
