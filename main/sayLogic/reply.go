package sayLogic

import (
	"MyBot/main/service"
)

/**
 *言语回复逻辑
 */
// 菜单说明
func SayMenu() {
	msg := "菜单"
	replyMsg := "●●机器人自动执行口令\n●●查询活动请发送:活动\n●●充值积分请发送:充值\n●●查询积分请发送:余额\n●●售后客服请发送:客服\n●●积分比例■1元=1🔮\n每推荐一个好友得1🔮\n累计推荐五名好友可升级代理\n📕如何操作升级代理发送:代理"
	service.HandleTextMessage(msg, replyMsg)
}

// 帮助说明
func SayHelp() {
	msg := "活动"
	replyMsg := "活动名--代理价格--零售价格\n\n沪上阿姨--3.2🔮--3.5🔮\n"
	service.HandleTextMessage(msg, replyMsg)
}
