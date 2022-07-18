/*******
* @Author:qingmeng
* @Description:
* @File:User_test
* @Date:2022/7/18
 */

package api

import (
	"Chess/dao"
	"Chess/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserLogin(t *testing.T) {
	tests := []struct {
		name   string
		param  string
		expect string
	}{
		{"good case", `{"user_name": "xy1111","pass_word":"1234567"}`, "登录成功"},
		{"bad case", `{"user_name": "a","pass_word":"1234567"}`, "登录失败"},
	}
	r := gin.Default()
	dao.InitDB()
	r.POST("/login", UserLogin)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// httptest 这个包是为了mock一个HTTP请求
			// 参数分别是请求方法，请求URL，请求Body
			// Body只能使用io.Reader
			req := httptest.NewRequest(
				"POST",                      // 请求方法
				"/login",                    // 请求URL
				strings.NewReader(tt.param), // 请求参数
			)

			// mock一个响应记录器
			w := httptest.NewRecorder()

			// 让server端处理mock请求并记录返回的响应内容
			r.ServeHTTP(w, req)

			// 校验状态码是否符合预期
			assert.Equal(t, http.StatusOK, w.Code)

			// 解析并检验响应内容是否复合预期
			var resp model.Result
			err := json.Unmarshal([]byte(w.Body.String()), &resp)
			assert.Nil(t, err)
			assert.Equal(t, tt.expect, resp.Msg)
		})
	}
}

func TestUserRegister(t *testing.T) {
	tests := []struct {
		name   string
		param  string
		expect string
	}{
		{"good case", `{"user_name": "xy11111","pass_word":"1234567"}`, "注册成功"},
		{"bad case", `{"user_name": "xy11111","pass_word":"1234567"}`, "该账号已经注册"},
	}
	r := gin.Default()
	dao.InitDB()
	r.POST("/register", UserRegister)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// httptest 这个包是为了mock一个HTTP请求
			// 参数分别是请求方法，请求URL，请求Body
			// Body只能使用io.Reader
			req := httptest.NewRequest(
				"POST",                      // 请求方法
				"/register",                 // 请求URL
				strings.NewReader(tt.param), // 请求参数
			)

			// mock一个响应记录器
			w := httptest.NewRecorder()

			// 让server端处理mock请求并记录返回的响应内容
			r.ServeHTTP(w, req)

			// 校验状态码是否符合预期
			assert.Equal(t, http.StatusOK, w.Code)

			// 解析并检验响应内容是否复合预期
			var resp model.Result
			err := json.Unmarshal([]byte(w.Body.String()), &resp)
			assert.Nil(t, err)
			assert.Equal(t, tt.expect, resp.Msg)
		})
	}
}
