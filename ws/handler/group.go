package handler

import (
	"im/dao"
	"im/global"
	"im/log"
	"im/model"
	"im/utils"
	"im/ws/entity"
	"strings"
)

// 获取群聊会话
func GetGroupSession(r *entity.Request) {
	msgs, _ := dao.QueryGroupMessage(r.ProcessId)
	if len(msgs) == 0 {
		return
	}
	sendMsgs := make([]*entity.GroupMessage, len(msgs))
	for i := 0; i < len(msgs); i++ {
		user := &model.User{ID: msgs[i].SenderId}
		dao.GetUserNameById(user)
		// 发布消息
		sendMsgs[i] = &entity.GroupMessage{
			ID: msgs[i].ID,
			SenderId: msgs[i].SenderId,
			SenderName: user.Name,
			Type: msgs[i].Type,
			Content: msgs[i].Content,
		}
	}
	err := utils.Send(r.SenderId, r.ID, entity.Text, r.ProcessId, sendMsgs)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 获取新群聊信息
func GetNewGroup(r *entity.Request) {
	// 获取当前用户创建过的所有群聊id
	groupIds, _ := dao.QueryGroupIdsByMasterId(r.SenderId)

	if groupIds == nil {
		log.Info.Println("-------------------------")
	}

	msgs, _ := dao.GetAddGroupMessagesByGroupIds(groupIds)
	log.Info.Println(msgs)
	if len(msgs) == 0 {
		return
	}

	sendMsgs := make([]*entity.NewGroupMessage, len(msgs))

	// 构造发送的消息
	for i := 0; i < len(sendMsgs); i++ {
		group := &model.Group{
			ID: msgs[i].GroupId,
		}
		dao.QueryGroupNameByGroupId(group)
		user := &model.User{
			ID: msgs[i].SenderId,
		}
		dao.GetUserNameById(user)
		sendMsgs[i] = &entity.NewGroupMessage{
			GroupId: group.ID,
			GroupName: group.Name,
			Username: user.Name,
		}
	}

	err := utils.Send(r.SenderId, r.ID, entity.Text, 0, sendMsgs)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 发送群聊消息
func SendGroupMsg(r *entity.Request) {
	// 群聊消息持久化
	msg := &model.GroupMessage {
		GroupId: r.ProcessId,
		SenderId: r.SenderId,
		Type: r.Type,
		Content: r.Payload,
	}
	user := &model.User{ID: r.SenderId}
	dao.GetUserNameById(user)
	if user.Name == "" {
		// 该用户不存在
		return
	} 
	// 消息持久化
	dao.CreateGroupMessage(msg)
	// 发布消息
	sendMsg := entity.GroupMessage{
		SenderId: user.ID,
		SenderName: user.Name,
		Type: r.Type,
		Content: r.Payload,
	}
	session := &model.GroupSession{
		GroupId: r.ProcessId,
	}
	ls, _ := dao.QueryGroupSessionsByGroupId(session)
	for _, session := range ls {
		utils.Send(session.UserId, r.ID, entity.Text, r.ProcessId, sendMsg)
	}
}

// 获取群聊列表
func GetGroupList(r *entity.Request) {
	ids, _ := dao.QueryGroupIdsByUserId(r.SenderId)
	if len(ids) == 0 {
		return
	}
	groups, db := dao.QueryGroupsByGroupIds(ids)
	if db.Error != nil {
		return
	}
	err := utils.Send(r.SenderId, r.ID, entity.Text, 0, groups)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 添加群聊
func AddGroup(r *entity.Request) {
	// 新群友消息持久化
	msg := &model.AddGroupMessage{
		SenderId: r.SenderId,
		GroupId: r.ProcessId,
	}

	if dao.CreateAddGroupMessage(msg).RowsAffected != 1 {
		// 请求失败
		log.Info.Println("创建添加群聊消息失败")
		return
	}

	group := &model.Group{
		ID: r.ProcessId,
	}
	dao.QueryGroupByGroupId(group)
	if group.Name == "" {
		log.Info.Println("群聊查询失败")
		return
	}
	sendMsg := entity.NewGroupMessage{
		GroupId: group.ID,
		GroupName: group.Name,
		Username: r.SenderName,
	}
	err := utils.Send(group.MasterId, r.ID, r.Type, r.ProcessId, sendMsg)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 通过群聊名称模糊查询群聊
func GetFuzzyGroupByGroupName(r *entity.Request) {
	ls, _ := dao.GetFuzzyGroupByGroupName(r.Payload)
	err := utils.Send(r.SenderId, r.ID, entity.Text, 0, ls)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 同意新群友请求
func AgreeNewGroup(r *entity.Request) {
	user := &model.User{
		Name: r.Payload,
	}

	// 通过 加好友的请求方名称 获取 加好友的请求方id
	dao.GetUserByName(user)
	
	if user.ID == 0 {
		return
	}

	var session = &model.GroupSession{
		GroupId: r.ProcessId,
		UserId: user.ID,
	}

	msg := &model.AddGroupMessage {
		SenderId: user.ID,
		GroupId:  r.ProcessId,
	}
	// 开启事务
	transaction := global.DB.Begin()
	
	db := dao.DeleteAddGroupMessageBySenderIdAndGroupId(msg)
	if db.RowsAffected != 1 {
		// 事务回滚
		transaction.Rollback()
		return
	}
	log.Info.Println("create session", session)
	db = dao.CreateGroupSession(session)
	if db.RowsAffected != 1 {
		// 事务回滚
		transaction.Rollback()
		return
	}

	group := &model.Group {
		ID: r.ProcessId,
	}
	dao.QueryGroupByGroupId(group)
	if group.Name == "" {
		log.Warn.Printf("id为%d的群聊信息获取失败\n", group.ID)
		transaction.Rollback()
		return
	}

	// 提交事务
	transaction.Commit()
	// 返回新群聊的信息
	err := utils.Send(user.ID, r.ID, entity.Text, 0, group)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 拒绝新群友请求
func RefuseNewGroup(r *entity.Request) {
	msg := &model.AddGroupMessage{
		SenderId: r.ProcessId,
		GroupId: r.SenderId,
	}
	dao.DeleteAddGroupMessageBySenderIdAndGroupId(msg)
}

// 获取群聊聊天记录
func GetGroupMsgs(r *entity.Request) {
	// 获取聊天记录
	ls, _ := dao.QueryGroupMessage(r.ProcessId)
	if len(ls) == 0 {
		return
	}
	// 响应聊天记录
	err := utils.Send(r.SenderId, r.ID, entity.Text, r.ProcessId, ls)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 创建群聊
func CreateGroup(r *entity.Request) {
	parts := strings.Split(r.Payload, " ")
	if len(parts) == 2 {
		group := &model.Group{
			Name: parts[0],
			Introduce: parts[1],
		}
		transaction := global.DB.Begin()
		if dao.CreateGroup(group).RowsAffected == 1 {
			if dao.GetGroupByName(group).Error == nil {
				utils.Send(r.SenderId, r.ID, entity.Text, group.ID, "ok" )
				session := &model.GroupSession{
					GroupId: group.ID,
					UserId: r.SenderId,
				}
				if dao.CreateGroupSession(session).RowsAffected == 1 {
					transaction.Commit()
					return
				}
			}
			transaction.Rollback()
		}
	}
	utils.Send(r.SenderId, r.ID, entity.Text, r.ProcessId, "err")
}