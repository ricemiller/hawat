
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
    fmt.Printf("TESTING KING MOVES\n")
    fmt.Printf("############################################################\n\n")
	var b board.Board

    var tests = [][2]string{}
    var moves = []board.Move{}
    tests = append(tests, [2]string{"FREE MOVEMENT (8 Moves)", "8/8/8/3K4/8/8/8/8 w - - 0 1"})
    tests = append(tests, [2]string{"4 CAPTURES, 2 BLOCKS (6 Moves)", "8/8/4Pp2/3PKp2/4pp2/8/8/8 w - - 0 1"})
    tests = append(tests, [2]string{"5 THREATS (3 Moves)", "5rb1/8/8/4K3/8/8/8/8 w - - 0 1"})
    tests = append(tests, [2]string{"WHITE BOTH CASTLING (7 Moves)", "8/8/8/8/8/8/8/R3K2R w KQ - 0 1"})
    tests = append(tests, [2]string{"WHITE IN CHECK, NO CASTLING (4 Moves)", "8/8/8/8/4r3/8/8/R3K2R w KQ - 0 1"})
    tests = append(tests, [2]string{"WHITE THREAT BOTH CASTLINGS (1 Moves)", "8/8/8/8/3r1r2/8/8/R3K2R w KQ - 0 1"})
    tests = append(tests, [2]string{"BLACK BOTH CASTLING (7 Moves)", "r3k2r/8/8/8/8/8/8/8 b kq - 0 1"})
    tests = append(tests, [2]string{"BLACK IN CHECK, NO CASTLING (4 Moves)", "r3k2r/8/8/4R3/8/8/8/8 b kq - 0 1"})
    tests = append(tests, [2]string{"BLACK THREAT BOTH CASTLINGS (1 Moves)", "r3k2r/8/8/3R1R2/8/8/8/8 b kq - 0 1"})

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
