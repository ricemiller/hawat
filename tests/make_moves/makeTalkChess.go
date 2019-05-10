package main

/*****************************
https://www.chessprogramming.org/Perft_Results#cite_note-6
DEPTH   NODES
1       44
2       1486
3       62369
4       2103487
5       89941194



******************************/
import (
	"fmt"
	"hawat/board"
    "strings"
)

func test() {
    fmt.Printf("TESTING TALKCHESS PERFT\n")
    fmt.Printf("############################################################\n\n")
	var b board.Board
    depth := 2

    var tests = [][2]string{}

    tests = append(tests, [2]string{"TALKCHESS DEPTH 1 (44 Moves)", "rnbq1k1r/pp1Pbppp/2p5/8/2B5/8/PPP1NnPP/RNBQK2R w KQ - 1 8"})

    for _, test := range tests{
        b.SetFEN(strings.Fields(test[1]))
        fmt.Printf("%s\n------------------------------\n", test[0])
        nodes := b.Perft(depth)
        fmt.Printf("\nNODES: %d\n", nodes)
    }
}

func main() {
    test()
}
