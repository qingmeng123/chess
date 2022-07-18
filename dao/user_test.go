/*******
* @Author:qingmeng
* @Description:
* @File:user_test
* @Date:2022/7/18
 */

package dao

import (
	"Chess/model"
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
)

var mock sqlmock.Sqlmock
var gormDB *gorm.DB

func init() {
	var err error
	var db *sql.DB
	db, mock, err = sqlmock.New()
	if err != nil {
		log.Fatalln("into sqlmock(mysql) db err:", err)
	}
	gormDB, err = gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("init DB with sqlmock(gorm) fail err:", err)
	}
	DB = gormDB
}

func TestUserDao_GetUserByName(t *testing.T) {
	user_name := "xy"
	// 新建字段
	rows := sqlmock.NewRows([]string{"id", "user_name", "pass_word", "phone"}).
		AddRow("1", "xy", "1234567", "13611111111")//"https",
		//"",
		//time.Now(),
		//nil,
		//time.Now().Add(10*time.Second),

	mock.ExpectQuery("^SELECT \\* FROM `user` WHERE user_name=\\? ORDER BY `user`.`id` LIMIT 1").
		WithArgs(user_name).WillReturnRows(rows)
	ud := UserDao{}
	res, err := ud.GetUserByName(user_name)
	if err != nil {
		t.Fatal("get user info err:", err)
	} else {
		fmt.Println(res)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}

func TestUserDao_UpdateUser(t *testing.T) {
	user := model.User{
		Id:       1,
		UserName: "xy",
		PassWord: "1234567",
		Token:    "bbbbbb",
	}

	mock.ExpectBegin()
	//错误示例
	//mock.ExpectExec("UPDATE `user` SET `token`=\\? WHERE `id`=\\?").
	//		WithArgs(user.Token,user.Id).WillReturnResult(sqlmock.NewResult(0,1))

	mock.ExpectExec("UPDATE `user`").
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectCommit()
	ud := UserDao{}
	err := ud.UpdateUser(&user)
	if err != nil {
		t.Fatal("update user info err:", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}

func TestUserDao_InsertUser(t *testing.T) {
	user := model.User{
		Id:       1,
		UserName: "xy",
		PassWord: "1234567",
		Token:    "bbbbbb",
	}

	mock.ExpectBegin()

	//错误示例
	//mock.ExpectExec("INSERT INTO `user`" +
	//		"(`user_name`,`pass_word`,`salt`,`create_time`,`update_time`,`last_ip`,`token`,`status`)" ).
	mock.ExpectExec("INSERT INTO `user`").
		WithArgs(user.UserName, user.PassWord, user.Salt, user.CreateTime, user.UpdateTime, user.LastIp, user.Token, user.Status).
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectCommit()
	ud := UserDao{}
	err := ud.InsertUser(&user)
	if err != nil {
		t.Fatal("insert user info err:", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}

func TestUserDao_GetUserByToken(t *testing.T) {
	token := "aaaaaa"
	// 新建字段
	rows := sqlmock.NewRows([]string{"id", "user_name", "pass_word", "token"}).
		AddRow("1", "xy", "1234567", "aaaaaa")//"https",
		//"",
		//time.Now(),
		//nil,
		//time.Now().Add(10*time.Second),

	mock.ExpectQuery("^SELECT \\* FROM `user` WHERE token=\\? ORDER BY `user`.`id` LIMIT 1").
		WithArgs(token).WillReturnRows(rows)
	ud := UserDao{}
	res, err := ud.GetUserByToken(token)
	if err != nil {
		t.Fatal("get user info err:", err)
	} else {
		fmt.Println(res)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}
