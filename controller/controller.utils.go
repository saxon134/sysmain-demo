package controller

import (
	"github.com/saxon134/go-utils/saData"
	"github.com/saxon134/sysmain-demo/conf"
	"time"
)

//func checkSign(r *http.Request) bool {
//	if conf.Conf.Sysmain.Secret == "" {
//		return true
//	}
//
//	var sign = r.Header.Get("sign")
//	var timestamp = r.Header.Get("timestamp")
//	if sign == "" || timestamp == "" {
//		return false
//	}
//
//	sign2 := saData.Md5(conf.Conf.Sysmain.Secret+timestamp, true)
//	return sign == sign2
//}

func genSign() (sign string, timestamp string) {
	if conf.Conf.Sysmain.Secret == "" {
		return
	}

	timestamp = saData.String(time.Now().Unix())
	sign = saData.Md5(conf.Conf.Sysmain.Secret+timestamp, true)
	return sign, timestamp
}
