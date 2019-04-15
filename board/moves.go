package board

type Move struct {
    FromSquare  int8
    ToSquare    int8
    Promotion   int8
}

func (b *Board) GenMoves(sq int8, piece int8) []Move{
    var moves = []Move{}
    var colour int8

    if piece > 0 {
       colour = WHITE_SIDE
    } else {
       colour = BLACK_SIDE
    }

    switch piece {
    case WP, BP:
        moves = b.genMovesPawn(sq, colour)
    case BN, WN:
        moves = b.genMovesKnight(sq, colour)
    case BB, WB:
        moves = b.genMovesBishop(sq, colour)
    case BR, WR:
        moves = b.genMovesRook(sq, colour)
    case BQ, WQ:
        moves = b.genMovesQueen(sq, colour)
    case BK, WK:
        moves = b.genMovesKing(sq, colour)
    }
    return moves
}


func (b *Board) genMovesPawn(sq int8, colour int8) []Move {
   var moves = []Move{}

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
   }
   return moves
}

func (b *Board) genMovesKnight(sq int8, colour int8) []Move{
   var moves = []Move{}
   var targets = []int8{0x0e, 0x1f, 0x21, 0x12, -0x0e, -0x1f, -0x21, -0x12}

   for _, target := range targets {
       target += sq
       if CheckValidSquare(target) && !b.CheckAlly(target, colour) {
           moves = append(moves, Move{sq, target, 0})
       }
   }

   return moves
}

func (b *Board) genMovesBishop(sq int8, colour int8) []Move{
   var moves = []Move{}
   var vectors = []int8{0x0f, 0x11, -0x0f, -0x11}

   for _, vector := range vectors {
       target := sq + vector

        for CheckValidSquare(target) && !b.CheckAlly(target, colour) {
           moves = append(moves, Move{sq, target, 0})
           if b.CheckEnemy(target, colour) { //slide until capture
               break
           }
           target += vector
        }
    }
   return moves
}

func (b *Board) genMovesRook(sq int8, colour int8) []Move{
   var moves = []Move{}

   return moves
}
func (b *Board) genMovesQueen(sq int8, colour int8) []Move{
   var moves = []Move{}

   return moves
}
func (b *Board) genMovesKing(sq int8, colour int8) []Move{
   var moves = []Move{}

   return moves
}
