package myidea

import "testing"

func TestPlayer_PlayerDoSkill(t *testing.T) {
	player := NewPlayer("quan", "Meteor Strike", 300)
	player.PlayerDoSkill()
	player.PlayerDoPay()
}
