package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alifnuryana/go-auth-jwt/helpers"
	"github.com/alifnuryana/go-auth-jwt/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func SetupDatabase() {
	loggerConfig := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", helpers.Load("DB_USER"), helpers.Load("DB_PASS"), helpers.Load("DB_HOST"), helpers.Load("DB_PORT"), helpers.Load("DB_NAME"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: loggerConfig,
	})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{}, &models.Todo{})

	DB = db
}
