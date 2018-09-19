package board

import (
    "fmt"
    "strings"
)

// 0x88 board representation
/******** SOURCES *************
https://mediocrechess.blogspot.com/2006/12/0x88-representation.html
http://sweng.epfl.ch/archive/2009/project/buddy-suite/milestone-3/0x88-chessboard
http://www.craftychess.com/hyatt/boardrep.html
http://www.chessengine.co.uk/2015/03/24/0x88-board/
*/

const WHITE_SIDE = 1
const BLACK_SIDE = -1

type Board struct {
    Square            [128]int8
    PosBK               int
    PosWK               int
    SidePlaying         int8
    WhiteCastleKing     bool
    WhiteCastleQueen    bool
    BlackCastleKing     bool
    BlackCastleQueen    bool
    EnPassant           int8
    HalfMoves           int
    FullMoves           int

}

func (b *Board) Init() {
    b.SetFEN(strings.Fields("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"))
    b.SidePlaying = WHITE_SIDE
}

func (b* Board) CheckEnemy(sq int) bool {
    return (b.Square[sq]*b.SidePlaying < 0)

}

func (b* Board) CheckEmpty(sq int) bool {
    return b.Square[sq] == EMPTY
}

func (b* Board) Moves(moves []string){
    //TO DO
}


func (b *Board) Print() {
    var s = []string{}
    var board = [][]string{}

    for i := 0; i < 0x80; i++ {
        if CheckValidSquare(i) {
            s = append(s, fmt.Sprintf("%d", b.Square[i]))

            if (i ^ 0x07) % 0x10 == 0 {
                board = append(board, s)
                s = make([]string, 0)
            }
        }
    }

    for i := len(board)-1; i >= 0 ; i--{
        fmt.Println(board[i])
    }

    fmt.Println("Turn:", b.SidePlaying)
    fmt.Println("White Castle King:", b.WhiteCastleKing)
    fmt.Println("White Castle Queen:", b.WhiteCastleQueen)
    fmt.Println("Black Castle King:", b.BlackCastleKing)
    fmt.Println("Black Castle Queen:", b.BlackCastleQueen)
    fmt.Println("En Passant:", b.EnPassant)
    fmt.Println("Half Moves:", b.HalfMoves)
    fmt.Println("Full Moves:", b.FullMoves)
}
