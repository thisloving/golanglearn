package main

import (
	"chatroom/client/process"
	"fmt"
	"os"
)

var userId int
var userPwd string
var userName string

func main() {
	var key int

	for {
		fmt.Println("---------------Welcome------------------")
		fmt.Println("\t\tlog in")
		fmt.Println("\t\tregister")
		fmt.Println("\t\tlog out")
		fmt.Println("\t\tselect(1-3)")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("log in")
			fmt.Println("id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("pwd:")
			fmt.Scanf("%s\n", &userPwd)

			up := &process.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("register")
			fmt.Println("input id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("input pwd:")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("input username:")
			fmt.Scanf("%s\n", &userName)

			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("log out!")
			os.Exit(0)
		default:
			fmt.Println("error")
		}
	}
}
