package main

import (
	payment "api_gateway/kitex_gen/payment/kitex_gen/payment/paymentsvc"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
	"github.com/cloudwego/kitex/pkg/klog"
	"log"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

func main() {
	r, err := registry.NewDefaultNacosRegistry()
	if err != nil {
		klog.Fatal(err)
	}
	svr := payment.NewServer(new(PaymentSvcImpl),server.WithRegistry(r),server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "payment"}),)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
