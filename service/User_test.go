/*******
* @Author:qingmeng
* @Description:
* @File:User_test
* @Date:2022/7/17
 */

package service

import (
	"Chess/dao"
	"Chess/model"
	"log"
	"testing"
)

var tokenTest = []struct {
	token string
	out   int
}{
	{"a9b4bdb0f71f4497a9538f902149dadf", 10},
	{"b11367f5109248f4b5d1c3bb207cb15c", 11},
	{"e99a8e3bad1347caa44d3638d505e06d", 12},
}

func TestUserService_GetIdByToken(t *testing.T) {
	userService := UserService{}
	dao.InitDB()
	for _, test := range tokenTest {
		get, err := userService.GetIdByToken(test.token)
		if err != nil {
			log.Println("err:", err)
		}
		if int(get) != test.out {
			t.Errorf("expected:%d,got:%d", test.out, get)
		}
	}
}

var usersTest = []struct {
	UserName string
	PassWord string
	out      int
}{
	{UserName: "a6", PassWord: "111111", out: 28},
	{UserName: "a7", PassWord: "111111", out: 29},
	{UserName: "a8", PassWord: "111111", out: 30},
	{UserName: "a9", PassWord: "111111", out: 31},
	{UserName: "a10", PassWord: "111111", out: 32},
}

func TestUserService_Register(t *testing.T) {
	userService := UserService{}
	dao.InitDB()

	for _, test := range usersTest {
		user := model.User{UserName: test.UserName, PassWord: test.PassWord}
		res, err := userService.Register(&user)
		if err != nil {
			log.Println("err:", err)
		}
		if int(res.Id) != test.out {
			t.Errorf("expected:%d,got:%d", test.out, res.Id)
		}
	}

}

func TestUserService_Login(t *testing.T) {
	userService := UserService{}
	dao.InitDB()

	for _, test := range usersTest {
		user := model.User{UserName: test.UserName, PassWord: test.PassWord}
		res, err := userService.Login(&user)
		if err != nil {
			log.Println("err:", err)
		}
		if int(res.Id) != test.out {
			t.Errorf("expected:%d,got:%d", test.out, res.Id)
		}
	}

}
