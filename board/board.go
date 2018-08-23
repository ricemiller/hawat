package board

import (
    "fmt"
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

const WHITE_TURN = true
const BLACK_TURN = false

type Board struct {
    Position            [128]int8
    PosBK               int
    PosWK               int
    SidePlaying         int8
    Turn                bool
    WhiteCastleKing     bool
    WhiteCastleQueen    bool
    BlackCastleKing     bool
    BlackCastleQueen    bool
    EnPassant           int8
    HalfMoves           int
    FullMoves           int

}

func (b *Board) Init(side int8) {
    b.SetFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
    b.PosBK = 0x74
    b.PosWK = 0x04
    b.SidePlaying = side
}

func (b* Board) CheckEnemy(pos int) bool {
    return (b.Position[pos]*b.SidePlaying < 0)

}

func (b* Board) CheckEmpty(pos int) bool {
    return b.Position[pos] == EMPTY
}

func (b *Board) Print() {
    var s = []string{}
    for i := 0; i < 0x88; i++ {
        if CheckValidPosition(i) {
            s = append(s, fmt.Sprintf("%d", b.Position[i]))

            if (i ^ 0x07) % 0x10 == 0 {
                fmt.Println(s)
                s = make([]string, 0)
            }
        }
    }
    fmt.Println("Turn:", b.Turn)
    fmt.Println("White Castle King:", b.WhiteCastleKing)
    fmt.Println("White Castle Queen:", b.WhiteCastleQueen)
    fmt.Println("Black Castle King:", b.BlackCastleKing)
    fmt.Println("Black Castle Queen:", b.BlackCastleQueen)
    fmt.Println("En Passant:", b.EnPassant)
    fmt.Println("Half Moves:", b.HalfMoves)
    fmt.Println("Full Moves:", b.FullMoves)
}
