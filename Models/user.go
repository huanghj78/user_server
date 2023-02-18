package Models

import (
	"errors"
	"gorm.io/gorm"
)

type UserInfo struct {
	BaseModels
	UserName 	string 		`json:"user_name"`
}

func (u *UserInfo) GetUserInfo() *gorm.DB {
	r := DBHelper.Where("user_name = ?", u.UserName).Find(&u)
	if r.RowsAffected <= 0 {
		r.Error = errors.New("user not found")
	}
	return r
}
