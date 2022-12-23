package config

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnect() (*gorm.DB, error) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", Config.DBHOST, Config.DBPORT, Config.DBUSER, Config.DBNAME, Config.DBPASS)

	db, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()

	err = sqlDb.Ping()

	if err != nil {
		return nil, err
	}

	sqlDb.SetMaxIdleConns(10)

	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func CloseConnection(db *gorm.DB) error {
	sqlDB, err := db.DB()

	if err != nil {
		return err
	}

	err = sqlDB.Ping()

	if err != nil {
		return err
	}

	sqlDB.Close()

	return nil
}
