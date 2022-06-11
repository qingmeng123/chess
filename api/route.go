package api

import (
	"Chess/websocket"
	"github.com/gin-gonic/gin"
)

func InitRoute(engine *gin.Engine) {
	//走棋信息
	engine.GET("/ws", websocket.WS)
	//登录
	engine.POST("/login", UserLogin)
	//注册
	engine.POST("/register", UserRegister)

}
