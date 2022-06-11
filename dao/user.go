package dao

import "Chess/model"

type UserDao struct {
}

func (d *UserDao) GetUserByName(userName string) (user *model.User, err error) {
	result := DB.Where("user_name=?", userName).First(&user)
	err = result.Error
	return
}

func (d *UserDao) UpdateUser(user *model.User) error {
	result := DB.Model(&user).Update("token", user.Token)
	return result.Error
}

func (d *UserDao) InsertUser(user *model.User) error {
	result := DB.Select("user_name", "pass_word", "salt", "create_time", "update_time", "last_ip", "token", "status").Create(&user)
	return result.Error
}

func (d *UserDao) GetUserByToken(token string) (user *model.User, err error) {
	result := DB.Where("token=?", token).First(&user)
	err = result.Error
	return
}
