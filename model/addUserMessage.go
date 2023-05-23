package model

type AddUserMessage struct {
	// 发送方id
	SenderId uint
	// 接收方id
	ReceiverId uint
}

func (m *AddUserMessage) TableName() string {
	return "add_user_message"
}