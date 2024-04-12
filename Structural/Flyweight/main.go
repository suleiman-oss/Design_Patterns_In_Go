package main

import (
	"fmt"

	game "github.com/suleiman-oss/Design_Patterns_in_Go/Structural/Flyweight/Game"
)

func main() {
	player1 := game.NewPlayer()
	player2 := game.NewPlayer()
	fmt.Println(player1 == player2)
	fmt.Println(player1.Common == player2.Common)
}
