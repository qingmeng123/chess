package main

import (
	"Chess/api"
	"Chess/dao"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	api.InitRoute(engine)
	dao.InitDB()
	engine.Run()
}
