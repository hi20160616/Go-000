# Week09 作业题目

用 Go 实现一个 tcp server ，用两个 goroutine 读写 conn，两个 goroutine 通过 chan 可以传递 message，能够正确退出

以上作业，要求提交到 GitHub 上面，Week09 作业地址：  
https://github.com/Go-000/Go-000/issues/82  

作业提交截止时间 1 月 28 日（周四）23:59 前。

学号：G20200607010680

# Coding logs

## Test this case

```
make build
make run
```

```
make test
```

Server
```
➜  tcpserver git:(main) ✗ make build
cd cmd/tcpserver && go build -o ../../bin/tcpserver && cd ../../ \
		&& cd cmd/tcpclient && go build -o ../../bin/tcpclient && cd ../../
➜  tcpserver git:(main) ✗ l
total 80
drwxr-xr-x  10 foobar  staff   320B  1 27 13:13 .
drwxr-xr-x   4 foobar  staff   128B  1 27 10:06 ..
-rw-r--r--   1 foobar  staff    12K  1 27 13:13 .Makefile.swp
-rw-r--r--   1 foobar  staff    12K  1 27 13:12 .README.md.swp
-rw-r--r--   1 foobar  staff   192B  1 27 13:13 Makefile
-rw-r--r--   1 foobar  staff   1.0K  1 27 13:12 README.md
drwxr-xr-x   4 foobar  staff   128B  1 27 13:13 bin
drwxr-xr-x   4 foobar  staff   128B  1 26 13:24 cmd
-rw-r--r--   1 foobar  staff   133B  1 27 10:06 go.mod
-rw-r--r--   1 foobar  staff   209B  1 27 10:06 go.sum
➜  tcpserver git:(main) ✗ make run
./bin/tcpserver
Hi there, server working at  :12345
->  Hi there
->  init 6
Server time: 2021-01-27T13:14:33+08:00
recive command: RESTART
Stop tcp server...
Start tcp server...
->  Hello world!
->  init 0
Server time: 2021-01-27T13:15:05+08:00
recive command: STOP
2021/01/27 13:15:05 tcp server stop now...
➜  tcpserver git:(main) ✗ make run
./bin/tcpserver
Hi there, server working at  :12345
^C
2021/01/27 13:15:26 signal caught: interrupt, ready to quit...
2021/01/27 13:15:26 tcp server stop now...
```
Client
```
➜  tcpserver git:(main) ✗ make test
./bin/tcpclient
Welcome to my dark side...
Hi there
->: Server time: 2021-01-27T13:14:13+08:00
RESTART
Send restart signal to TCP Server...
->: Server time: 2021-01-27T13:14:33+08:00
EXIT
TCP Client exiting...
➜  tcpserver git:(main) ✗ make test
./bin/tcpclient
Welcome to my dark side...
Hello world!
->: Server time: 2021-01-27T13:14:58+08:00
STOP
Send stop signal to TCP Server...
->: Server time: 2021-01-27T13:15:05+08:00
client routine stoped.
EXIT
TCP Client exiting...
➜  tcpserver git:(main) ✗ make test
./bin/tcpclient
Welcome to my dark side...
client routine stoped.
^C
EXIT
TCP Client exiting...
```

## A simple TCP Server and Client

this case at parent folder: `simple_tcpserver`

Reference:
- https://www.linode.com/docs/guides/developing-udp-and-tcp-clients-and-servers-in-go/

## Test

```
make build
make run
```

```
make test1
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
make test1
./bin/tcpclient 127.0.0.1:1234
>> Hello
->: 2021-01-25T12:51:49+08:00
>> Good day!
->: 2021-01-25T12:52:02+08:00
>> STOP
->: TCP client exiting...
```
