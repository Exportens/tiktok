# API_gateway
项目文件解释：

1.请忽略client和test文件夹

2.cmd文件夹包含一个名为hello的idl的注册服务以及服务端，其idl位置为./idl/idl3/hello.thrift，使用前运行nacos后，在该目录下运行：go run .
3.hertz-gateway为网关，四个phg为main运行go run ./hertz-gateway，router.go里面针对http和json进行不同的泛化调用步骤

              该文件夹目录下：biz包含error，router，types，middleware（忽略），和handler，handler包含gateway，内容是泛化调用客户端内容
              
4.idl文件下包含不同idl文件：问题在于使用/:svc分支post，期待能进行优化！

5.kitex_gen包含common（返回error），echo文件夹（请忽略！），payment：是针对idl/idl4/payment.thrift的使用kitex命令生成的服务框架和内容；其余文件，均为hello.thrift的服务框架

6.pkg包含授权，但是请忽略

7.go.mod名为api_gateway

8.go.sum 为go（衍生）生成文件

测试：
本地运行nacos（略）127.0.0.1:8848/nacos

go run .
listening on address[::]8888 listening
并返回日志在nacos中！

go run ./hertz-gateway
listening on address[::]8080 listening

curl 127.0.0.1:8080
hertz-gateway is running
连通！

curl http://127.0.0.1:8080/ping
{"message":"pong"}
证明网关在运行！

