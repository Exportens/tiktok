
namespace go api

struct Request {
  1: string message (api.body='message')
}

struct Response {
  1: string message (api.body='message')
}
struct CulRequest {
  1: i64 first (api.body='first')
  2: i64 second (api.body='second')
}

struct MulResponse {
  1: i64 mul (api.body='mul')
}

struct AddResponse {
  1: i64 sum (api.body='sum')
}
service Echo {
    Response echo(1: Request req)
    MulResponse mul(1: CulRequest req)
    AddResponse sum(1: CulRequest req)
}
