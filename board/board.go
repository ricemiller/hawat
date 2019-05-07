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

type Status struct {
	WhiteCastleKing  bool
	WhiteCastleQueen bool
	BlackCastleKing  bool
	BlackCastleQueen bool
	EnPassant        int8 // Points at the square behind the pawn (e2-e4 -> EnPassant=e3)
	HalfMoves        int
	FullMoves        int
}

type Board struct {
	Square           [0x80]int8
	ThreatToBlack    [0x80]bool
	ThreatToWhite    [0x80]bool
	PosBK            int8
	PosWK            int8
	SidePlaying      int8
    BoardStatus      Status
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
		if b.CheckEnemy(i, colour) {
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

func (b *Board) DoMove(move Move) (bool, Status) {
    legal := true
    oldStatus := b.BoardStatus

    // Save old status for undo


    // Castling
    switch move.Promotion {
    case WHITE_CASTLE_QS:
        b.Square[0x00] = 0
        b.Square[0x03] = WR

    case WHITE_CASTLE_KS:
        b.Square[0x07] = 0
        b.Square[0x05] = WR

    case BLACK_CASTLE_QS:
        b.Square[0x70] = 0
        b.Square[0x73] = BR

    case BLACK_CASTLE_KS:
        b.Square[0x77] = 0
        b.Square[0x75] = BR
    }


    // Update king position and castling booleans
    if b.Square[move.FromSquare] == WK {
        b.PosWK = move.ToSquare
        b.BoardStatus.WhiteCastleKing = false
        b.BoardStatus.WhiteCastleQueen = false

    }
    if b.Square[move.FromSquare] == BK {
        b.PosBK = move.ToSquare
        b.BoardStatus.BlackCastleKing = false
        b.BoardStatus.BlackCastleQueen = false
    }

    if b.Square[move.FromSquare] == WR {
        if move.FromSquare == 0x00 {
            b.BoardStatus.WhiteCastleQueen = false
        } else if move.FromSquare == 0x07 {
            b.BoardStatus.WhiteCastleKing = false
        }
    }
    if b.Square[move.FromSquare] == BR {
        if move.FromSquare == 0x70 {
            b.BoardStatus.BlackCastleQueen = false
        } else if move.FromSquare == 0x77 {
            b.BoardStatus.BlackCastleKing = false
        }
    }

    // Rook captured
    if move.CapturedPiece == WR {
        if move.ToSquare == 0x00 {
            b.BoardStatus.WhiteCastleQueen = false
        } else if move.ToSquare == 0x07{
            b.BoardStatus.WhiteCastleKing = false
        }
    }
    if move.CapturedPiece == BR {
        if move.ToSquare == 0x70 {
            b.BoardStatus.BlackCastleQueen = false
        } else if move.ToSquare == 0x77{
            b.BoardStatus.BlackCastleKing = false
        }
    }



    if b.Square[move.FromSquare] == WP*b.SidePlaying { // Piece moved is a pawn
        // New en passant
        if move.FromSquare+0x20*b.SidePlaying == move.ToSquare {
            b.BoardStatus.EnPassant = move.FromSquare+0x10*b.SidePlaying
        } else {
            b.BoardStatus.EnPassant = -1
            // En passant special capture
            if b.BoardStatus.EnPassant == move.ToSquare && move.CapturedPiece == BP*b.SidePlaying { // Pawn is on EnPassant square and piece captured is a pawn
                b.Square[b.BoardStatus.EnPassant-(0x10*b.SidePlaying)] = 0 // Locate and clear the captured en passant pawn
            }
        }
    } else {
        fmt.Printf("%d\n", b.Square[move.FromSquare])
        fmt.Printf("%d\n", WP*b.SidePlaying)
        b.BoardStatus.EnPassant = -1
    }

    // Make move
    b.Square[move.ToSquare] = b.Square[move.FromSquare]
    b.Square[move.FromSquare] = 0

    // Recalculate threat tables
    b.setThreats(WHITE_SIDE)
    b.setThreats(BLACK_SIDE)

    // King in check?
    if b.SidePlaying == WHITE_SIDE && b.CheckThreat(b.PosWK, b.SidePlaying) || b.SidePlaying == BLACK_SIDE && b.CheckThreat(b.PosBK, b.SidePlaying) { // If the king is in check, the move is illegal
        legal = false // Needs to complete the move making for a correct undo
    }

    // Update side playing
    b.SidePlaying = -b.SidePlaying

    // No of moves and half moves
    return legal, oldStatus

}


func (b *Board) UndoMove(move Move, oldStatus Status) {

    // Castling : fix rooks
    switch move.Promotion {
    case WHITE_CASTLE_QS:
        b.Square[0x03] = 0
        b.Square[0x00] = WR

    case WHITE_CASTLE_KS:
        b.Square[0x05] = 0
        b.Square[0x07] = WR

    case BLACK_CASTLE_QS:
        b.Square[0x73] = 0
        b.Square[0x70] = BR

    case BLACK_CASTLE_KS:
        b.Square[0x75] = 0
        b.Square[0x77] = BR
    }

    // Update king position
    if b.Square[move.ToSquare] == WK {
        b.PosWK = move.FromSquare
    }
    if b.Square[move.ToSquare] == BK {
        b.PosBK = move.FromSquare
    }

    // En passant special capture
    if move.ToSquare == oldStatus.EnPassant && b.Square[move.ToSquare] == BP * b.SidePlaying { // Enemy pawn occupies old en passant square
        b.Square[oldStatus.EnPassant-0x10*b.SidePlaying] = WP*b.SidePlaying // Replace captured en passant pawn
    }

    // No of moves and half moves

    b.Square[move.FromSquare] = b.Square[move.ToSquare]
    b.Square[move.ToSquare] = move.CapturedPiece
    b.BoardStatus = oldStatus

    // Update side playing
    b.SidePlaying = -b.SidePlaying

    // Recalculate threat tables
    b.setThreats(WHITE_SIDE)
    b.setThreats(BLACK_SIDE)

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
	fmt.Println("White Castle King:", b.BoardStatus.WhiteCastleKing)
	fmt.Println("White Castle Queen:", b.BoardStatus.WhiteCastleQueen)
	fmt.Println("Black Castle King:", b.BoardStatus.BlackCastleKing)
	fmt.Println("Black Castle Queen:", b.BoardStatus.BlackCastleQueen)
	fmt.Println("En Passant:", b.BoardStatus.EnPassant)
	fmt.Println("Half Moves:", b.BoardStatus.HalfMoves)
	fmt.Println("Full Moves:", b.BoardStatus.FullMoves)
}
