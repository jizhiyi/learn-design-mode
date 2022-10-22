package sample

import "testing"

func Test_ChatRoom(t *testing.T) {
	chatRoom := &ChatRoom{}
	user1 := NewUser("jizhihong")
	user2 := NewUser("jizhiyi")
	user3 := NewUser("zhugufang")
	chatRoom.AddUser(user1)
	chatRoom.AddUser(user2)
	chatRoom.AddUser(user3)

	user1.SendMessage(user2, "hello")
	user3.SendMessage(nil, "Hello")
}
