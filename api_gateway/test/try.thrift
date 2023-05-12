namespace go MyService

// 定义MyService服务接口
service MyService {
    // 定义POST请求方法，接收请求参数为MyRequest结构体，返回响应参数为MyResponse结构体
    MyResponse post(MyRequest request) throws (1: MyException e),
}

// 定义请求参数结构体
struct MyRequest {
    1: string name,
    2: string email,
}

// 定义响应参数结构体
struct MyResponse {
    1: string message,
}

// 定义异常结构体
exception MyException {
    1: string code,
    2: string message,
}
///
