package terminal

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/TechnicallyMay/frankly/array"
	"github.com/TechnicallyMay/frankly/game"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
    clear = make(map[string]func())
    clear["linux"] = func() { 
        cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls")
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func DrawGame(state game.GameState) {
    displayFiles := array.GetCopyOfArray(state.Board.FileNames)
    displayRanks := array.GetCopyOfArray(state.Board.RankNames)

    if state.PlayerColor == game.White {
        // print 1st rank last (at bottom) for white's perspective
        array.ReverseArray(displayRanks)
    } else {
        // print h-file first (on left) for black's perspective
        array.ReverseArray(displayFiles)
    }

    clearTerminal()

    for _, rank := range displayRanks {
        for _, file := range displayFiles {

            boardElement := state.Board.GetElementAtCoordinates(file, rank)
            if state.Board.SelectedFile == file && state.Board.SelectedRank == rank {
                fmt.Printf(" >%v< ", string(boardElement.GetSymbol()))
            } else {
                fmt.Printf("  %v  ", string(boardElement.GetSymbol()))
                // fmt.Printf("  %v%v:%v  ", file, rank, string(boardelement.getsymbol()))
            }
        }

        fmt.Println()
        fmt.Println()
    }

    // prompt
    fmt.Print("> ")
}

func clearTerminal() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
}

