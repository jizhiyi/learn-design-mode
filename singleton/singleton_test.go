package singleton

import (
	"learn/singleton/sample"
	"testing"
)

func TestSingleton(t *testing.T) {
	playerMgr1 := sample.NewPersonMgr()
	playerMgr2 := sample.NewPersonMgr()
	playerMgr1.PlayerCount = 12
	t.Log(playerMgr2.PlayerCount)
}
