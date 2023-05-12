package main

import (
    "fmt"
    "context"
    "github.com/cloudwego/kitex/pkg/generic"
    "github.com/cloudwego/kitex/server/genericserver"
)

func main() {
    // 本地文件 idl 解析
    // YOUR_IDL_PATH thrift 文件路径: e.g. ./idl/example.thrift
    p, err := generic.NewThriftFileProvider("/home/ubuntu/go/src/github.com/cloudwego/api_gateway/test/example_service.thrift")
    if err != nil {
        panic(err)
    }
    // 构造 JSON 请求和返回类型的泛化调用
    g, err := generic.JSONThriftGeneric(p)
    if err != nil {
        panic(err)
    }
    svc := genericserver.NewServer(new(GenericServiceImpl), g)
    if err != nil {
        panic(err)
    }
    err = svc.Run()
    if err != nil {
        panic(err)
    }
    // resp is a JSON string
}

type GenericServiceImpl struct {
}

func (g *GenericServiceImpl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
        // use jsoniter or other json parse sdk to assert request
        m := request.(string)
        fmt.Printf("Recv: %v\n", m)
        return  "{\"Msg\": \"world\"}", nil
}

