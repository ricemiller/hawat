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
    fmt.Printf("TESTING KNIGHT MOVES\n")
    fmt.Printf("############################################################\n\n")
	var b board.Board

    var tests = [][2]string{}
    var moves = []board.Move{}

    tests = append(tests, [2]string{"FREE MOVEMENT (8 Moves)", "8/8/8/8/4N3/8/8/8 w - - 0 1"})
    tests = append(tests, [2]string{"8 CAPTURES (8 Moves)", "8/8/3q1q2/2q3q1/4N3/2q3q1/3q1q2/8 w - - 0 1"})
    tests = append(tests, [2]string{"8 BLOCKS (0 Moves)", "8/8/3Q1Q2/2Q3Q1/4N3/2Q3Q1/3Q1Q2/8 w - - 0 1"})
    tests = append(tests, [2]string{"CORNER, 6 OOB (2 Moves)", "8/8/8/8/8/8/8/7N w - - 0 1"})
    tests = append(tests, [2]string{"4 OOB, 1 CAPTURE, 1 BLOCK (3 Moves)", "8/8/8/8/5R2/4p3/6N1/8 w - - 0 1"})

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
