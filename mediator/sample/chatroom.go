package sample

type ChatRoom struct {
	users []*User
}

func (chatRoom *ChatRoom) AddUser(user *User) {
	chatRoom.users = append(chatRoom.users, user)
	user.ChatRoom = chatRoom
}

func (chatRoom *ChatRoom) ForwardedMessage(fromUser, toUser *User, msg string) {
	for _, user := range chatRoom.users {
		if toUser == nil || toUser.GetName() == user.GetName() {
			user.RecvMessage(fromUser, msg)
		}
	}
}
