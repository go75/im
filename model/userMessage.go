package model

type UserMessage struct {
	ID uint `gorm:"primary_key"`
	// 发送方id
	SenderId uint
	// 接收方id
	ReceiverId uint
	// 消息类型 文字为0, 图片为1, 音频为2
	Type uint
	// 消息内容
	Content string
}

func (m *UserMessage) TableName() string {
	return "user_message"
}