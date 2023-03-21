package main

import (
	"github.com/saxon134/go-utils/saLog"
	"github.com/saxon134/sysmain-demo/conf"
	"github.com/saxon134/sysmain-demo/http"
	"github.com/saxon134/sysmain-demo/sdp"
	"github.com/saxon134/sysmain-demo/task"
)

func main() {
	//初始化
	conf.Init()

	//初始化日志
	saLog.Init(saLog.WarnLevel, saLog.ZapType)

	//初始化http服务
	go http.Init()

	//初始化SDP
	sdp.Init()

	//初始化task
	task.Init()

	//防止应用退出
	<-make(chan bool)
}
