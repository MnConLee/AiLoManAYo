package service

import (
	"MyBot/main/constant"
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"sync"
)

var (
	bot     *openwechat.Bot
	onceBot sync.Once
)

// 全局获取bot
func GetBot() *openwechat.Bot {
	onceBot.Do(func() {
		bot = openwechat.DefaultBot(openwechat.Desktop)
	})
	return bot
}

// 获取主人号
func GetMaster() *openwechat.Friend {
	self, _ := bot.GetCurrentUser()
	friends, _ := self.Friends(true)
	// 查找指定的好友
	friendsRes := friends.SearchByUserName(1, constant.MASTERUSERNAME)
	if friendsRes.Count() < 1 {
		fmt.Println("主人没有添加机器人")
		return nil
	}
	friend := friendsRes.First()
	return friend
}

// 获取主人号
func GetFriend(nikename string) *openwechat.Friend {
	self, _ := bot.GetCurrentUser()
	friends, _ := self.Friends(true)
	// 查找指定的好友
	friendsRes := friends.SearchByNickName(1, nikename)
	if friendsRes.Count() < 1 {
		fmt.Println("没有此好友")
		return nil
	}
	friend := friendsRes.First()
	return friend
}

// 获取好友列表
func GetFriends(nikename string) openwechat.Friends {
	self, _ := bot.GetCurrentUser()
	friends, _ := self.Friends(true)
	return friends
}
