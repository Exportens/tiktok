# API gateway with CloudWeGo

## Use the Hertz framework to build the gateway server, and use the generic-call capability of the Kitex framework to provide HTTP/JSON <-> Thrift protocol conversion capability. Here is the IDL specification document: IDL Definition Specification for Mapping between Thrift and HTTP .

## Implement HTTP/JSON <-> gRPC Generic-Call based on Kitex extension API; IDL specification should comply with Transcoding HTTP/JSON to gRPC

## Provide a basic management platform to manage IDLs of services, and the gateway server needs to update the IDLs regularly.

## Optional: Service discovery and load balancing are already implemented in Kitex, you can either choose a registry extension from kitex-contrib or implement one yourself; other service governance capabilities such as customized traffic routering, traffic dyeing, authentication, and rate limiter are optional.
