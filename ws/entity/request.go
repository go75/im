package entity

type Request struct {
	// 发送方id
	SenderId uint
	// 发送方名称
	SenderName string
	// 请求对象
	Obj
}