package model

type Group struct {
	ID   uint `gorm:"primary_key"`
	// 群聊名称
	Name string
	// 群主id
	MasterId uint
	// 群聊简介
	Introduce string
}

func (g *Group) TableName() string {
	return "group"
}
