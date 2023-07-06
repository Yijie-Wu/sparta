package dao

import (
	"sparta/model"
	"sparta/service/dto"
)

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

func (d *UserDao) GetUserByName(username string) (model.User, error) {
	var iUser model.User
	err := d.ORM.Model(&iUser).Where("name=?", username).First(&iUser).Error
	return iUser, err
}

func (d *UserDao) AddUser(iUserAddDTO *dto.UserAddDTO) error {
	var iUser model.User
	iUserAddDTO.ConvertToModel(&iUser)

	err := d.ORM.Save(&iUser).Error
	if err == nil {
		iUserAddDTO.ID = iUser.ID
		iUserAddDTO.Password = ""
	}
	return err
}

func (d *UserDao) UpdateUser(iUserUpdateDTO *dto.UserUpdateDTO) error {
	var iUser model.User
	d.ORM.First(&iUser, iUserUpdateDTO.ID)
	iUserUpdateDTO.ConvertToModel(&iUser)

	return d.ORM.Save(&iUser).Error
}

func (d *UserDao) CheckUserExist(NT string) bool {
	var nTotal int64
	d.ORM.Model(&model.User{}).Where("nt = ?", NT).Count(&nTotal)

	return nTotal > 0
}

func (d *UserDao) GetUserByID(id uint) (model.User, error) {
	var iUser model.User

	err := d.ORM.First(&iUser, id).Error
	return iUser, err
}

func (d *UserDao) GetUserList(iUserListDTO *dto.UserListDTO) ([]model.User, int64, error) {
	var iUserList []model.User
	var nTotal int64

	err := d.ORM.Model(&model.User{}).
		Scopes(Paginate(iUserListDTO.Paginate)).
		Find(&iUserList).Offset(-1).Limit(-1).
		Count(&nTotal).Error

	return iUserList, nTotal, err
}

func (d *UserDao) DeleteUserByID(id uint) error {
	return d.ORM.Delete(&model.User{}, id).Error
}
