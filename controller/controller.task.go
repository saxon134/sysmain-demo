package controller

import (
	"encoding/json"
	"github.com/saxon134/go-utils/saLog"
	"github.com/saxon134/sysmain-demo/api"
	"github.com/saxon134/sysmain-demo/task"
	"net/http"
)

func TaskEventReceiver(w http.ResponseWriter, r *http.Request) {
	var params map[string]string
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&params)
	saLog.Log("接收到Task event：", params)

	//签名校验用
	params["sign"] = r.Header.Get("sign")
	params["timestamp"] = r.Header.Get("timestamp")

	err := task.Client.Event(params)
	if err != nil {
		api.ResError(w, err.Error())
		return
	}
	api.ResSuccess(w, nil, 0)
}

func TaskStatusReceiver(w http.ResponseWriter, r *http.Request) {
	var params = map[string]string{}
	if r.URL.Query().Get("key") == "" {
		api.ResError(w, "缺少任务key")
		return
	}
	params["key"] = r.URL.Query().Get("key")

	//签名校验用
	params["sign"] = r.Header.Get("sign")
	params["timestamp"] = r.Header.Get("timestamp")

	res, err := task.Client.Status(params)
	if err != nil {
		api.ResError(w, err.Error())
		return
	}
	api.ResSuccess(w, res, 0)
}
