package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sparta/database"
)

var (
	Logger      *zap.SugaredLogger
	DB          *gorm.DB
	RedisClient *database.RedisClient
)
