package service

import (
	"Chess/dao"
	"Chess/global"
	"Chess/model"
	"Chess/tools"
	"errors"
	"time"
)

type UserService struct {
}

var userDao *dao.UserDao

//用户登录
func (s *UserService) Login(user *model.User) (*model.User, error) {
	var err error
	tmp := new(model.User)
	tmp, err = userDao.GetUserByName(user.UserName)
	if err != nil {
		return nil, err
	}
	if !tools.ValidatePwd(user.PassWord, tmp.Salt, tmp.PassWord) {
		return nil, errors.New("账号或密码不正确")
	}
	//清除临时信息
	tmp.Token = tools.GetUUID()
	tmp.UpdateTime = time.Now()
	err = userDao.UpdateUser(tmp)
	if err != nil {
		return nil, err
	}
	//更新Token
	global.Lock.Lock()
	delete(global.AllUserByToken, tmp.Token)
	global.AllUserByToken[tmp.Token] = tmp

	delete(global.AllUserById, tmp.Id)
	global.AllUserById[tmp.Id] = tmp
	global.Lock.Unlock()
	return tmp, nil
}

//用户注册
func (s *UserService) Register(user *model.User) (*model.User, error) {
	u, _ := userDao.GetUserByName(user.UserName)
	if u.Id > 0 {
		return nil, errors.New("该账号已经注册")
	}

	user.Salt = tools.GetUUID()
	user.PassWord = tools.MakeDbPwd(user.PassWord, user.Salt)
	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()
	user.Token = tools.GetUUID()
	user.Status = "0"
	err := userDao.InsertUser(user)
	if err != nil {
		return nil, err
	}
	//更新缓存
	global.Lock.Lock()
	delete(global.AllUserByToken, user.Token)
	global.AllUserByToken[user.Token] = user

	delete(global.AllUserById, user.Id)
	global.AllUserById[user.Id] = user
	global.Lock.Unlock()

	return user, nil
}

// GetIdByToken 根据token获取用户id并判断是否有效
func (s *UserService) GetIdByToken(token string) (int64, error) {
	id := s.GetIdByTokenCache(token)
	if id != 0 {
		return id, nil
	}
	u, err := userDao.GetUserByToken(token)
	if err != nil {
		return 0, err
	}
	if u.Id > 0 {
		return u.Id, nil
	}
	return 0, errors.New("没有找到用户")
}

//根据token获取临时信息
func (s *UserService) GetIdByTokenCache(token string) int64 {
	user, ok := global.AllUserByToken[token]
	if ok {
		return user.Id
	}
	return 0
}
