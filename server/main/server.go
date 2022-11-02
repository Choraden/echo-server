package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

type Writer struct {
	client string
}

func (w Writer) Write(p []byte) (n int, err error) {
	return fmt.Printf("received from client %s: %s\n", w.client, string(p))
}

func main() {
	addr := os.Args[1] + ":" + os.Args[2]
	server, err := net.Listen("tcp", addr)

	if err != nil {
		panic(err)
	}
	defer server.Close()

	fmt.Println("Server is running on:", server.Addr().Network(), server.Addr().String())

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Failed to accept conn:", err)
			continue
		}

		fmt.Println("Connected user", conn.RemoteAddr().String())

		go func(conn net.Conn) {
			defer func() {
				conn.Close()
				fmt.Println("User", conn.RemoteAddr().String(), "disconnected")
			}()
			reader := io.TeeReader(conn, Writer{client: conn.RemoteAddr().String()})
			io.Copy(conn, reader)
		}(conn)
	}
}
