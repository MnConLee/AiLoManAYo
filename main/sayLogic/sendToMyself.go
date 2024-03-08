package sayLogic

import (
	"MyBot/main/service"
)

func sendToMyself(msg string) {
	bot := service.GetBot()
	// 获取登录用户
	self, _ := bot.GetCurrentUser()
	// 查找指定的好友
	friends, _ := self.Friends(true)
	friends.SearchByUserName(1, msg)
}
