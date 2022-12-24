package config

import (
	"blog-chi-gorm/entity"
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

	err = db.AutoMigrate(&entity.User{}, &entity.Tag{}, &entity.SubMenu{}, &entity.Article{}, &entity.Role{}, &entity.PostTag{}, &entity.PostCategory{}, &entity.Post{}, &entity.Permission{}, &entity.Menu{}, &entity.Category{})

	if err != nil {
		return nil, err
	}

	sqlDb.SetMaxIdleConns(10)

	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)

	return db, nil
}
