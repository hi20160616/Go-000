# Week09 作业题目

用 Go 实现一个 tcp server ，用两个 goroutine 读写 conn，两个 goroutine 通过 chan 可以传递 message，能够正确退出

# Coding logs

## A simple TCP Server and Client

Reference:
- https://www.linode.com/docs/guides/developing-udp-and-tcp-clients-and-servers-in-go/

## Test

```
make build
make run
```

```
make test
```

### Output

Server
```
make run
./bin/tcpserver 1234
-> Hello world!
-> Good day!
Exiting TCP server!
```

Client
```
make test
./bin/tcpclient 127.0.0.1:1234
>> Hello
->: 2021-01-25T12:51:49+08:00
>> Good day!
->: 2021-01-25T12:52:02+08:00
>> STOP
->: TCP client exiting...
```
