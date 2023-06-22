package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"sparta/model"
	"time"
)

func InitDB() (*gorm.DB, error) {
	logMode := logger.Info
	if !viper.GetBool("mode.development") {
		logMode = logger.Error
	}

	name := viper.GetString("mysql.name")
	port := viper.GetString("mysql.port")
	server := viper.GetString("mysql.server")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, server, port, name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "oss_",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logMode),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(viper.GetInt("mysql.MaxIdleConn"))
	sqlDB.SetMaxOpenConns(viper.GetInt("mysql.MaxOpenConn"))
	sqlDB.SetConnMaxLifetime(time.Hour)

	// init tables
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
