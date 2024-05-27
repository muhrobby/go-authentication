package database

import (
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/muhrobby/go-authentication/models/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {

	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/go_auth?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Info("Failed to connect Database: ", err)
	}
	db, errdb := DB.DB()
	if errdb != nil {
		log.Info("Failed to connect Database: ", err)
	}

	db.SetConnMaxIdleTime(time.Hour * 1)
	log.Info("successfully connected")

	err = DB.AutoMigrate(
		&entity.User{},
		&entity.Role{})

	if err != nil {
		log.Info("Failed to migrate")
	}
	log.Info("successfully migrated")

	return DB

}
