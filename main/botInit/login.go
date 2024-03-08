package botInit

import (
	"MyBot/main/regist"
	"MyBot/main/service"
	"fmt"
	"github.com/eatmoreapple/openwechat"
)

func Login() {
	bot := service.GetBot()
	// 创建热存储容器对象
	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	defer reloadStorage.Close()
	// 执行热登录
	//bot.HotLogin(reloadStorage, openwechat.NewRetryLoginOption())
	// 执行免扫码登录
	bot.PushLogin(reloadStorage, openwechat.NewRetryLoginOption())
	// 接收消息
	regist.Receive()
	fmt.Println("登陆成功")
}
