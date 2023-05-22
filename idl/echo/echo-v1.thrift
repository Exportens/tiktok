
namespace go api


struct Request {
        1: optional string message(api.body = 'message')
}

struct Response {
        1: optional string message (api.body='message')
}

struct AddRequest {
  1: optional i64 first (api.body = 'first')
  2: optional i64 second (api.body = 'second')
}

struct AddResponse {
  1: optional i64 sum (api.body='sum')
}

service Hello {
    Response echo(1: Request req)(
	api.get = '/echo-v1/echo', api.baseurl = '127.0.0.1:8881',api.param = 'true'
	)
    AddResponse add(1: AddRequest req)( api.get = '/echo-v1/add',api.baseurl = '127.0.0.1:8881',api.param = 'true')
}
