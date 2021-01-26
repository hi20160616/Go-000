build:
	cd server && go build -o ../bin/tcpserver && cd ../ \
		&& cd client && go build -o ../bin/tcpclient && cd ../

run:
	./bin/tcpserver 1234

test1:
	./bin/tcpclient 127.0.0.1:1234
test2:
	./bin/greeter_client gopher
