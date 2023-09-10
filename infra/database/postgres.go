package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	// DB is the connection handle
	// for the database
	DB  *gorm.DB
	err error
)

// DBConnection opens a database connection
func DBConnection(dsn string) (*gorm.DB, error) {
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	return DB, nil
}

// GetDB helps you to get a connection
func GetDB() *gorm.DB {
	return DB
}

// CloseDB closes database connection (unnecessary)
func CloseDB() error {
	dbInstance, _ := GetDB().DB()
	return dbInstance.Close()
}

// DBUseAble returns true if the database is available
func DBUseAble() bool {
	return err == nil
}
