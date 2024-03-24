package main

import (
	"fmt"

	"github.com/TechnicallyMay/frankly/game"
	"github.com/TechnicallyMay/frankly/terminal"
)

func main() {
    gameState := game.GameState{PlayerColor: game.White, Board: game.GetDefaultBoard()}
    for true {
        terminal.DrawGame(gameState)
        var input string
        fmt.Scan(&input)

        runes := []rune(input)
        fmt.Println(runes[0] + runes[1] + runes[2] + runes[3])
        gameState.Board.MovePiece(runes[0], runes[1], runes[2], runes[3])
    }
} 

