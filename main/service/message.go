package service

import (
	"MyBot/main/dao"
	"github.com/eatmoreapple/openwechat"
)

var globalMsg *openwechat.Message

func Init(msg *openwechat.Message) {
	globalMsg = msg
}

// 处理对话文本内容
func handleTextContent(message string, replyMsg string) {
	if globalMsg.Content == message {
		globalMsg.ReplyText(replyMsg)
	}
}

// 处理对话文本消息
func HandleTextMessage(message string, replyMsg string) {
	// 判断消息是否为文本消息
	if globalMsg.IsText() {
		// 调用处理文本内容的函数
		handleTextContent(message, replyMsg)
	}
}

// 处理新好友招呼文本消息
func HandleAddFriendMessage(replyMsg string) {
	// 判断消息是否有好友申请
	if globalMsg.IsFriendAdd() {
		// 同意好友
		friend, err := globalMsg.Agree()
		if err != nil {
			GetMaster().SendText("无法同意好友申请")
			return
		}
		// 保存好友信息
		dao.SaveFriend(friend)
		// 发送消息
		friend.SendText(replyMsg)
	}
}
