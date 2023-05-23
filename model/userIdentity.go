package model

type UserIdentity struct {
	// 名称
	Name string `gorm:"not null;unique"`
	// 密码
	Pwd string `gorm:"not null"`
	// 盐值
	Salt string `gorm:"not null"`
}

func (identity *UserIdentity) TableName() string {
    return "user_identity"
}