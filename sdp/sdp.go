package sdp

import (
	"fmt"
	"github.com/saxon134/go-utils/saHttp"
	"github.com/saxon134/go-utils/saLog"
	"github.com/saxon134/sysmain-demo/conf"
	"github.com/saxon134/sysmain-demo/errs"
	"github.com/saxon134/sysmain/sdp"
	"time"
)

var Client *sdp.Client

func Init() {
	Client = sdp.NewClient(sdp.Conf{
		SysmainUrl: conf.Conf.Sysmain.Url,
		ClientRoot: conf.Conf.Sysmain.ClientRoot,
		Secret:     conf.Conf.Sysmain.Secret,
		PingSecond: 3,
		Redis: struct {
			Host string
			Pass string
		}{Host: conf.Conf.Redis.Host, Pass: conf.Conf.Redis.Pass},
	})
	var host = conf.Conf.Http.Host
	if host == "" {
		ary, err := saHttp.GetLocalIP()
		if len(ary) == 0 {
			panic("get local ip error: " + err.Error())
		}
		host = ary[0]
	}
	Client.Register(conf.Conf.Name, host, conf.Conf.Http.Port)

	//////// 以下是使用DEMO ////////
	go func() {
		time.Sleep(time.Second * 3)

		//发现服务
		host, port := Client.Discovery(conf.Conf.Name)
		saLog.Log(fmt.Sprintf("发现服务 %s：%s:%d", conf.Conf.Name, host, port))

		//发送消息
		err := Client.SendMsg(conf.Conf.Name, host, port, "SDP发送消息测试")
		if err != nil {
			saLog.Err(errs.New(err))
		}
	}()
}
