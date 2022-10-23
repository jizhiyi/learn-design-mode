package sample

import "fmt"

type state interface {
	doThing()
}

var workState *WorkState = &WorkState{}

type WorkState struct {
}

func (*WorkState) doThing() {
	fmt.Println("拼命工作")
}

var restState *RestState = &RestState{}

type RestState struct {
}

func (*RestState) doThing() {
	fmt.Println("不一定能好好地休息")
}

func doClock(w *Worker, hour int) {
	if hour < 9 || hour > 21 {
		w.setState(restState)
		return
	}
	w.setState(workState)
}
