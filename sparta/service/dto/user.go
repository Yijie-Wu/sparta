package dto

import "sparta/model"

type UserLoginDTO struct {
	Name     string `form:"name" json:"name" binding:"required" message:"用户名填写错误" required_err:"用户名不能为空"`
	Password string `form:"password" json:"password" binding:"required" message:"密码填写错误" required_err:"密码不能为空"`
}

type UserAddDTO struct {
	ID       uint
	NT       string `form:"nt" json:"nt" binding:"required" message:"NT填写错误" required_err:"NT不能为空"`
	Name     string `form:"name" json:"name" binding:"required" message:"用户名填写错误" required_err:"用户名不能为空"`
	Email    string `form:"email" json:"email" binding:"required" message:"邮箱填写错误" required_err:"邮箱不能为空"`
	Avatar   string `form:"avatar" json:"avatar"`
	Password string `form:"password" json:"password,omitempty" binding:"required" message:"密码填写错误" required_err:"密码不能为空"`
}

type UserUpdateDTO struct {
	ID    uint   `json:"id" form:"id" uri:"id"`
	NT    string `form:"nt" json:"nt"`
	Name  string `form:"name" json:"name"`
	Email string `form:"email" json:"email"`
}

type UserListDTO struct {
	Paginate
}

func (d *UserAddDTO) ConvertToModel(iUser *model.User) {
	iUser.NT = d.NT
	iUser.Name = d.Name
	iUser.Email = d.Email
	iUser.Avatar = d.Avatar
	iUser.Password = d.Password
}

func (d *UserUpdateDTO) ConvertToModel(iUser *model.User) {
	iUser.ID = d.ID
	iUser.NT = d.NT
	iUser.Name = d.Name
	iUser.Email = d.Email
}
