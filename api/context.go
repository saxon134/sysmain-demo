package api

import (
	"github.com/saxon134/go-utils/saData"
	"net/http"
)

type Headers struct {
	// jwt token解析出的数据
	Account AccountJwt
}

type AccountJwt struct {
	Id int64
}

type Response struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
	Total  int64       `json:"total"` //分页时的总数
}

// Paging 获取分页数据
// limit: 默认20， offset:默认0
func Paging(r *http.Request) (offset, limit int) {
	limit, _ = saData.ToInt(r.URL.Query().Get("pageSize"))
	if limit <= 0 {
		limit = 20
	}

	num, _ := saData.ToInt(r.URL.Query().Get("pageNumber"))
	if num < 1 {
		num = 1
	}

	offset = limit * (num - 1)
	return offset, limit
}

func ResSuccess(w http.ResponseWriter, data interface{}, pagingTotal ...int64) {
	w.WriteHeader(200)
	if data == nil {
		data = []struct{}{}
	}

	var total int64
	if pagingTotal != nil && len(pagingTotal) > 0 {
		total = pagingTotal[0]
	}
	_, _ = w.Write([]byte(saData.String(Response{Code: 0, Result: data, Total: total})))
	return
}

func ResError(w http.ResponseWriter, errMsg string) {
	w.WriteHeader(400)
	_, _ = w.Write([]byte(saData.String(Response{Code: 0, Result: []struct{}{}, Msg: errMsg})))
	return
}
