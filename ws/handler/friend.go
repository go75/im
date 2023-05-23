package handler

import (
	"im/dao"
	"im/global"
	"im/log"
	"im/model"
	"im/utils"
	"im/ws/entity"
)

// 获取好友会话
func GetFriendSession(r *entity.Request) {
	msgs, _ := dao.QueryUserMessageByIds(r.SenderId, r.ProcessId)
	err := utils.Send(r.SenderId, r.ID, entity.Text, 0, msgs)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 获取新好友信息
func GetNewFriend(r *entity.Request) {
	ids, _ := dao.GetAddUserSenderIdsByReveiverId(r.SenderId)
	if len(ids) == 0 {
		return
	}
	users, _ := dao.GetUsersByIds(ids)
	err := utils.Send(r.SenderId, r.ID, entity.Text, 0, users)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 发送好友消息
func SendFriendMsg(r *entity.Request) {
	// 用户消息持久化
	msg := &model.UserMessage {
		ReceiverId: r.ProcessId,
		SenderId: r.SenderId,
		Type: r.Type,
		Content: r.Payload,
	}
	dao.CreateUserMessage(msg)
	err := utils.Send(r.ProcessId, r.ID, r.Type, r.SenderId, msg)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 获取好友列表
func GetFriendList(r *entity.Request) {
	// 查询用户会话
	users := make([]*model.User, 0)
	
	db := global.DB.Raw("select user.id, user.name, user.introduce from user, user_session where user_session.small_id = ? and user.id = user_session.big_id", r.SenderId).Scan(&users)
	log.Info.Println(users)
	
	tmp := make([]*model.User, 0)
	db.Raw("select user.id, user.name, user.introduce from user, user_session where user_session.big_id = ? and user.id = user_session.small_id", r.SenderId).Scan(&tmp)
	log.Info.Println(tmp)

	users = append(users, tmp...)
	err := utils.Send(r.SenderId, r.ID, r.Type, r.ProcessId, users)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 添加好友
func AddFriend(r *entity.Request) {
	// 新好友消息持久化
	msg := &model.AddUserMessage {
		SenderId: r.SenderId,
		ReceiverId: r.ProcessId,
	}
	dao.CreateAddUserMessage(msg)
	user := &model.User {
		ID: msg.SenderId,
	}
	dao.GetUserById(user)
	err := utils.Send(r.ProcessId, r.ID, entity.Text, r.ProcessId, user)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 通过用户名称模糊查询用户
func GetFuzzyUserByUserName(r *entity.Request) {
	ls, _ := dao.GetFuzzyUserByUserName(r.Payload)
	err := utils.Send(r.SenderId, r.ID, entity.Text, 0, ls)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 同意新好友请求
func AgreeNewFriend(r *entity.Request) {
	var session *model.UserSession
	if r.ProcessId < r.SenderId {
		session = &model.UserSession {
			SmallId: r.ProcessId,
			BigId: r.SenderId,
		}
	} else if r.ProcessId > r.SenderId {
		session = &model.UserSession {
			SmallId: r.SenderId,
			BigId: r.ProcessId,
		}
	} else {
		return
	}

	// 开启事务
	transaction := global.DB.Begin()
	db := dao.DeleteAddUserMessage(r.ProcessId, r.SenderId)
	if db.RowsAffected != 1 {
		// 事务回滚
		transaction.Rollback()
		return
	}
	db = dao.CreateUserSession(session)
	if db.RowsAffected != 1 {
		// 事务回滚
		transaction.Rollback()
		return
	}

	// 提交事务
	transaction.Commit()

	user := &model.User {
		ID: r.SenderId,
	}
	dao.GetUserById(user)
	if user.Name == "" {
		log.Warn.Printf("id为%d的用户信息获取失败\n", user.ID)
		return
	}

	// 返回新好友的信息
	err := utils.Send(r.ProcessId, r.ID, entity.Text, 0, user)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 拒绝新好友请求
func RefuseNewFriend(r *entity.Request) {
	dao.DeleteAddUserMessage(r.ProcessId, r.SenderId)
}