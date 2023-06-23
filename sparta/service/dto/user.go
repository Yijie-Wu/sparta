package dto

type UserLoginDTO struct {
	Name     string `form:"name" json:"name" binding:"required,my_validator" message:"用户名填写错误" required_err:"用户名不能为空"`
	Password string `form:"password" json:"password" binding:"required,my_validator" message:"密码填写错误" required_err:"密码不能为空"`
}
