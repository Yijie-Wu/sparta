package dao

import (
	"gorm.io/gorm"
	"sparta/global"
)

type BaseDao struct {
	ORM *gorm.DB
}

func NewBaseDao() BaseDao {
	return BaseDao{
		ORM: global.DB,
	}
}
