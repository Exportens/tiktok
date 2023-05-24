package main

import (
	"github.com/cloudwego/hertz/pkg/app"
    	"github.com/cloudwego/hertz/pkg/app/server"
    	"github.com/cloudwego/hertz/pkg/common/utils"
    	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"os"
	"io"
)

func main() {
	f, err := os.OpenFile("/home/ubuntu/go/src/github.com/cloudwego/api_gateway/log/gateway/output.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    	if err != nil {
        	panic(err)
    	}
    	defer f.Close()
	fileWriter := io.MultiWriter(f,os.Stdout)
    	klog.SetOutput(fileWriter)	

	h := server.Default(server.WithHostPorts(":8080"))
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
            ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
    })
	register(h)
	fmt.Println("正在监听 8080 端口...")
	fmt.Println("Listening on port 8080...")
	h.Spin()
}
