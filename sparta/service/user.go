package service

import (
	"errors"
	"fmt"
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

func (u *UserService) AddUser(iUserAddDTO *dto.UserAddDTO) error {
	if u.Dao.CheckUserExist(iUserAddDTO.NT) {
		return errors.New(fmt.Sprintf("user %s already exist", iUserAddDTO.NT))
	}

	return u.Dao.AddUser(iUserAddDTO)
}

func (u *UserService) GetUserByID(iCommonIDDTO *dto.CommonIDDTO) (model.User, error) {
	return u.Dao.GetUserByID(iCommonIDDTO.ID)
}

func (u *UserService) GetUserList(iUserListDTO *dto.UserListDTO) ([]model.User, int64, error) {
	return u.Dao.GetUserList(iUserListDTO)
}

func (u *UserService) UpdateUser(iUserUpdateDTO *dto.UserUpdateDTO) error {
	if iUserUpdateDTO.ID == 0 {
		return errors.New("Invalid user id")
	}

	return u.Dao.UpdateUser(iUserUpdateDTO)
}

func (u *UserService) DeleteUserByID(iCommonIDDTO *dto.CommonIDDTO) error {
	return u.Dao.DeleteUserByID(iCommonIDDTO.ID)
}
