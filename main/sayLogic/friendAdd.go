package sayLogic

import (
	"MyBot/main/service"
)

// 新好友发送
func SayfriendAdd() {
	replyMsg := "●●机器人自动执行口令\n●●查询活动请发送:活动\n●●查询全菜单发送:菜单"
	service.HandleAddFriendMessage(replyMsg)
}
