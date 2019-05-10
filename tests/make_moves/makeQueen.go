package main

import (
	"fmt"
	"hawat/board"
    "strings"
)

func test() {
    fmt.Printf("TESTING TALKCHESS PERFT\n")
    fmt.Printf("############################################################\n\n")
	var b board.Board

    var tests = [][2]string{}

    tests = append(tests, [2]string{"FREE MOVEMENT (27 Moves)", "rnbq1k1r/pp1Pbppp/2p5/8/2B5/8/PPP1NnPP/RNBQK2R w KQ - 1 8"})

    for _, test := range tests{
        b.SetFEN(strings.Fields(test[1]))
        fmt.Printf("%s\n------------------------------\n", test[0])
        b.Perft(1)
    }
}

func main() {
    test()
}
