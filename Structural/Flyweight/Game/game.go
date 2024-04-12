package game

import "sync"

type playerIntrinsic struct {
	height int
	color  string
}

var pi *playerIntrinsic
var once sync.Once

func newPlayerIntrinsic() *playerIntrinsic {
	once.Do(func() {
		pi = &playerIntrinsic{height: 5, color: "white"}
	})
	return pi
}

type PlayerExtrinsic struct {
	X      int
	Y      int
	Health float32
	Speed  float32
	Common *playerIntrinsic
}

func NewPlayer() *PlayerExtrinsic {
	return &PlayerExtrinsic{Common: newPlayerIntrinsic()}
}
