package service

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"sparta/dao"
	"sparta/global"
	"sparta/global/constants"
	"sparta/model"
	"sparta/service/dto"
	"sparta/utils"
	"strconv"
	"time"
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

func SetLoginUserTokenToRedis(uid uint, token string) error {
	return global.RedisClient.Set(constants.LOGIN_USER_TOKEN_REDIS_KEY+strconv.Itoa(int(uid)), token, viper.GetDuration("jwt.TokenExpire")*time.Minute)
}

func (u *UserService) Login(iUserDTO dto.UserLoginDTO) (model.User, string, error) {
	var errResult error
	var token string

	iUser, err := u.Dao.GetUserByName(iUserDTO.Name)

	if err != nil || !utils.CompareHashAndPassword(iUser.Password, iUserDTO.Password) {
		errResult = errors.New("invalid username or password")
	} else {
		token, err = utils.GenerateToken(iUser.ID, iUser.Name)
		if err != nil {
			errResult = errors.New(fmt.Sprintf("generate token error:%s", err.Error()))
		}
	}

	return iUser, token, errResult
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
