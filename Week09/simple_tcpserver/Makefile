build:
	cd server && go build -o ../bin/tcpserver && cd ../ \
		&& cd client && go build -o ../bin/tcpclient && cd ../

run:
	./bin/tcpserver 1234

test:
	./bin/tcpclient 127.0.0.1:1234
