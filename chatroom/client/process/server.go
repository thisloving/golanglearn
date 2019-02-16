package process

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func ShowMenu() {
	fmt.Println("---------------Congratulation---------------")
	fmt.Println("---------------1.show user   ---------------")
	fmt.Println("---------------2.send msg    ---------------")
	fmt.Println("---------------3.show msg    ---------------")
	fmt.Println("---------------4.logout      ---------------")
	fmt.Println("---------------select(1-4)   ---------------")

	var key int
	var content string
	smsProcess := &SmsProcess{}

	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		outputOnlineUser()
		fmt.Println("show user:")
	case 2:
		fmt.Println("send msg:")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("show msg:")
	case 4:
		fmt.Println("logout!")
		os.Exit(0)
	default:
		fmt.Println("wrong input!")
	}
}

func serverProcessMes(Conn net.Conn) {

	tf := &utils.Transfer{Conn: Conn}
	for {
		fmt.Println("client wait to read msg.")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}

		switch mes.Type {
		case message.NotifyUserStatusMesType:
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType:
			outputGroupMes(&mes)
		default:
			fmt.Println("server return unknown msg type")
		}

		fmt.Printf("mes=%v\n", mes)
	}
}
