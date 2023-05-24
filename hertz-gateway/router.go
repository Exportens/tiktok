package main

import (
	"context"
	"net/http"
	"os"
	"strings"
	"github.com/zz1219/API_gateway/blob/masterapi_gateway/hertz-gateway/biz/handler"
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
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// customizeRegister registers customize routers.

func customizedRegister2(r *server.Hertz) {
	r.GET("/", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(http.StatusOK, "hertz-gateway is running")
	})

	registerGateway2(r)
}


func registerGateway2(r *server.Hertz) {
	if handler.SvcMap == nil {
		handler.SvcMap = make(map[string]genericclient.Client)
	}
	nacosResolver, err := resolver.NewDefaultNacosResolver()
	if err != nil {
		hlog.Fatalf("err:%v", err)
	}
	idlPath := "/home/ubuntu/go/src/github.com/cloudwego/api_gateway/idl/"
	subDirs, err := ioutil.ReadDir(idlPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}
	for _, subDir := range subDirs {
		if subDir.IsDir() {
			subDirPath := filepath.Join(idlPath, subDir.Name())
			c, err := os.ReadDir(subDirPath)
			if err != nil {
				fmt.Println("Error reading directory:", err)
				return
			}
			fmt.Println(subDirPath)
			for _, entry := range c {
				svcName := strings.ReplaceAll(entry.Name(), ".thrift", "")
				group := r.Group(svcName)
				provider, err := generic.NewThriftFileProvider(entry.Name(), subDirPath)
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
					client.WithLoadBalancer(loadbalance.NewWeightedRoundRobinBalancer()),)
				if err != nil {
					panic(err)
				}
				handler.SvcMap[svcName] = jcli
				group.POST("/*action", handler.JGateway)

			}
		}
	}
}
