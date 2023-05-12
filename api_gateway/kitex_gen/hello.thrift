namespace go hello

struct Request {
        1: optional string message(api.query = 'message')
}

struct Response {
        1: optional string message (api.body='message')
}

struct AddRequest {
  1: optional i64 first (api.query = 'first')
  2: optional i64 second (api.query = 'second')
}

struct AddResponse {
  1: optional i64 sum (api.body='sum')
}

service Hello {
    Response echo(1: Request req)(
	api.get = '/hello/echo', api.baseurl = '127.0.0.1:8888',api.param = 'true'
	)
    AddResponse add(1: AddRequest req)( api.get = '/hello/add',api.baseurl = '127.0.0.1:8888',api.param = 'true')
}
