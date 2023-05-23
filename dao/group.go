package dao

import (
	"im/global"
	"im/model"

	"gorm.io/gorm"
)

// 创建群聊
func CreateGroup(group *model.Group) *gorm.DB {
	return global.DB.Create(group)
}

// 删除群聊
func DeleteGroup(group *model.Group) *gorm.DB {
	return global.DB.Delete(group)
}

// 通过群聊id获取群聊
func GetGroupById(group *model.Group) *gorm.DB {
	return global.DB.Where("id = ?", group.ID).First(group)
}

// 通过群聊名称获取群聊
func GetGroupByName(group *model.Group) *gorm.DB {
	return global.DB.Where("name = ?", group.Name).First(group)
}

// 修改群聊信息
func UpdateGroup(group *model.Group) *gorm.DB {
	return global.DB.Model(group).Updates(model.Group{
		Name: group.Name,
		Introduce: group.Introduce,
	})
}

// 通过群主id查找所有的群聊id
func QueryGroupIdsByMasterId(masterId uint) ([]uint, *gorm.DB) {
	var list []uint
	db := global.DB.Model(&model.Group{}).Select("id").Where("master_id = ?", masterId).Find(&list)
	return list, db
}

func GetFuzzyGroupByGroupName(name string) (groups []*model.Group, db *gorm.DB){
	db = global.DB.Where("name like ?", "%"+name+"%").Find(&groups)
	return
}

func QueryGroupByGroupId(group *model.Group) *gorm.DB {
	return global.DB.First(group)
}

func QueryGroupNameByGroupId(group *model.Group) *gorm.DB {
	return global.DB.Select("name").Where("id=?",group.ID).Find(group)
}

func QueryGroupsByGroupIds(ids []uint) ([]*model.Group, *gorm.DB) {
	groups := make([]*model.Group, 0)
	db := global.DB.Find(&groups, ids)
	return groups, db
}