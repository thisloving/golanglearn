package process

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
)

func outputGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err.Error())
		return
	}

	info := fmt.Sprintf("id:%d%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
}
