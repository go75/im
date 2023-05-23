package model

// 用户会话
type UserSession struct {
	// 两个用户中id较小值的名称
	SmallId uint
	// 两个用户中id较大值的名称
	BigId uint
}

func (s *UserSession) TableName() string {
	return "user_session"
}
