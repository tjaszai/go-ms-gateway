package db

import (
	"fmt"
	"github.com/tjaszai/go-ms-gateway/config"
	"github.com/tjaszai/go-ms-gateway/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var modelList = []interface{}{
	&model.Microservice{},
	&model.MicroserviceVersion{},
	&model.User{},
}

type DatabaseManager struct {
	DB *gorm.DB
}

func NewDatabaseManager() *DatabaseManager {
	dsn := config.Config("DATABASE_DSN")
	if dsn == "" {
		log.Println("Warning: Missing DATABASE_DSN environment variable")
		panic("Database DSN is required")
	}
	var db *gorm.DB
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Println(err)
		panic("Failed to connect to database")
	}
	err = db.AutoMigrate(modelList...)
	if err != nil {
		log.Printf("Failed to migrate model list: %v\n", err)
		panic("Database migration failed")
	}
	fmt.Println("Connected to database")
	return &DatabaseManager{DB: db}
}

func (m *DatabaseManager) GetDB() *gorm.DB {
	return m.DB
}

func (m *DatabaseManager) CheckConnection() error {
	sqlDB, err := m.DB.DB()
	if err != nil {
		log.Printf("Failed to connect to database: %v\n", err)
		return fmt.Errorf("failed to get database connection: %w", err)
	}
	if err = sqlDB.Ping(); err != nil {
		log.Printf("Failed to ping database: %v\n", err)
		return fmt.Errorf("database is not reachable: %w", err)
	}
	return nil
}
