package service

import (
	"errors"
	"sparta/dao"
	"sparta/model"
	"sparta/service/dto"
)

var userService *UserService

type UserService struct {
	BaseService
	Dao *dao.UserDao
}

func NewUserService() *UserService {
	if userService == nil {
		userService = &UserService{
			Dao: dao.NewUserDao(),
		}
	}

	return userService
}

func (u *UserService) Login(iUserDTO dto.UserLoginDTO) (model.User, error) {
	var errResult error

	iUser := u.Dao.GetUserByNameAndPassword(iUserDTO.Name, iUserDTO.Password)
	if iUser.ID == 0 {
		errResult = errors.New("invalid name or password")
	}
	return iUser, errResult
}





