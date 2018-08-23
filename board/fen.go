package board

import (
    "strings"
    "strconv"
)

// Getters and setters for FEN format
/******* SOURCES *******
https://mediocrechess.blogspot.com/2006/12/guide-fen-notation.html
http://www.chessengine.co.uk/2015/03/24/0x88-board/
************************/

func (b *Board) setPieces(pieces string){
    pos := 0x70

    for _, p := range pieces{
        switch p {

        case 'p':
            b.Position[pos] = BP
        case 'r':
            b.Position[pos] = BR
        case 'n':
            b.Position[pos] = BN
        case 'b':
            b.Position[pos] = BB
        case 'q':
            b.Position[pos] = BQ
        case 'k':
            b.Position[pos] = BK
        case 'P':
            b.Position[pos] = WP
        case 'R':
            b.Position[pos] = WR
        case 'N':
            b.Position[pos] = WN
        case 'B':
            b.Position[pos] = WB
        case 'Q':
            b.Position[pos] = WQ
        case 'K':
            b.Position[pos] = WK

        case '/':
            pos &= 0xf0 // go to beginning of row
            pos -= 0x10 // change row
            pos -= 0x01 // counters pos++

        case '1':
        case '2':
            pos += 1
        case '3':
            pos += 2
        case '4':
            pos += 3
        case '5':
            pos += 4
        case '6':
            pos += 5
        case '7':
            pos += 6
        }
        pos++
    }
}

func (b* Board) setSide(side string) {
    if side == "w" {
        b.Turn = WHITE_TURN
    } else if side == "b" {
        b.Turn = BLACK_TURN
    }
}

func (b* Board) setCastling(castling string){
    for _, c := range castling{
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

func (b *Board) setEnPassant(enPassant string){
    if enPassant == "-" {
       b.EnPassant = -1
   } else {
        b.EnPassant = ParseAlg_0x88(enPassant)
   }
}

func (b *Board) setHalfMoves(halfMoves string){
    b.HalfMoves, _ = strconv.Atoi(halfMoves)
}

func (b *Board) setFullMoves(fullMoves string){
    b.FullMoves, _ = strconv.Atoi(fullMoves)
}

func (b *Board) SetFEN(Fen string) {
    // CREATE RESET BOARD FUNCTION
    fen := strings.Fields(Fen)

    b.setPieces(fen[0])
    b.setSide(fen[1])
    b.setCastling(fen[2])
    b.setEnPassant(fen[3])
    b.setHalfMoves(fen[4])
    b.setFullMoves(fen[5])

}
