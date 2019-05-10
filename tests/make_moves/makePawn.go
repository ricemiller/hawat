package main

import (
	"fmt"
	"hawat/board"
    "strings"
)

func testWhite() {
    fmt.Printf("TESTING WHITE PAWN MOVES\n")
    fmt.Printf("############################################################\n\n")
	var b board.Board

    var whiteTests = [][2]string{}

    whiteTests = append(whiteTests, [2]string{"BLOCK (0 Moves)", "8/8/8/8/8/4p3/4P3/8 w - - 0 1"})
    whiteTests = append(whiteTests, [2]string{"INITIAL + BLOCK DOUBLE (1 Moves)", "8/8/8/8/4p3/8/4P3/8 w - - 0 1"})
    whiteTests = append(whiteTests, [2]string{"INITIAL + DOUBLE (2 Moves)", "8/8/8/8/8/8/4P3/8 w - - 0 1"})
    whiteTests = append(whiteTests, [2]string{"2 CAPTURES + INITIAL + DOUBLE (4 Moves)", "8/8/8/8/8/3p1p2/4P3/8 w - - 0 1"})
    whiteTests = append(whiteTests, [2]string{"PROMOTION (2 Moves)", "8/4P3/8/8/8/8/8/8 w - - 0 1"})
    whiteTests = append(whiteTests, [2]string{"2 CAPTURES + PROMOTIONS (6 Moves)", "3p1p2/4P3/8/8/8/8/8/8 w - - 0 1"})
    whiteTests = append(whiteTests, [2]string{"EN PASSANT (2 Moves)", "8/8/8/3pP3/8/8/8/8 w - d6 0 1"})

    for _, test := range whiteTests{
        b.SetFEN(strings.Fields(test[1]))
        fmt.Printf("%s\n------------------------------\n", test[0])
        b.Perft(1)
    }
}

func testBlack() {
    fmt.Printf("TESTING BLACK PAWN MOVES\n")
    fmt.Printf("############################################################\n\n")
	var b board.Board

    var blackTests = [][2]string{}

    blackTests = append(blackTests, [2]string{"BLOCK (0 Moves)", "8/4p3/4P3/8/8/8/8/8 b - - 0 1"})
    blackTests = append(blackTests, [2]string{"INITIAL + BLOCK DOUBLE (1 Moves)", "8/4p3/8/4P3/8/8/8/8 b - - 0 1"})
    blackTests = append(blackTests, [2]string{"INITIAL + DOUBLE (2 Moves)", "8/4p3/8/8/8/8/8/8 b - - 0 1"})
    blackTests = append(blackTests, [2]string{"2 CAPTURES + INITIAL + DOUBLE (4 Moves)", "8/4p3/3P1P2/8/8/8/8/8 b - - 0 1"})
    blackTests = append(blackTests, [2]string{"PROMOTION (2 Moves)", "8/8/8/8/8/8/4p3/8 b - - 0 1"})
    blackTests = append(blackTests, [2]string{"2 CAPTURES + PROMOTIONS (6 Moves)", "8/8/8/8/8/8/4p3/3P1P2 b - - 0 1"})
    blackTests = append(blackTests, [2]string{"EN PASSANT (2 Moves)", "8/8/8/8/2Pp4/8/8/8 b - c3 0 1"})

    for _, test := range blackTests{
        b.SetFEN(strings.Fields(test[1]))
        fmt.Printf("%s\n------------------------------\n", test[0])
        b.Perft(1)
    }
}


func main() {
    testWhite()
    testBlack()
}
