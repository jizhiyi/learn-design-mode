package sample

import (
	"learn/memento/sample/game"
	"testing"
	"time"
)

func Test_Game(t *testing.T) {
	gamer := game.NewGamer(100)
	memento := gamer.CreateMemento()

	for i := 0; i < 100; i++ {
		t.Log("=============", i)
		t.Log("当前状态:", gamer)

		gamer.Bet()

		t.Log("所持金钱为", gamer.GetMoney(), "元.")

		if gamer.GetMoney() > memento.GetMoney() {
			t.Log("    （所持金钱增加了许多，因此保存游戏当前的状态）")
			memento = gamer.CreateMemento()
		} else if gamer.GetMoney() < memento.GetMoney()/2 {
			t.Log("    （所持金钱减少了许多，因此将游戏恢复至以前的状态）")
			gamer.RestoreMemento(memento)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
