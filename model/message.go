package model

import "time"

type Message struct {
	Id         int64     `json:"id"`                             //房间ID
	UserId     int64     `json:"user_id" form:"user_id"`         //用户ID
	DestId     int64     `json:"dest_id" form:"dest_id" `        //接收ID
	Cmd        int64     `json:"cmd" form:"cmd"`                 //操作
	IsRed      bool      `json:"is_red"`                         //是否红方
	Content    string    `json:"content" form:"content"`         //消息内容
	Status     int       `json:"status" form:"status"`           //状态
	CreateTime time.Time `json:"create_time" form:"create_time"` //发送时间
	Token      string    `json:"token"`                          //token
	Move       move      `json:"move"`                           //棋子移动
}

//从坐标(x1,y1)到（x2,y2)
type move struct {
	X1 int `json:"x1"`
	Y1 int `json:"y1"`
	X2 int `json:"x2"`
	Y2 int `json:"y2"`
}
