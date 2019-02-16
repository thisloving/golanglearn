package process

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(content string) (err error) {

	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail=", err.Error())
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail=", err.Error())
		return
	}

	tf := &utils.Transfer{Conn: CurUser.Conn}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes err=", err.Error())
	}

	return
}
