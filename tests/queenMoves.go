package main

import (
	"fmt"
	"hawat/board"
    "strings"
)


func printMoves(moves []board.Move) {
    for i, move := range moves {
        fmt.Printf("Move %d: ", i+1)
        fmt.Printf("%s -> %s : %d\n", board.Parse0x88_Alg(move.FromSquare), board.Parse0x88_Alg(move.ToSquare), move.Promotion)
    }
    fmt.Printf("\n\n")
}

func test() {
    fmt.Printf("TESTING QUEEN MOVES\n")
    fmt.Printf("############################################################\n\n")
	var b board.Board

    var tests = [][2]string{}
    var moves = []board.Move{}

    tests = append(tests, [2]string{"FREE MOVEMENT (27 Moves)", "8/8/8/3Q4/8/8/8/8 w - - 0 1"})
    tests = append(tests, [2]string{"6 CAPTURES, 9 MOVES (14 Moves)", "8/8/8/2rQr3/2rrr3/8/8/8 w - - 0 1"})
    tests = append(tests, [2]string{"3 BLOCKS, 2 CAPTURE (11 Moves)", "8/8/8/2rQr3/2KKK3/8/8/8 w - - 0 1"})

    for _, test := range tests{
        b.SetFEN(strings.Fields(test[1]))
        moves = b.Moves()
        fmt.Printf("%s\n------------------------------\n", test[0])
        printMoves(moves)
    }
}

func main() {
    test()
}
