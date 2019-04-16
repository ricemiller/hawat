package board

import (
	"strconv"
)

// Getters and setters for FEN format
/******* SOURCES *******
https://mediocrechess.blogspot.com/2006/12/guide-fen-notation.html
http://www.chessengine.co.uk/2015/03/24/0x88-board/
************************/

func (b *Board) setPieces(pieces string) {
	sq := 0x70

	for _, p := range pieces {
		switch p {

		case 'p':
			b.Square[sq] = BP
		case 'r':
			b.Square[sq] = BR
		case 'n':
			b.Square[sq] = BN
		case 'b':
			b.Square[sq] = BB
		case 'q':
			b.Square[sq] = BQ
		case 'k':
			b.Square[sq] = BK
			b.PosBK = sq
		case 'P':
			b.Square[sq] = WP
		case 'R':
			b.Square[sq] = WR
		case 'N':
			b.Square[sq] = WN
		case 'B':
			b.Square[sq] = WB
		case 'Q':
			b.Square[sq] = WQ
		case 'K':
			b.Square[sq] = WK
			b.PosWK = sq

		case '/':
			sq &= 0xf0 // go to beginning of row
			sq -= 0x10 // change row
			sq -= 0x01 // counters sq++

		case '1':
		case '2':
			sq += 1
		case '3':
			sq += 2
		case '4':
			sq += 3
		case '5':
			sq += 4
		case '6':
			sq += 5
		case '7':
			sq += 6
		}
		sq++
	}
}

func (b *Board) setSide(side string) {
	if side == "w" {
		b.SidePlaying = WHITE_SIDE
	} else if side == "b" {
		b.SidePlaying = BLACK_SIDE
	}
}

func (b *Board) setCastling(castling string) {
	for _, c := range castling {
		switch c {
		case 'K':
			b.WhiteCastleKing = true
		case 'Q':
			b.WhiteCastleQueen = true
		case 'k':
			b.BlackCastleKing = true
		case 'q':
			b.BlackCastleQueen = true
		case '-':
		}
	}
}

func (b *Board) setEnPassant(enPassant string) {
	if enPassant == "-" {
		b.EnPassant = -1
	} else {
		b.EnPassant = ParseAlg_0x88(enPassant)
	}
}

func (b *Board) setHalfMoves(halfMoves string) {
	b.HalfMoves, _ = strconv.Atoi(halfMoves)
}

func (b *Board) setFullMoves(fullMoves string) {
	b.FullMoves, _ = strconv.Atoi(fullMoves)
}

func (b *Board) clearBoard() {
	for i := 0; i < 0x80; i++ {
		b.Square[i] = 0
	}
	b.WhiteCastleKing = false
	b.WhiteCastleQueen = false
	b.BlackCastleKing = false
	b.BlackCastleQueen = false
}

func (b *Board) SetFEN(fen []string) {
	b.clearBoard()

	b.setPieces(fen[0])
	b.setSide(fen[1])
	b.setCastling(fen[2])
	b.setEnPassant(fen[3])
	b.setHalfMoves(fen[4])
	b.setFullMoves(fen[5])
    b.setThreats(WHITE_SIDE)
    b.setThreats(BLACK_SIDE)

}
