package main

import (
	"MyBot/main/base"
	"MyBot/main/botInit"
)

func main() {
	// 初始化数据库连接gorm
	base.InitDB()
	// 初始化bot
	botInit.Login()
}
