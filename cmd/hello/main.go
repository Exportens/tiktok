
package main

import (
	//"net"

	"api_gateway/kitex_gen/kitex_gen/hello/hello"
	"github.com/cloudwego/kitex/pkg/klog"
	//"github.com/cloudwego/kitex/pkg/rpcinfo"
	//"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
)

func main() {
	addr := "127.0.0.1:8081"
	r, err := registry.NewDefaultNacosRegistry()
	if err != nil {
		klog.Fatal(err)
	}

	//hdlr := initHandler()

	//svr := hello.NewServer(
		//hdlr,
		//server.WithRegistry(r),
		//server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "hello"}),
		//server.WithServiceAddr(&net.TCPAddr{Port: 8081}),
		//server.WithMetaHandler(transmeta.ServerTTHeaderHandler),
	//)
///
	svr := hello.Default(
        	server.WithHostPorts(addr),
        	server.WithRegistry(r, &registry.Info{
            		ServiceName: "hello",
            		Addr:        utils.NewNetAddr("tcp", addr),
            		Weight:      10,
            		Tags:        nil,
        	}),
    	)


	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
