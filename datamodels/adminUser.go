package datamodels

import (
	"time"
)

type AdminUser struct {
	Id           int64  `gorm:"PRIMARY_KEY;AUTO_INCREMENT",json:"id"`
	LoginName    string `gorm:"type:varchar(30);not null",json:"login_name"`//用户名
	NickName     string `gorm:"type:varchar(30)",json:"nick_name"`//昵称
	Email        string `gorm:"type:varchar(50)",json:"email"`
	Tel          string `gorm:"type:varchar(50)",json:"tel"`
	Pwd          string `gorm:"type:char(32)",json:"pwd"`
	Avatar       string `gorm:"type:varchar(255)",json:"avatar"`
	Status       int64  `gorm:"_"`
	CityId       string `gorm:"type:varchar(20)"`
	PositionId   string `gorm:"type:varchar(150)"`
	DepartmentId int64  `gorm:"-"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
