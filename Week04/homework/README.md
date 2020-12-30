> 按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。

## Tree

```shell
.
├── Makefile
├── README.md
├── api
│   └── helloworld
│       └── v1
│           ├── helloworld.pb.go
│           ├── helloworld.proto
│           └── helloworld_grpc.pb.go
├── bin
│   ├── greeter_client
│   └── greeter_server
├── cmd
│   ├── greeter_client
│   │   └── main.go  # grpc client
│   └── greeter_server
│       ├── main.go  # grpc server
│       ├── wire.go
│       └── wire_gen.go
├── go.mod
├── go.sum
└── internal
    ├── biz
    │   └── biz.go
    ├── data
    │   └── data.go
    ├── pkg
    │   └── service_handler
    │       └── service_handler.go  # encapsulate Service.Start, invoke GracefulStop
    └── service
        └── service.go  # api implement

13 directories, 17 files

```

## Conclusion

Test passed and log is below ↓

Server:  
```
➜  helloworld git:(main) ✗ make build
cd cmd/greeter_server && go build -o ../../bin/greeter_server && cd ../../ \
		&& cd cmd/greeter_client && go build -o ../../bin/greeter_client && cd ../../
➜  helloworld git:(main) ✗ make run
./bin/greeter_server
2020/12/30 17:45:37
grpc server start at: :50051
2020/12/30 17:45:41 Hi there is data package!
2020/12/30 17:45:56 Hi there is data package!
^C
2020/12/30 17:47:21 signal caught: interrupt, reday to quit...
2020/12/30 17:47:21 grpc server gracefully stopped.
```

Client:
```
➜  helloworld git:(main) ✗ make test1
./bin/greeter_client
2020/12/30 17:45:41 Greeting: hello pitt, Server recived your message: hi guy~ ID: 100
➜  helloworld git:(main) ✗ make test2
./bin/greeter_client gopher
2020/12/30 17:45:56 Greeting: hello gopher, Server recived your message: hi guy~ ID: 100
```

## Prerequisites

### gRPC

Reference: https://grpc.io/docs/languages/go/quickstart/#prerequisites

```
brew install protobuf
export GO111MODULE=on  # Enable module mode
go get google.golang.org/protobuf/cmd/protoc-gen-go \
         google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/protobuf/cmd/protoc-gen-go
```

copy `export PATH="$PATH:$(go env GOPATH)/bin"` append to `.zshrc`

reload iTerm2 and run below to generate pb files:
```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/helloworld/v1/helloworld.proto
```

### wire

Reference: https://github.com/google/wire

```
go get github.com/google/wire/cmd/wire
```


## Thanks

repo: https://github.com/AngelovLee/Go-001/tree/main/Week04  



