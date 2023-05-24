package main

import (
	router "api_gateway/hertz-gateway/biz/router"
	"github.com/cloudwego/hertz/pkg/app/server"
	//"api_gateway/hertz-gateway/biz/rt"
)


func register(r *server.Hertz) {
	router.GeneratedRegister(r)

	customizedRegister2(r)
}
