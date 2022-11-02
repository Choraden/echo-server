package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	serverAddr := os.Args[1] + ":" + os.Args[2]
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected!")

	reader := bufio.Reader{}
	reader.Reset(os.Stdin)
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			fmt.Println("Bye!")
			conn.Close()
			return
		} else if err != nil {
			panic(err)
		}
		line = line[:len(line)-1] // skip '\n'

		_, err = conn.Write(line)
		if err != nil {
			panic(err)
		}

		buf := make([]byte, 1024)
		_, err = conn.Read(buf)
		if err != nil {
			panic(err)
		}

		fmt.Println("received:", string(buf))
	}
}
