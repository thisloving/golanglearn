package process

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type UserProcess struct {
}

func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	conn, err := net.Dial("tcp", "localhost:7889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()

	var mes message.Message
	mes.Type = message.RegisterMesType

	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json marshal err=", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	tf := &utils.Transfer{Conn: conn}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("register send msg err=", err)
	}

	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("register success")
	} else {
		fmt.Println(registerResMes.Error)
	}

	os.Exit(0)
	return
}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	fmt.Printf(" userId = %d, userPwd = %s\n", userId, userPwd)

	conn, err := net.Dial("tcp", "localhost:7889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()

	var mes message.Message
	mes.Type = message.LoginMesType

	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json marshal err=", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn write(len) err=", err)
		return
	}

	fmt.Printf("client send success=%d data=%s\n", len(data), string(data))

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn write(data) err=", err)
		return
	}

	tf := &utils.Transfer{Conn: conn}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline

		fmt.Println("usersids:")
		for _, id := range loginResMes.UsersIds {
			if id == userId {
				continue
			}

			fmt.Println("user id:", id)

			user := &message.User{
				UserId:     id,
				UserStatus: message.UserOnline,
			}
			onlineUsers[id] = user
		}

		fmt.Printf("\n")

		go serverProcessMes(conn)

		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}

	return
}
