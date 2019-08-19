package main

import (
	"gin/Router"
)

func main() {
	// 初始化引擎
	engine := Router.InitRouter()
	// 绑定端口，然后启动应用
	engine.Run(":9205")
}
