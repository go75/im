package ws

import (
	"encoding/json"
	"im/log"
	"im/ws/api"
	"im/ws/entity"
	"im/ws/handler"
	"im/ws/manager"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func Router(path string, r *gin.Engine) {
	// 升级为websocket协议
	r.GET("/upgrade", upgrade)
	// 注册friend模块路由
	registFriend()
	// 注册group模块路由
	registGroup()
	// 注册五子棋模块路由
	registFiveChess()
}

// websocket默认配置
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 升级协议
func upgrade(c *gin.Context) {
	id := c.GetUint("id")
	name := c.GetString("name")
	socket, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Warn.Println("websocket升级失败, err: ", err)
		c.JSON(http.StatusBadRequest, "websocket升级失败")
		c.Abort()
		return
	}
	defer socket.Close()
	defer manager.Remove(id)
	log.Info.Println(id, "websocket open")
	err = manager.Put(id, socket)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	var p []byte
	var obj = entity.Obj{}
	for {
		// 接收上游生产的数据
		_, p, err = socket.ReadMessage()
		if err != nil {
			log.Warn.Println("websocket read json err: ",err)
			manager.Remove(id)
			return
		}
		err := json.Unmarshal(p, &obj)
		if err != nil {
			log.Warn.Println(err)
			continue
		}

		// 填充字段
		req := &entity.Request {
			SenderId: id,
			SenderName: name,
			Obj: obj,
		}
		// 发布消息
		channel <- req
	}
}

// 注册friend模块路由
func registFriend() {
	// 获取好友会话
	api.Regist(entity.GetFriendSession, handler.GetFriendSession)
	// 获取新好友信息
	api.Regist(entity.GetNewFriend, handler.GetNewFriend)
	// 获取好友列表
	api.Regist(entity.GetFriendList, handler.GetFriendList)
	// 添加好友
	api.Regist(entity.AddFriend, handler.AddFriend)
	// 发送好友消息
	api.Regist(entity.SendFriendMsg, handler.SendFriendMsg)
	// 通过用户名称模糊查询用户
	api.Regist(entity.GetFuzzyUserByUserName, handler.GetFuzzyUserByUserName)
	// 同意新好友请求	
	api.Regist(entity.AgreeNewFriend, handler.AgreeNewFriend)
	// 拒绝新好友请求
	api.Regist(entity.RefuseNewFriend, handler.RefuseNewFriend)
}

// 注册group模块路由
func registGroup() {
	// 获取群聊会话
	api.Regist(entity.GetGroupSession, handler.GetGroupSession)
	// 获取新群友信息
	api.Regist(entity.GetNewGroup, handler.GetNewGroup)
	// 获取群聊列表
	api.Regist(entity.GetGroupList, handler.GetGroupList)
	// 添加群聊
	api.Regist(entity.AddGroup, handler.AddGroup)
	// 发送群聊消息
	api.Regist(entity.SendGroupMsg, handler.SendGroupMsg)
	// 通过群聊名称模糊查询群聊
	api.Regist(entity.GetFuzzyGroupByGroupName, handler.GetFuzzyGroupByGroupName)
	// 同意新群友请求
	api.Regist(entity.AgreeNewGroup, handler.AgreeNewGroup)
	// 拒绝新群友请求
	api.Regist(entity.RefuseNewGroup, handler.RefuseNewGroup)
}

// 注册五子棋模块路由
func registFiveChess() {
	// 请求五子棋对战
	api.Regist(entity.RequestFiveChess, handler.FiveChess)
	// 同意五子棋对战
	api.Regist(entity.AgreeFiveChess, handler.AgreeFiveChess)
	// 拒绝五子棋对战
	api.Regist(entity.RefuseFiveChess, handler.RefuseFiveChess)	
	// 发送下棋位置
	api.Regist(entity.DownChessPosition, handler.DownChessPosition)
}