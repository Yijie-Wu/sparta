package dto

import "github.com/spf13/viper"

// 通用ID对应的dto
type CommonIDDTO struct {
	ID uint `json:"id" form:"id" uri:"id"`
}

// 分页dto
type Paginate struct {
	Page  int `json:"page,omitempty" form:"page"`
	Limit int `json:"limit,omitempty" form:"limit"`
}

func (d *Paginate) GetPage() int {
	if d.Page <= 0 {
		d.Page = 1
	}
	return d.Page
}

func (d *Paginate) GetLimit() int {
	if d.Limit <= 0 {
		d.Limit = viper.GetInt("page.usersLimit")
	}
	return d.Limit
}
