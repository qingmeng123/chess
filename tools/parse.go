package tools

import (
	"Chess/model"
	"github.com/gin-gonic/gin"
	"log"
)

//绑定用户
func ParseUser(c *gin.Context) *model.User {
	m := new(model.User)
	if err := c.Bind(m); err != nil {
		log.Println(err.Error())
	}
	m.LastIp = c.ClientIP()
	return m
}
func Success(c *gin.Context, code int, msg string, data interface{}) {
	m := model.Result{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	c.JSON(200, m)
}
func Error(c *gin.Context, code int, msg string, data interface{}) {
	m := model.Result{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	c.JSON(200, m)
}
