package global

import (
	"Chess/model"
	"sync"
)

var (
	// ClientMap 所有用户信息
	ClientMap map[int64]*model.Node
	Lock sync.RWMutex
	// AllUserByToken 临时存放所有用户信息
	AllUserByToken map[string]*model.User
	AllUserById    map[int64]*model.User

	// ClientMatch 存放匹配的用户
	ClientMatch []string
	// ClientStart 存放开始游戏的用户
	ClientStart map[int64]int64
)

func init() {
	ClientMap = make(map[int64]*model.Node)
	AllUserByToken = make(map[string]*model.User)
	AllUserById = make(map[int64]*model.User)
	ClientStart = make(map[int64]int64)
}
