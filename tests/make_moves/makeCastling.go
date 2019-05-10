package main

import (
	"fmt"
	"hawat/board"
    "strings"
)


func test() {
	var b board.Board

    fmt.Printf("TESTING CASTLING\n")
    fmt.Printf("############################################################\n\n")
    b.SetFEN(strings.Fields("r3k2r/8/8/8/8/8/8/R3K2R w KQkq - - 0 1"))
    fmt.Printf("Initial status (WHITE TO MOVE) -------------------------------\n")
    b.Print()

    moves := b.Moves()

    for _, m := range moves {
        fmt.Printf("\n\n\n\nAfter 1.%s -------------------------------\n", board.Parse0x88_Alg(m.ToSquare))
        _, s := b.DoMove(m)
        b.Print()
        b.UndoMove(m, s)

    }

    b.SetFEN(strings.Fields("r3k2r/8/8/8/8/8/8/R3K2R b KQkq - - 0 1"))
    fmt.Printf("Initial status (BLACK TO MOVE) -------------------------------\n")
    b.Print()

    moves = b.Moves()

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
