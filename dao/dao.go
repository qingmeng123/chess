package dao

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var (
	name   = "root"
	pwd    = "@XUEHUI."
	host   = "localhost"
	port   = "3306"
	dbname = "player"
)

func InitDB() {
	dsn := name + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	//gormDB连接mysql
	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		fmt.Println("mysql conn err:", err)
		return
	}
	DB = gormDB
}
