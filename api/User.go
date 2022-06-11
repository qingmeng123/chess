package api

import (
	"Chess/service"
	"Chess/tools"
	"fmt"
	"github.com/gin-gonic/gin"
)

var userService *service.UserService

func UserLogin(c *gin.Context) {
	u := tools.ParseUser(c)
	fmt.Println(u)
	loginUser, err := userService.Login(u)
	if err != nil {
		fmt.Println("err:", err.Error())
		tools.Error(c, -1, err.Error(), "")
	} else {
		tools.Success(c, 0, "登录成功", loginUser)
	}
}

func UserRegister(c *gin.Context) {
	u := tools.ParseUser(c)
	u, err := userService.Register(u)
	if err != nil {
		fmt.Println("err:", err.Error())
		tools.Error(c, -1, err.Error(), "")
	} else {
		tools.Success(c, 0, "注册成功", u)
	}
}

func init() {
	userService = new(service.UserService)
}
