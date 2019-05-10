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
    fmt.Printf("TESTING ROOK MOVES\n")
    fmt.Printf("############################################################\n\n")
	var b board.Board

    var tests = [][2]string{}
    var moves = []board.Move{}

    tests = append(tests, [2]string{"FREE MOVEMENT (14 Moves)", "8/8/8/3R4/8/8/8/8 w - - 0 1"})
    tests = append(tests, [2]string{"4 CAPTURES, 2 MOVES (6 Moves)", "8/3q4/8/1q1Rq3/3q4/8/8/8 w - - 0 1"})
    tests = append(tests, [2]string{"2 BLOCKS, 1 CAPTURE (5 Moves)", "8/8/8/1q1RK3/3K4/8/8/8 w - - 0 1"})

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
