package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"pendaftaran-sidang/internal/model/entity"
	"time"
)

var DB *gorm.DB

type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

var config = Config{}

func OpenConnection() *gorm.DB {
	config.Read()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Sidang{})

	sqlDB, err := db.DB()
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	return db
}

func (c *Config) Read() {
	config.Host = os.Getenv("DB_HOST")
	config.User = os.Getenv("DB_USER")
	config.Password = os.Getenv("DB_PASSWORD")
	config.DBName = os.Getenv("DB_NAME")
	config.Port = os.Getenv("DB_PORT")
}
