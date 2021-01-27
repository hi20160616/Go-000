build:
	cd cmd/tcpserver && go build -o ../../bin/tcpserver && cd ../../ \
		&& cd cmd/tcpclient && go build -o ../../bin/tcpclient && cd ../../

run:
	./bin/tcpserver

test:
	./bin/tcpclient
