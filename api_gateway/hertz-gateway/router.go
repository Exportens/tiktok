package main

import (
	"context"
	"net/http"
	"os"
	"strings"

	"api_gateway/hertz-gateway/biz/handler"
	//"api_gateway/hertz-gateway/biz/middleware"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/registry-nacos/resolver"
	//"github.com/hertz-contrib/registry/nacos"
	//"github.com/cloudwego/hertz/pkg/app/middlewares/client/sd"
	//"log"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(http.StatusOK, "hertz-gateway is running")
	})

	registerGateway(r)
}
func customizedRegister2(r *server.Hertz) {
	r.GET("/", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(http.StatusOK, "hertz-gateway is running")
	})

	registerGateway2(r)
}
func registerGateway(r *server.Hertz) {
	//group := r.Group("/gateway").Use(middleware.GatewayAuth()...)
	group := r.Group("/gateway")
	if handler.SvcMap == nil {
		handler.SvcMap = make(map[string]genericclient.Client)
	}
	//idlPath := "./idl/"
	idlPath := "/home/ubuntu/go/src/github.com/cloudwego/api_gateway/idl/idl3/"
	c, err := os.ReadDir(idlPath)
	if err != nil {
		hlog.Fatalf("new thrift file provider failed: %v", err)
	}
	nacosResolver, err := resolver.NewDefaultNacosResolver()
	if err != nil {
		hlog.Fatalf("err:%v", err)
	}
	
	for _, entry := range c {
		//if entry.IsDir() || entry.Name() == "common.thrift" {
		//	continue
		//}
		svcName := strings.ReplaceAll(entry.Name(), ".thrift", "")

		provider, err := generic.NewThriftFileProvider(entry.Name(), idlPath)
		if err != nil {
			hlog.Fatalf("new thrift file provider failed: %v", err)
			break
		}
//http
		g, err := generic.HTTPThriftGeneric(provider)
		if err != nil {
			hlog.Fatal(err)
		}



//http
		hcli, err := genericclient.NewClient(
			svcName,
			g,
			client.WithResolver(nacosResolver),
			client.WithTransportProtocol(transport.TTHeader),
			client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
			client.WithLoadBalancer(loadbalance.NewWeightedBalancer()),
		)
		if err != nil {
			hlog.Fatal(err)
		}
		
		//hcli.Use(sd.Discovery(r))

///
		handler.SvcMap[svcName] = hcli
		//group.POST("/:svc", handler.Gateway)
		group.POST("/:svc/*action", handler.Gateway)
	}
}

//////////

func registerGateway2(r *server.Hertz) {
	//group := r.Group("/gateway").Use(middleware.GatewayAuth()...)
	group := r.Group("/gateway2")
	if handler.SvcMap == nil {
		handler.SvcMap = make(map[string]genericclient.Client)
	}
	//idlPath := "./idl/"
	idlPath := "/home/ubuntu/go/src/github.com/cloudwego/api_gateway/idl/idl2/"
	c, err := os.ReadDir(idlPath)
	if err != nil {
		hlog.Fatalf("new thrift file provider failed: %v", err)
	}
	nacosResolver, err := resolver.NewDefaultNacosResolver()
	if err != nil {
		hlog.Fatalf("err:%v", err)
	}
	for _, entry := range c {
		//if entry.IsDir() || entry.Name() == "common.thrift" {
		//	continue
		//}
		svcName := strings.ReplaceAll(entry.Name(), ".thrift", "")

		provider, err := generic.NewThriftFileProvider(entry.Name(), idlPath)
		if err != nil {
			hlog.Fatalf("new thrift file provider failed: %v", err)
			break
		}


//json
		j, err := generic.JSONThriftGeneric(provider)
    		if err != nil {
        		panic(err)
    		}


//json
		jcli, err := genericclient.NewClient(
			svcName,
			j,
			client.WithResolver(nacosResolver),
			client.WithTransportProtocol(transport.TTHeader),
			client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
			client.WithLoadBalancer(loadbalance.NewWeightedBalancer()),)
    		if err != nil {
        		panic(err)
    		}
//zhuce
		//r, err := nacos.NewDefaultNacosResolver()
		//if err != nil {
			//log.Fatal(err)
			//return
		//}
		//jcli.Use(sd.Discovery(r))
		handler.SvcMap[svcName] = jcli
		group.POST("/:svc", handler.JGateway)
	}
}

