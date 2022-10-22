package sample

import "fmt"

type User struct {
	name string
	*ChatRoom
}

func NewUser(name string) *User {
	return &User{name: name}
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) SendMessage(toUser *User, msg string) {
	user.ChatRoom.ForwardedMessage(user, toUser, msg)
}

func (user *User) RecvMessage(from *User, msg string) {
	fmt.Printf("[Recv Msg by from %s]: %s\n", from.GetName(), msg)
}
