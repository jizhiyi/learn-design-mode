package command

import (
	"fmt"
	"learn/command/sample"
	"testing"
)

func Test_Command(t *testing.T) {
	commands := sample.ActionCommand{}
	actor := sample.NewActor(&commands)
	commands.AddStep(sample.NewMove(sample.MoveType_Left))
	commands.AddStep(sample.NewMove(sample.MoveType_Right))
	commands.AddStep(sample.NewMove(sample.MoveType_Right))
	commands.AddStep(sample.NewMove(sample.MoveType_Stay))
	commands.AddStep(sample.NewMove(sample.MoveType_Right))

	actor.DoAction()
	fmt.Println("当前位置:", actor.GetCurPos())

	commands.Undo()
	actor.DoAction()
	fmt.Println("当前位置:", actor.GetCurPos())
}
