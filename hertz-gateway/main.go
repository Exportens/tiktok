package main

import (
	//"github.com/cloudwego/hertz/pkg/app/server"
	//"src/github.com/cloudwego/api_gateway/hertz-gateway"
	//"api_gateway/hertz-gateway/biz/rtg"
	"github.com/cloudwego/hertz/pkg/app"
    	"github.com/cloudwego/hertz/pkg/app/server"
    	"github.com/cloudwego/hertz/pkg/common/utils"
    	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"context"
	"fmt"
	//hello "api_gateway/kitex_gen/kitex_gen/hello/hello"
	//"github.com/cloudwego/kitex/server"
	//"github.com/kitex-contrib/registry-nacos/registry"
	"github.com/cloudwego/kitex/pkg/klog"
	//"log"
	//"github.com/cloudwego/kitex/pkg/rpcinfo"
	//"net"
	"os"
	"io"
)

func main() {
	//addr := "127.0.0.1:8080"
	//fmt.Println("请输入1或者2进行选择：")
	//fmt.Println("1. 用户输入http不包含json")
	//fmt.Println("2. 用户输入http包含json")
	//var choice int
	//fmt.Scan(&choice)
///	
	//r, err := registry.NewDefaultNacosRegistry()
	//if err != nil {
	//	klog.Fatal(err)
	//}
	//h := hello.NewServer(new(HelloImpl),server.WithRegistry(r),server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "hello"}),server.WithServiceAddr(&net.TCPAddr{Port: 8080}),)
///
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
	//switch choice {
	//case 1:
		//http.HandleFunc("/", helloHandler)
		//register(h)
		//fmt.Println("正在监听 8080 端口...")
		//http.ListenAndServe(":8080", nil)
	//case 2:
		//http.HandleFunc("/", jsonHandler)
	register2(h)
	fmt.Println("正在监听 8080 端口...")
		//http.ListenAndServe(":8080", nil)
	//default:
		//fmt.Println("无效的选择...")
	//}

	//register(h)
	h.Spin()
}
