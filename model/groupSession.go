package model

// 群聊会话
type GroupSession struct {
	// 群聊id
	GroupId uint
	// 用户id
	UserId uint
}

func (s *GroupSession) TableName() string {
	return "group_session"
}