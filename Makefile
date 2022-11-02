.PHONY: all

all:
	go build -o tcpclient ./client/main
	go build -o tcpserver ./server/main

clean:
	rm tcpclient tcpserver