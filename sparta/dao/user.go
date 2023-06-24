package dao

import "sparta/model"

var userDao *UserDao

type UserDao struct {
	BaseDao
}

func NewUserDao() *UserDao {
	if userDao == nil {
		userDao = &UserDao{
			NewBaseDao(),
		}
	}

	return userDao
}

func (d *UserDao) GetUserByNameAndPassword(username, password string) model.User {
	var iUser model.User
	d.ORM.Model(&iUser).Where("name=? and password=?", username, password).Find(&iUser)
	return iUser
}
