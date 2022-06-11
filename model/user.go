package model

import "time"

type User struct {
	Id         int64     `json:"id" form:"id"`
	UserName   string    `json:"user_name" form:"user_name"`
	PassWord   string    `json:"pass_word" form:"pass_word"`
	Phone      string    `json:"phone"`
	Sex        string    `json:"sex"`
	Salt       string    `json:"salt"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	QQ         string    `json:"qq"`
	LastIp     string    `json:"last_ip"`
	Token      string    `json:"token"`
	//用户状态(0-正常)
	Status string `json:"status" form:"status"`
}

func (User) TableName() string {
	return "user"
}
