package task

import (
	"github.com/saxon134/go-utils/saHttp"
	"github.com/saxon134/sysmain-demo/conf"
	"github.com/saxon134/sysmain/task"
)

var Client *task.Client

func Init() {
	var host = conf.Conf.Http.Host
	if host == "" {
		ary, err := saHttp.GetLocalIP()
		if len(ary) == 0 {
			panic("get local ip error: " + err.Error())
		}
		host = ary[0]
	}
	Client = task.Init(conf.Conf.Name, conf.Conf.Sysmain.Url, conf.Conf.Sysmain.Secret, host, conf.Conf.Http.Port)
	Client.Register(
		task.Case{Key: "SayHello", Spec: "", Local: false, Handler: SayHello},
		//task.Case{Key: "LocalTask", Spec: "*/2 * * * * *", Local: true, Handler: LocalTask},
	)
}
