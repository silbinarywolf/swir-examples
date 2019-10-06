package game

import (
	"math"

	"github.com/silbinarywolf/swir-examples/squaregame/internal/game/input"
)

const (
	playerMoveSpeed = 4
)

var (
	player *Player = NewPlayer()
)

type Player struct {
	X, Y float64
}

func NewPlayer() *Player {
	self := new(Player)
	self.X = 32
	self.Y = 32
	return self
}

func (self *Player) Update() {
	if input.ButtonCheck(input.Left) {
		player.X -= playerMoveSpeed
	}
	if input.ButtonCheck(input.Right) {
		player.X += playerMoveSpeed
	}
	if input.ButtonCheck(input.Up) {
		player.Y -= playerMoveSpeed
	}
	if input.ButtonCheck(input.Down) {
		player.Y += playerMoveSpeed
	}
}

// GetPlayerPos is used by tests
func GetPlayerPos() (int, int) {
	x := int(math.Floor(player.X))
	y := int(math.Floor(player.Y))
	return x, y
}
