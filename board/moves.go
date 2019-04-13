package board

type Move struct {
    FromSquare  int8
    ToSquare    int8
    Promotion   int8
}

func (b *Board) GenMoves(sq int8, piece int8) []Move{
    var moves = []Move{}
    switch piece {
    case WP:
        moves = b.genMovesWPawn(sq)
    case BP:
        moves = b.genMovesBPawn(sq)
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


func (b *Board) genMovesWPawn(sq int8) []Move {
   var moves = []Move{}

   if sq >= 0x60  { // Last rank
       moves = append(moves, Move{sq, sq+0x10, WN})
       moves = append(moves, Move{sq, sq+0x10, WQ}) // Only promotes to queen and knight, no advantage on bishop or rook over queen
   } else {
       moves = append(moves, Move{sq, sq+0x10, 0}) // Normal move
   }
   // Initial x2 move
   if sq >= 0x10 && sq <= 0x17 && b.CheckEmpty(sq+0x20) {
       moves = append(moves, Move{sq, sq+0x20, 0})
   }

   // Capture right
   if CheckValidSquare(sq+0x11) {
       if b.CheckEnemy(sq+0x11) || b.EnPassant == sq+0x11 {
           if sq >= 0x60  {
               moves = append(moves, Move{sq, sq+0x11, WN})
               moves = append(moves, Move{sq, sq+0x11, WQ}) // Only promotes to queen and knight, no advantage on bishop or rook over queen
           } else {
               moves = append(moves, Move{sq, sq+0x11, 0}) // Normal move
           }

       }
   }

   // Capture left
   if CheckValidSquare(sq+0x0f) {
       if b.CheckEnemy(sq+0x0f) || b.EnPassant == sq+0x0f {
           if sq >= 0x60  {
               moves = append(moves, Move{sq, sq+0x0f, WN})
               moves = append(moves, Move{sq, sq+0x0f, WQ}) // Only promotes to queen and knight, no advantage on bishop or rook over queen
           } else {
               moves = append(moves, Move{sq, sq+0x0f, 0}) // Normal move
           }

       }
   }

    // TODO: En passant

   return moves
}

func (b *Board) genMovesBPawn(sq int8) []Move{
   var moves = []Move{}
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
