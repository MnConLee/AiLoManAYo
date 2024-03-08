package regist

import (
	"MyBot/main/sayLogic"
	"MyBot/main/service"
	"github.com/eatmoreapple/openwechat"
)

// 注册回复
func Receive() {
	bot := service.GetBot()
	// 注册消息处理函数
	bot.MessageHandler = func(msg *openwechat.Message) {
		service.Init(msg)
		// 注册言语逻辑
		// 新好友发送
		sayLogic.SayfriendAdd()
		// 对话回复
		sayLogic.SayMenu()
		sayLogic.SayHelp()
	}
	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	bot.Block()
}
