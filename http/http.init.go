package http

import (
	"github.com/saxon134/go-utils/saData"
	"github.com/saxon134/go-utils/saData/saUrl"
	"github.com/saxon134/go-utils/saLog"
	"github.com/saxon134/sysmain-demo/conf"
	"github.com/saxon134/sysmain-demo/controller"
	"net/http"
)

// Init 阻塞
func Init() {
	if conf.Conf.Http.Port <= 0 {
		panic("http port can not be empty")
	}

	//SDP
	http.HandleFunc(saUrl.ConnPath(conf.Conf.Sysmain.ClientRoot, "sdp/msg"), controller.SDPMsgReceiver)

	//Task
	http.HandleFunc(saUrl.ConnPath(conf.Conf.Sysmain.ClientRoot, "task/event"), controller.TaskEventReceiver)
	http.HandleFunc(saUrl.ConnPath(conf.Conf.Sysmain.ClientRoot, "task/status"), controller.TaskStatusReceiver)

	saLog.Log("Http listening on " + saData.String(conf.Conf.Http.Port))
	err := http.ListenAndServe(":"+saData.String(conf.Conf.Http.Port), nil)
	if err != nil {
		panic("http err:" + err.Error())
	}
}
