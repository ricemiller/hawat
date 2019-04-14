package board

type Move struct {
    FromSquare  int8
    ToSquare    int8
    Promotion   int8
}

func (b *Board) GenMoves(sq int8, piece int8) []Move{
    var moves = []Move{}
    switch piece {
    case WP, BP:
        moves = b.genMovesPawn(sq, piece)
    case BN, WN:
        moves = b.genMovesKnight(sq)
    case BB, WB:
        moves = b.genMovesBishop(sq)
    case BR, WR:
        moves = b.genMovesRook(sq)
    case BQ, WQ:
        moves = b.genMovesQueen(sq)
    case BK, WK:
        moves = b.genMovesKing(sq)
    }
    return moves
}


func (b *Board) genMovesPawn(sq int8, piece int8) []Move {
   var moves = []Move{}
   var colour int8

   if piece == WP {
       colour = 1
   } else {
       colour = -1
   }

   push         := sq + 0x10 * colour
   doublePush   := sq + 0x20 * colour
   captureLeft  := sq + 0x0f * colour
   captureRight := sq + 0x11 * colour

   if ( sq >= 0x60 && colour == WHITE_SIDE ) || ( sq <= 0x17 && colour == BLACK_SIDE ) { // Last rank
       // Push + Promotion
       if b.CheckEmpty(push) {
           moves = append(moves, Move{sq, push, WN*colour})
           moves = append(moves, Move{sq, push, WQ*colour}) // Only promotes to queen and knight, no advantage on bishop or rook over queen
       }
       // Captures + Promotion
       if b.CheckEnemy(captureLeft, colour) {
           moves = append(moves, Move{sq, captureLeft, WN*colour})
           moves = append(moves, Move{sq, captureLeft, WQ*colour})
       }
       if b.CheckEnemy(captureRight, colour) {
           moves = append(moves, Move{sq, captureRight, WN*colour})
           moves = append(moves, Move{sq, captureRight, WQ*colour})
       }
   } else {
       if b.CheckEmpty(push) {
           moves = append(moves, Move{sq, push, 0}) // Push
           if ( ( sq <= 0x17  && colour == WHITE_SIDE ) || ( sq >= 0x60  && colour == BLACK_SIDE ) ) && b.CheckEmpty(doublePush) { // Double push
               moves = append(moves, Move{sq, doublePush, 0})
           }
       }
       //Captures
       if b.CheckEnemy(captureLeft, colour) || b.EnPassant == captureLeft {
           moves = append(moves, Move{sq, captureLeft, 0})
       }
       if b.CheckEnemy(captureRight, colour) || b.EnPassant == captureRight {
           moves = append(moves, Move{sq, captureRight, 0})
       }
   // Initial x2 move
   }

   return moves
}

func (b *Board) genMovesKnight(sq int8) []Move{
   var moves = []Move{}

   return moves
}
func (b *Board) genMovesBishop(sq int8) []Move{
   var moves = []Move{}

   return moves
}
func (b *Board) genMovesRook(sq int8) []Move{
   var moves = []Move{}

   return moves
}
func (b *Board) genMovesQueen(sq int8) []Move{
   var moves = []Move{}

   return moves
}
func (b *Board) genMovesKing(sq int8) []Move{
   var moves = []Move{}

   return moves
}
