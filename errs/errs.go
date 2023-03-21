package errs

import (
	"fmt"
	"github.com/saxon134/go-utils/saData/saError"
	"runtime"
)

var (
	ErrorTooFrequent  = saError.Error{Code: saError.BeDisplayedErrorCode, Msg: "操作太频繁"}
	ErrorParams       = saError.Error{Code: saError.NormalErrorCode, Msg: "缺少必要参数"}
	ErrorData         = saError.Error{Code: saError.NormalErrorCode, Msg: "数据有误"}
	ErrorEmpty        = saError.Error{Code: saError.NormalErrorCode, Msg: "数据为空"}
	ErrorNotSupport   = saError.Error{Code: saError.NormalErrorCode, Msg: "暂不支持"}
	ErrorUnauthorized = saError.Error{Code: saError.LoggedFailErrorCode, Msg: "未授权"}
	ErrorNotExisted   = saError.Error{Code: saError.BeDisplayedErrorCode, Msg: "不存在"}
)

func New(err interface{}, params ...interface{}) error {
	if err == nil {
		return nil
	}

	var resErr = saError.Error{Code: saError.NormalErrorCode, Msg: "", Caller: ""}
	if s, ok := err.(string); ok {
		resErr.Msg = s
		resErr.Code = saError.NormalErrorCode
	} else {
		e, ok := err.(*saError.Error)
		if ok == false {
			var e2 saError.Error
			if e2, ok = err.(saError.Error); ok {
				e = &e2
			}
		}

		if e != nil {
			if len(e.Msg) > 0 {
				resErr.Msg = e.Msg
			}
			if e.Code > 0 {
				resErr.Code = e.Code
			}
			if e.Caller != "" {
				if resErr.Caller == "" {
					resErr.Caller = e.Caller
				} else {
					resErr.Caller = e.Caller + "\n" + resErr.Caller
				}
			}
		} else if e, ok := err.(error); ok {
			resErr.Msg = e.Error()
			resErr.Code = saError.SensitiveErrorCode
		} else {
			return nil
		}
	}

	if params != nil {
		for _, v := range params {
			if code, ok := v.(int); ok {
				if code > 0 {
					resErr.Code = code
				}
			} else if s, ok := v.(string); ok {
				if s != "" {
					resErr.Msg += s
				}
			} else {
				e, ok := err.(*saError.Error)
				if ok == false {
					var e2 saError.Error
					if e2, ok = err.(saError.Error); ok {
						e = &e2
					}
				}

				if e != nil {
					if len(e.Msg) > 0 {
						resErr.Msg += s
					}
					if e.Code > 0 {
						resErr.Code = e.Code
					}
					if e.Caller != "" {
						resErr.Caller = e.Caller + "\n" + resErr.Caller
					}
				} else if e, ok := err.(error); ok {
					resErr.Msg += e.Error()
				}
			}
		}
	}

	//获取调用栈
	pc := make([]uintptr, 1)
	n := runtime.Callers(2, pc)
	if n >= 1 {
		f := runtime.FuncForPC(pc[0])
		file, line := f.FileLine(pc[0])
		resErr.Caller = fmt.Sprintf("%s:%d\n", file, line) + resErr.Caller
	}
	return &resErr
}
