package main

import (
	"chatroom/common/message"
	"chatroom/server/process"
	"chatroom/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		up := &process2.UserProcess{Conn: this.Conn}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		up := &process2.UserProcess{Conn: this.Conn}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		smsProcess := &process2.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("wrong msg type")
	}

	return
}

func (this *Processor) processConn() (err error) {
	for {
		tf := &utils.Transfer{Conn: this.Conn}

		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("client logout")
			}
			fmt.Println("readPkg err=", err)
			return err
		}

		fmt.Println("mes=", mes)

		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
