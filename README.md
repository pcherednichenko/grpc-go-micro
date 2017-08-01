# Simple GRPC Go-Micro service

Count request, write logs, use Redis and Go-Micro, with generated code for Go and Ruby client

## What's here?

- **srv** - a GRPC counter service with Redis
- **cli** - a GRPC client that calls the service once


Before run service you need to run everything you need:
```
$ consul agent -dev
$ redis-server
```

Run Service
```
$ go run srv/main.go
2017/07/31 13:40:26 Listening on [::]:55490
2017/07/31 13:40:26 Broker Listening on [::]:55491
2017/07/31 13:40:26 Registering node: go.micro.srv.greeter-aa015b05-75dc-11e7-821e-784f43936a40
```

Test Go Client
```
$ go run cli/main.go
Hello Cryptopay Counter = get counter: 1
$ go run cli/main.go
Hello Cryptopay Counter = get counter: 2
$ go run cli/main.go
Hello Cryptopay Counter = get counter: 3
```
Test Ruby Client
```
cli/ruby> $ ruby test.rb
"Greeting: Hello Cryptopay Ruby Counter = get counter: 1"
cli/ruby> $ ruby test.rb
"Greeting: Hello Cryptopay Ruby Counter = get counter: 2"
cli/ruby> $ ruby test.rb
"Greeting: Hello Cryptopay Ruby Counter = get counter: 3"
```

To compile Ruby code from .proto file
```
$ grpc_tools_ruby_protoc -I ./srv/proto/hello/ --ruby_out=./srv/proto/hello --grpc_out=./srv/proto/hello ./srv/proto/hello/hello.proto
```

To compile Golang code from .proto file
```
$ protoc -I ./srv/proto/hello/ ./srv/proto/hello/hello.proto --go_out=micro:./srv/proto/hello/
```