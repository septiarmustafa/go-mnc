package database

import (
	"go-mnc/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
    db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=go sslmode=disable password=postgres")
    if err != nil {
        return nil, err
    }

    db.AutoMigrate(&models.User{})
    DB = db

    return db, nil
}
