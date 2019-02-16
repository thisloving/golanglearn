package process

import (
	"chatroom/client/model"
	"chatroom/common/message"
	"fmt"
)

var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser

func outputOnlineUser() {
	fmt.Println("current user list:")

	for id, _ := range onlineUsers {
		fmt.Println("id：\t", id)
	}
}

func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{UserId: notifyUserStatusMes.UserId, UserStatus: notifyUserStatusMes.UserStatus}
	} else {
		user.UserStatus = notifyUserStatusMes.UserStatus
	}

	onlineUsers[notifyUserStatusMes.UserId] = user

	outputOnlineUser()
}
