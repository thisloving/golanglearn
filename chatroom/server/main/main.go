package main

import (
	"chatroom/server/model"
	"fmt"
	"net"
	"time"
)

func init() {
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
}

func process(conn net.Conn) {
	defer conn.Close()

	processor := &Processor{Conn: conn}
	err := processor.processConn()
	if err != nil {
		fmt.Println("client err connect with server. err=", err)
	}
}

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	fmt.Println("server listen 7889...")

	listen, err := net.Listen("tcp", "0.0.0.0:7889")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	defer listen.Close()

	for {
		fmt.Println("wait for client...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}

		go process(conn)
	}
}
