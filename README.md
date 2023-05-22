Explanation of project documents:

1. Cache is a git generated file, please ignore it.
 
2. The cmd folder contains a call to the server called hello, echo generalization, and its idl location is ./idl/
 
3. The hertz gateway folder serves as the gateway, and in router.go, there are generalization calls for JSON
 
4. In this folder directory: biz contains error, router, types, middleware (pls ignored), and handler. The handler contains gateway, which is the content of the generalized call client

5. The idl folder contains different idl folders:/:svc (with version)/:action should be used for calling in http

6. kitex_gen folder contains common (return error)
 
7. The log folder is the gateway and all service logs
 
8. go.mod named API_ gateway
 
9. go.sum generates files for go (derived)

######################################################################

Usage method:

localhost(can be replaced by 127.0.0.1)

Start nacos:

Run in nacos/bin with the nacos version file directory

bash startup.sh - m standalone 

Login 127.0.0.1:8848/nacos

The idl for subsequent implementation and testing is in/idl (pay attention to version)

Run go run on the hertz gateway

cmd/hello input: go run .                 -------------hello(add)

cmd/echo/echo-v1 input: go run .	-----echo-v1(add, just same as hello)

cmd/echo/echo-v2 input: go run .	-------------echo-v2(mul)

######################################################################

Gateway testing:

Base test:

ps: running with port 8080.

curl 127.0.0.1:8080 

or

curl http://127.0.0.1:8080/ping     ------will return {"message":"pong"}

Other tests:

Latest kitex-json test statements:2023.5.16（"996ers" is the return-code of running seccessfully）

The format should be

curl -X POST HTTP/1.1(optional) -H "Content-Type: application/json"(optional) 'http://domain(:port)/gateway-svcName-version/svcName-version/method' -d '{"message": "hi"}'(can be replaced)

port is the same as your gateway setting.

curl -X POST http://localhost:8080/gateway-hello-v1/hello-v1/echo -d '{"message":"Updated"}'

-------------return 996ers{\"message\": \"OK! Both svr and gateway are on the way to restart!\"}(After the gateway updates the idl directory, it receives a message from the management platform reminding it to restart)


curl -X POST http://localhost:8080/gateway-hello-v1/hello-v1/add -d '{"first":1,"second":2}' 

-------------return 996ers{\"sum\":"3"}


curl -X POST http://localhost:8080/gateway-hello-v1/hello-v1/echo -d '{"message":"hello"}'

-------------return 996ers{\"hello\":\"world\"}


curl -X POST http://localhost:8080/gateway-echo-v1/echo-v1/add -d '{"first":48,"second":51}'

-------------return 996ers{\"sum\":"99"}


curl -X POST http://localhost:8080/gateway-echo-v1/echo-v1/echo -d '{"message":"Paylah@163.com hi"}'

-------------return 996ers{\"hello\":\"world\"}


Echo has v2 version (echo v2.thrift) (the add method was renamed as sum and the mul method was added)

curl -X POST http://localhost:8080/gateway-echo-v2/echo-v2/echo -d '{"message":"Paylah@163.com hi"}'

-------------return 996ers{\"message\": \"OK! Successfully sent!\"}


curl -X POST http://localhost:8080/gateway-echo-v2/echo-v2/sum -d '{"first":5,"second":4}'

-------------return 996ers{\"sum\":"9"}


curl -X POST http://localhost:8080/gateway-echo-v2/echo-v2/mul -d '{"first":4,"second":5}'

-------------return 996ers{\"mul\":"20"}


######################################################################

Error situation:

404 no found indicates an issue with the input HTTP

10001 indicates an error in the request content

10005 indicates that request lacks of method

10006 indicates generalization failure (method does not exist)

996ers indicates that the program ran successfully without any interruption!


