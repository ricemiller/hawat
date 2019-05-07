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


type Board struct {
	Square           [0x80]int8
	ThreatToBlack    [0x80]bool
	ThreatToWhite    [0x80]bool
	PosBK            int
	PosWK            int
	SidePlaying      int8
	WhiteCastleKing  bool
	WhiteCastleQueen bool
	BlackCastleKing  bool
	BlackCastleQueen bool
	EnPassant        int8 // Points at the square behind the pawn (e2-e4 -> EnPassant=e3)
	HalfMoves        int
	FullMoves        int
}

func (b *Board) Init() {
	b.SetFEN(strings.Fields("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"))
}

func (b *Board) CheckThreat(sq int8, colourPlaying int8) bool {
    if colourPlaying == WHITE_SIDE {
        return b.ThreatToWhite[sq]
    } else {
        return b.ThreatToBlack[sq]
    }
}

func (b *Board) CheckAlly(sq int8, colourPlaying int8) bool {
	return (b.Square[sq]*colourPlaying > 0)
}

func (b *Board) CheckEnemy(sq int8, colourPlaying int8) bool {
	return (b.Square[sq]*colourPlaying < 0)
}

func (b *Board) CheckEmpty(sq int8) bool {
	return (b.Square[sq] == EMPTY)
}

func (b *Board) Moves() []Move { //ADD DEPTH LATER ON
    var moves = []Move{}
	for i := int8(0); i < 0x78; i++ {
		if b.CheckAlly(i, b.SidePlaying) {
            moves = append(moves, b.GenMoves(i, b.Square[i])...)
			}
		}
    return moves
}

func (b *Board) setThreats(colour int8) {
    var threats = []Move{}
	for i := int8(0); i < 0x78; i++ {

        // Get all enemmy attack moves
		if b.CheckEnemy(i, b.SidePlaying) {
            threats = append(threats, b.GenThreats(i, b.Square[i])...)
			}
		}

        // Update threats table
        if colour == WHITE_SIDE {
            for i := 0; i < 0x80; i++ {
                b.ThreatToWhite[i] = false
            }

            for _, threat := range threats {
                b.ThreatToWhite[threat.ToSquare] = true
            }
        } else {
            for i := 0; i < 0x80; i++ {
                b.ThreatToBlack[i] = false
            }
            for _, threat := range threats {
                b.ThreatToBlack[threat.ToSquare] = true
            }
        }
}

func (b *Board) Perft(moves []string) {
	//TO DO
}

func (b *Board) DoMove(move Move){
    b.Square[move.ToSquare] = b.Square[move.FromSquare]
    b.Square[move.FromSquare] = 0
}


func (b *Board) UndoMove(move Move) {
    b.Square[move.FromSquare] = b.Square[move.ToSquare]
    b.Square[move.ToSquare] = move.CapturedPiece
}


func (b *Board) Print() {
	var s = []string{}
	var board = [][]string{}

	for i := int8(0); i < 0x78; i++ {
		if CheckValidSquare(i) {
			s = append(s, fmt.Sprintf("%s", PieceToChar(b.Square[i])))

			if (i^0x07)%0x10 == 0 {
                board = append(board, s)
				s = make([]string, 0)
			}
		}
	}

	for i := len(board) - 1; i >= 0; i-- {
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
