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

type Board struct {
    Position    [128]int8
    PosBK       int
    PosWK       int
    SidePlaying int8

}

func (b *Board) Init(side int8) {
    b.SetFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
    b.PosBK = 0x74
    b.PosWK = 0x04
    b.SidePlaying = side
}

func CheckValidPosition(pos int) bool {
    return ((pos & 0x88) == 0)
}

func (b* Board) CheckEnemy(pos int) bool {
    return (b.Position[pos]*b.SidePlaying < 0)

}

func (b* Board) CheckEmpty(pos int) bool {
    return b.Position[pos] == EMPTY
}

func (b *Board) Print() { //fix print
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
}

func (b *Board) PosibleP(x, y int) []string {
    var s = []string{}

    return s
}
func (b *Board) PosibleR(x, y int) []string {
    var s = []string{}

    return s
}
func (b *Board) PosibleN(x, y int) []string {
    var s = []string{}

    return s
}
func (b *Board) PosibleB(x, y int) []string {
    var s = []string{}

    return s
}
func (b *Board) PosibleQ(x, y int) []string {
    var s = []string{}

    return s
}
func (b *Board) PosibleK(x, y int) []string {
    var s = []string{}
    for i := x-1; i <= x+1; i++{
        for j := y-1; j <= y+1; j++{
            if i != 0 && j != 0 {

                //if b.CheckEnemy(i,j) || b.CheckEmpty(i,j) {
                //}

            }
        }
    }

    return s
}

/*func (b *Board) PosibleMoves() []string {
    var s = []string{}

    for i := 0; i < 8; i++ {
        var s = []string{}
        for j :=0; j <8 ; j++{
            switch b.Position[i][j] {
            case 'P':
                s = append(s, b.PosibleP(i,j)...)

            case 'R':
                s = append(s, b.PosibleR(i,j)...)

            case 'N':
                s = append(s, b.PosibleN(i,j)...)

            case 'B':
                s = append(s, b.PosibleB(i,j)...)

            case 'Q':
                s = append(s, b.PosibleQ(i,j)...)

            case 'K':
                s = append(s, b.PosibleK(i,j)...)
            }
        }
    }

    return s
}
*/

