package websocket

import (
	"Chess/global"
	"Chess/model"
	"Chess/service"
	"Chess/tools"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"time"
)

var userService *service.UserService

func WS(c *gin.Context) {
	//检测连接是否合法
	token := c.Query("token")
	log.Println("收到连接")
	//检查token
	userId, err := userService.GetIdByToken(token)
	if err != nil {
		tools.Error(c, -1, err.Error(), nil)
		return
	}

	//升级协议
	upgraded := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgraded.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}

	conn.SetCloseHandler(closeConn)

	node := &model.Node{Conn: conn, Send: make(chan []byte), Heart: 1}
	global.Lock.Lock()
	_, ok := global.ClientMap[userId]
	if ok {
		delete(global.ClientMap, userId)
	}
	global.ClientMap[userId] = node
	global.Lock.Unlock()
	//发送函数
	go SendProc(node)
	//接收函数
	go RecvProc(node)
	//心跳函数
	go heartbeat(conn, userId)
	//更新棋盘
	go updateBoard(conn, userId)
	//连接成功，发送消息
	log.Println("新连接:", token)
	SendSingleMsg(userId, []byte(`{"success":true}`))

}

//更新棋盘
func updateBoard(conn *websocket.Conn, userId int64) {
	global.Lock.Lock()
	node := global.ClientMap[userId]
	global.Lock.Unlock()
	for {
		if node.Heart == 2 {
			time.Sleep(30 * time.Second)
			chessBoard := model.ChessBoard{
				Board: model.Board,
			}
			b, err := json.Marshal(chessBoard)
			err = conn.WriteMessage(websocket.TextMessage, b)
			if err != nil {
				log.Println("初始化棋盘错误")
				return
			}
		}
	}
}

// 心跳检测
func heartbeat(conn *websocket.Conn, userId int64) {
	global.Lock.Lock()
	node := global.ClientMap[userId]
	global.Lock.Unlock()
	for {
		node.Heart = 1
		time.Sleep(15 * time.Second)
		err := sendPing(conn)
		if err != nil {
			fmt.Println("send heartbeat stop")
			node.Heart = 0
			return
		}
	}
}
func sendPing(conn *websocket.Conn) error {
	str := `{"cmd":0}`
	chessBoard := model.ChessBoard{
		Board: model.Board,
	}
	b, err := json.Marshal(chessBoard)
	err = conn.WriteMessage(websocket.TextMessage, b)
	err = conn.WriteMessage(websocket.TextMessage, []byte(str))
	return err
}
func closeConn(code int, text string) error {
	log.Println("断开连接", code, text)
	fmt.Println("断开连接", code, text)
	message := websocket.FormatCloseMessage(code, text)
	log.Println(string(message))
	return nil
}

// SendProc 监听发送
func SendProc(node *model.Node) {
	for {
		select {
		case data := <-node.Send:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	}
}

// RecvProc 监听接收
func RecvProc(node *model.Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			return
		}
		//处理数据
		JsonPatch(data)
	}
}

// SendSingleMsg 发送消息
func SendSingleMsg(userId int64, msg []byte) {
	global.Lock.Lock()
	node, _ := global.ClientMap[userId]
	log.Println("开始发送", node)
	global.Lock.Unlock()
	node.Send <- msg
}

// JsonPatch 解析消息
func JsonPatch(data []byte) {
	msg := new(model.Message)
	err := json.Unmarshal(data, msg)
	log.Println("新消息:", msg)
	if err != nil {
		log.Println(err.Error())
		return
	}
	switch msg.Cmd {
	case 0:
		//心跳
		id1, err1 := userService.GetIdByToken(msg.Token)
		if err1 != nil {
			log.Println(err1)
			return
		}
		global.Lock.Lock()
		node, ok := global.ClientMap[id1]
		global.Lock.Unlock()
		if ok {
			//收到心跳
			node.Heart = 2
		}
		return
	case 1:
		//准备
		log.Println(msg.Token)
		global.Lock.Lock()
		//加入到匹配队列
		if len(global.ClientMatch) == 1 && global.ClientMatch[0] == msg.Token {
			global.Lock.Unlock()
			return
		}
		global.ClientMatch = append(global.ClientMatch, msg.Token)
		global.Lock.Unlock()
		if len(global.ClientMatch) >= 2 {
			//匹配成功,准备开始游戏
			id1, err1 := userService.GetIdByToken(global.ClientMatch[0])
			id2, err2 := userService.GetIdByToken(global.ClientMatch[1])
			if id1 == id2 {
				global.Lock.Lock()
				global.ClientMatch = global.ClientMatch[1:2]
				global.Lock.Unlock()
				return
			}
			//先判断第一位是否断开连接
			node := global.ClientMap[id1]
			err := sendPing(node.Conn)
			if err != nil {
				return
			}
			count := 0
			for {
				//判断心跳和时间
				if node.Heart == 2 || count > 50 {
					break
				}
				time.Sleep(time.Millisecond * 100)
				count++
			}
			if node.Heart != 2 {
				//说明已经断开
				log.Println("之前用户已经断开")
				global.Lock.Lock()
				global.ClientMatch = global.ClientMatch[1:2]
				global.Lock.Unlock()
				return
			}
			log.Println("匹配成功,开始游戏", id1, id2, err1)
			if err1 != nil || err2 != nil {
				//获取失败
				log.Println(err1, err2)
			} else {
				//组装消息
				str1 := `{"cmd":1,"dstId":` + strconv.FormatInt(id2, 10) + `,"isRed":true}`
				str2 := `{"cmd":1,"dstId":` + strconv.FormatInt(id1, 10) + `,"isRed":false}`
				log.Println("组装消息发送")
				log.Println(str1, str2)
				SendSingleMsg(id1, []byte(str1))
				SendSingleMsg(id2, []byte(str2))
				//清空匹配数组
				global.Lock.Lock()
				global.ClientStart[id1] = id2
				global.ClientStart[id2] = id1
				global.ClientMatch = global.ClientMatch[0:0]
				global.Lock.Unlock()
			}
		}
		return
	case 2:
		//收到移动消息
		id1, err1 := userService.GetIdByToken(msg.Token)
		if err1 != nil {
			log.Println(err1)
			return
		}
		//移动棋子
		info := model.Move(msg.IsRed, msg.Move.X1, msg.Move.Y1, msg.Move.X2, msg.Move.Y2)
		if info =="红方胜"||info=="黑方胜"{
			return
		}
		if err != nil {
			SendSingleMsg(id1, []byte(info))
		}
		//获取id1对应的id2进行转发
		id2, ok := global.ClientStart[id1]
		if !ok {
			return
		}
		//开始发送
		str1 := `{"cmd":2,"dstId":` + strconv.FormatInt(id2, 10) + `,"x1":` + strconv.Itoa(msg.Move.X1) +
			`,"y1":` + strconv.Itoa(msg.Move.Y1) + `,"x2":` + strconv.Itoa(msg.Move.X2) +
			`,"y2":` + strconv.Itoa(msg.Move.Y2) +
			`}`
		SendSingleMsg(id2, []byte(str1))
		return
	case 10:
		//聊天
		id1, err1 := userService.GetIdByToken(msg.Token)
		if err1 != nil {
			log.Println(err1)
			return
		}
		id2, ok := global.ClientStart[id1]
		if !ok {
			//还没有开始游戏
			return
		}
		name := global.AllUserById[id1].UserName
		str1 := `{"cmd":10,"dstUsername":"` + name + `","content":"` + msg.Content + `"}`
		SendSingleMsg(id2, []byte(str1))
		return
	}
}

func init() {
	userService = new(service.UserService)
}
