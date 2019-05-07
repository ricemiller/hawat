package main

import (
	"fmt"
	"hawat/board"
)


func test() {
	var b board.Board

    fmt.Printf("TESTING ALL INITIAL MOVES\n")
    fmt.Printf("############################################################\n\n")
    b.Init()
    fmt.Printf("Initial status -------------------------------\n")
    b.Print()

    moves := b.Moves()

    for _, m := range moves {
        fmt.Printf("\n\n\n\nAfter 1.%s -------------------------------\n", board.Parse0x88_Alg(m.ToSquare))
        _, s := b.DoMove(m)
        b.Print()
        b.UndoMove(m, s)
    }


}

func main() {
    test()
}
