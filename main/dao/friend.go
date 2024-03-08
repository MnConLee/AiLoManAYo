package dao

import (
	"MyBot/main/base"
	"MyBot/main/form"
	"github.com/eatmoreapple/openwechat"
)

// 保存好友信息
func SaveFriend(friend *openwechat.Friend) {
	// 对用户信息进行保存
	user := form.User{
		NickName:   friend.NickName,
		UserName:   friend.UserName,
		City:       friend.City,
		HeadImgUrl: friend.HeadImgUrl,
		Sex:        friend.Sex,
	}
	base.DB().Create(user)
}
