package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}

	defer conn.Close()

	go func() {
		buf := make([]byte, 2048)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("conn.Read err = ", err)
				return
			}
			addr := conn.RemoteAddr().String()
			fmt.Printf("from %s: response :%s", addr, string(buf[:n]))
		}
	}()

	str := make([]byte, 2048)
	for {
		n, err := os.Stdin.Read(str)
		if err != nil {
			fmt.Println("os.Stdin. err = ", err)
			return
		}

		if "exit" == string(str[:n-1]) {
			return
		}

		conn.Write(str[:n])
	}
}
