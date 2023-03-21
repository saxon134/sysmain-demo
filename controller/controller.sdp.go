package controller

import (
	"encoding/json"
	"github.com/saxon134/go-utils/saLog"
	"github.com/saxon134/sysmain-demo/api"
	"net/http"
)

func SDPMsgReceiver(w http.ResponseWriter, r *http.Request) {
	var params map[string]string
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&params)
	saLog.Log("接收到SDP信息：", params["msg"])

	api.ResSuccess(w, nil, 0)
}
