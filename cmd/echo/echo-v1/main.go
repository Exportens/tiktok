package main

import (
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"net"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/server/genericserver"
	"context"
	"fmt"
	"encoding/json" 
	"strconv"
	"os"
	"io"
)

func main() {
	f, err := os.OpenFile("/home/ubuntu/go/src/github.com/cloudwego/api_gateway/log/svr/echo/output_v1.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    	if err != nil {
        	panic(err)
    	}
    	defer f.Close()
    	fileWriter := io.MultiWriter(f,os.Stdout)
    	klog.SetOutput(fileWriter)
	r, err := registry.NewDefaultNacosRegistry()
	if err != nil {
		klog.Fatal(err)
	}
	p, err := generic.NewThriftFileProvider("/home/ubuntu/go/src/github.com/cloudwego/api_gateway/idl/echo/echo-v1.thrift")
    	if err != nil {
        	panic(err)
    	}
    	// 构造 JSON 请求和返回类型的泛化调用
    	g, err := generic.JSONThriftGeneric(p)
    	if err != nil {
        	panic(err)
    	}
	svr := genericserver.NewServer(new(GenericServiceImpl), g,server.WithRegistry(r),server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "echo-v1"}),server.WithServiceAddr(&net.TCPAddr{Port: 8881}),)
    	if err != nil {
        	panic(err)
    	}
    	err = svr.Run()
    	if err != nil {
        	panic(err)
    	}
    	// resp is a JSON string
}

type GenericServiceImpl struct {
}
type AddRequest struct {
    First   int `json:"\first"`
    Second    int    `json:"\second"`
}
type  AddResponse struct {
	Sum interface{} `json:"sum"`
}
type Response struct {
	Message string `json:"message"`
}
func (g *GenericServiceImpl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	reqBytes, err := json.Marshal(request)
    	if err != nil {
        	return nil, err
    	}
	str, err := strconv.Unquote(string(reqBytes))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if method == "add" {
		fmt.Printf("Recv: %v\n", str)
		var person AddRequest
		err := json.Unmarshal([]byte(str), &person)
		if err != nil {
    			fmt.Println("Error:", err)
		} else {
			c := person.First+person.Second
			addResponse :=AddResponse{ Sum: c,}
			jsonData, err := json.Marshal(addResponse)
			if err != nil {
        			return nil, err
    			}
			return  (string(jsonData)), nil
		}
        	
	}
	resp := Response{Message: "hello"}
	data, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
        if method == "echo" {
		if str == string(data){
			fmt.Printf("Recv: %v\n", str)
			fmt.Printf("Someone has found this surprise!")
			return  "{\"message\": \"world! Congratulations! You have found this surprise! So, You have a 50% or higher chance of being an IT professional...\"}", nil
		}		
	}
	fmt.Println(str)
        return  "{\"message\": \"OK! Successfully sent!\"}", nil
}
