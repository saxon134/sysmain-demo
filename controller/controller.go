package controller

import (
	"github.com/saxon134/sysmain-demo/api"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	api.ResSuccess(w, "hi", 0)
}
