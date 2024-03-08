package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

// User undefined
type User struct {
	ID         int64                 `json:"id" gorm:"id"`
	NickName   string                `json:"nick_name" gorm:"nick_name"` // 微信昵称
	MagicBall  int64                 `json:"magic_ball" gorm:"magic_ball"`
	UserName   string                `json:"user_name" gorm:"user_name"`                                           // 微信用户唯一标识
	InviteId   string                `json:"Invite_id" gorm:"Invite_id"`                                           // 邀请的人id
	City       string                `json:"city" gorm:"city"`                                                     // 用户所在城市
	HeadImgUrl string                `json:"head_img_url" gorm:"head_img_url"`                                     // 用户头像
	Sex        int                   `json:"sex" gorm:"sex"`                                                       // 性别 0：男 1：女
	CreateTime time.Time             `json:"createtime" gorm:"column:create_time;type:datetime(0);autoUpdateTime"` // 创建时间
	UpdateTime time.Time             `json:"updatetime" gorm:"column:update_time;type:datetime(0);autoUpdateTime"` // 更新时间
	DeletedAt  soft_delete.DeletedAt `json:"is_deleted" gorm:"column:is_deleted;softDelete:flag"`
}

// TableName 表名称
func (*User) TableName() string {
	return "user"
}
