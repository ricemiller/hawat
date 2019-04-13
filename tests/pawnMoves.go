package main

import (
	"fmt"
	"hawat/board"
    "strings"
)

func main() {
	var b board.Board
	//b.SetFEN(strings.Fields("8/8/8/8/8/4p3/4P3/8 w - - 0 1")) // Block -> 0 moves
	//b.SetFEN(strings.Fields("8/8/8/8/4p3/8/4P3/8 w - - 0 1")) // Initial + block on Double -> 1 move
	//b.SetFEN(strings.Fields("8/8/8/8/8/8/4P3/8 w - - 0 1")) // Initial + Double -> 2 moves
	//b.SetFEN(strings.Fields("8/8/8/8/8/3p1p2/4P3/8 w - - 0 1")) // 2 captures + initial move + double-> 4 moves
	//b.SetFEN(strings.Fields("8/4P3/8/8/8/8/8/8 w - - 0 1")) // Promotion -> 2 moves
	//b.SetFEN(strings.Fields("3p1p2/4P3/8/8/8/8/8/8 w - - 0 1")) // 2 captures + promotion -> 6 moves
	b.SetFEN(strings.Fields("8/8/8/3pP3/8/8/8/8 w - d6 0 1")) // En passant -> 2 moves
    moves := b.Moves()

    for i, move := range moves {
        fmt.Println("Move", i, ":", move)
    }

}
