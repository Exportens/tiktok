package main

import (
	"context"
	hello "api_gateway/kitex_gen/kitex_gen/hello"
)

// HelloImpl implements the last service interface defined in the IDL.
type HelloImpl struct{}

// Echo implements the HelloImpl interface.
func (s *HelloImpl) Echo(ctx context.Context, req *hello.Request) (resp *hello.Response, err error) {
	// TODO: Your code here...
	resp = &hello.Response{Message: req.Message}
	return
}

// Add implements the HelloImpl interface.
func (s *HelloImpl) Add(ctx context.Context, req *hello.AddRequest) (resp *hello.AddResponse, err error) {
	// TODO: Your code here...
	a := (*req.First) + (*req.Second)
	resp = &hello.AddResponse{Sum: &a}	
	//resp = &hello.AddResponse{Sum: *int64(*req.First) + *int64(*req.Second)}
	return
}
