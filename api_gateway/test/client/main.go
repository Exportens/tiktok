package main

import (
    //ctx"context"
    "github.com/cloudwego/kitex/pkg/generic"
    "github.com/cloudwego/kitex/client/genericclient"
)

func main() {
    // 本地文件 idl 解析
    // YOUR_IDL_PATH thrift 文件路径: 举例 ./idl/example.thrift
    // includeDirs: 指定 include 路径，默认用当前文件的相对路径寻找 include
    p, err := generic.NewThriftFileProvider("/home/ubuntu/go/src/github.com/cloudwego/api_gateway/test/example_service.thrift")
    if err != nil {
        panic(err)
    }
    // 构造 JSON 请求和返回类型的泛化调用
    g, err := generic.JSONThriftGeneric(p)
    if err != nil {
        panic(err)
    }
    cli, err := genericclient.NewClient("destServiceName", g)
    if err != nil {
        panic(err)
    }
    // 'ExampleMethod' 方法名必须包含在 idl 定义中
    resp, err := cli.GenericCall("", "ExampleMethod", "{\"Msg\": \"hello\"}")
    // resp is a JSON string
    
}
