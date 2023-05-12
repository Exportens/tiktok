package main

import (
	hello "api_gateway/kitex_gen/kitex_gen/hello/hello"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
	"github.com/cloudwego/kitex/pkg/klog"
	"log"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"net"
)

func main() {
	r, err := registry.NewDefaultNacosRegistry()
	if err != nil {
		klog.Fatal(err)
	}
	svr := hello.NewServer(new(HelloImpl),server.WithRegistry(r),server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "hello"}),server.WithServiceAddr(&net.TCPAddr{Port: 8888}),)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
