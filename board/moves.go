package board

type Move struct {
    FromSquare      int8
    ToSquare        int8
    Promotion       int8 // Also denotes castling
    CapturedPiece   int8
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

func (b *Board) GenThreats(sq int8, piece int8) []Move{
    var moves = []Move{}
    var colour int8

    if piece > 0 {
       colour = WHITE_SIDE
    } else {
       colour = BLACK_SIDE
    }

    switch piece {
    case WP, BP:
        moves = b.genThreatsPawn(sq, colour)
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


func (b *Board) genThreatsPawn(sq int8, colour int8) []Move { //Only piece that threats are different to moves, necessary for King moves
    var threats = []Move{}

    captureLeft  := sq + 0x0f * colour
    captureRight := sq + 0x11 * colour

    if CheckValidSquare(captureLeft) {
        threats = append (threats, Move{sq, captureLeft, 0, b.Square[captureLeft]})
    }
    if CheckValidSquare(captureRight) {
        threats = append (threats, Move{sq, captureRight, 0, b.Square[captureRight]})
    }
    return threats
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
           moves = append(moves, Move{sq, push, WN*colour, 0})
           moves = append(moves, Move{sq, push, WQ*colour, 0}) // Only promotes to queen and knight, no advantage on bishop or rook over queen
       }
       // Captures + Promotion
       if b.CheckEnemy(captureLeft, colour) {
           moves = append(moves, Move{sq, captureLeft, WN*colour, b.Square[captureLeft]})
           moves = append(moves, Move{sq, captureLeft, WQ*colour, b.Square[captureLeft]})
       }
       if b.CheckEnemy(captureRight, colour) {
           moves = append(moves, Move{sq, captureRight, WN*colour, b.Square[captureRight]})
           moves = append(moves, Move{sq, captureRight, WQ*colour, b.Square[captureRight]})
       }
   } else {
       if b.CheckEmpty(push) {
           moves = append(moves, Move{sq, push, 0, 0}) // Push
           if ( ( sq <= 0x17  && colour == WHITE_SIDE ) || ( sq >= 0x60  && colour == BLACK_SIDE ) ) && b.CheckEmpty(doublePush) { // Double push
               moves = append(moves, Move{sq, doublePush, 0, 0})
           }
       }
       //Captures
       if b.CheckEnemy(captureLeft, colour) {
           moves = append(moves, Move{sq, captureLeft, 0, b.Square[captureLeft]})
       }
       if b.EnPassant == captureLeft {
           moves = append(moves, Move{sq, captureLeft, 0, BP*colour}) // captures enemy pawn
       }
       if b.CheckEnemy(captureRight, colour) {
           moves = append(moves, Move{sq, captureRight, 0, b.Square[captureRight]})
       }
       if b.EnPassant == captureRight {
           moves = append(moves, Move{sq, captureRight, 0, BP*colour})
       }
   }
   return moves
}

func (b *Board) genMovesKnight(sq int8, colour int8) []Move {
   var moves = []Move{}
    var targets = []int8{0x0e, 0x1f, 0x21, 0x12, -0x0e, -0x1f, -0x21, -0x12}

   for _, target := range targets {
       target += sq
       if CheckValidSquare(target) && !b.CheckAlly(target, colour) {
           moves = append(moves, Move{sq, target, 0, b.Square[target]})
       }
   }
   return moves
}

func (b *Board) genMovesSliding(sq int8, colour int8, vectors []int8) []Move {
   var moves = []Move{}

   for _, vector := range vectors {
       target := sq + vector

        for CheckValidSquare(target) && !b.CheckAlly(target, colour) {
           moves = append(moves, Move{sq, target, 0, b.Square[target]})
           if b.CheckEnemy(target, colour) { //slide until capture
               break
           }
           target += vector
        }
    }
   return moves
}


func (b *Board) genMovesBishop(sq int8, colour int8) []Move {
    var vectors = []int8{0x0f, 0x11, -0x0f, -0x11}
    return b.genMovesSliding(sq, colour, vectors)
}

func (b *Board) genMovesRook(sq int8, colour int8) []Move {
    var vectors = []int8{0x10, 0x01, -0x10, -0x01}
    return b.genMovesSliding(sq, colour, vectors)
}

func (b *Board) genMovesQueen(sq int8, colour int8) []Move {
    return append (b.genMovesBishop(sq, colour), b.genMovesRook(sq, colour)...)
}

func (b *Board) genMovesKing(sq int8, colour int8) []Move{
   var moves = []Move{}
   var targets = []int8{0x10, 0x01, -0x10, -0x01, 0x0f, 0x11, -0x0f, -0x11}
   var squaresCastlingWQ = []int8{0x01, 0x02, 0x03}
   var squaresCastlingWK = []int8{0x05, 0x06}
   var squaresCastlingBQ = []int8{0x71, 0x72, 0x73}
   var squaresCastlingBK = []int8{0x75, 0x76}
   var castle bool

   for _, target := range targets {
       target += sq
       if CheckValidSquare(target) && !b.CheckAlly(target, colour) && !b.CheckThreat(target, colour) {
           moves = append(moves, Move{sq, target, 0, b.Square[target]})
       }
   }

   // Castling
   if !b.CheckThreat(sq, colour) { // King cannot castle if in check
       if colour == WHITE_SIDE {
           if b.WhiteCastleKing {
               castle = true
               for _, square := range squaresCastlingWK {
                   if b.CheckThreat(square, colour) || !b.CheckEmpty(square) {
                       castle = false
                       break
                   }
               }
               if castle {
                   moves = append(moves, Move{sq, sq-2, WHITE_CASTLE_KS, 0}) //Promotion also describes castling
               }
           }
           if b.WhiteCastleQueen {
               castle = true
               for _, square := range squaresCastlingWQ {
                   if b.CheckThreat(square, colour) || !b.CheckEmpty(square) {
                       castle = false
                       break
                   }
               }
               if castle {
                   moves = append(moves, Move{sq, sq+2, WHITE_CASTLE_QS, 0})
               }
           }
       } else {
           if b.BlackCastleKing {
               castle = true
               for _, square := range squaresCastlingBK {
                   if b.CheckThreat(square, colour) || !b.CheckEmpty(square) {
                       castle = false
                       break
                   }
               }
               if castle {
                   moves = append(moves, Move{sq, sq-2, BLACK_CASTLE_KS, 0})
               }
           }
           if b.BlackCastleQueen {
               castle = true
               for _, square := range squaresCastlingBQ {
                   if b.CheckThreat(square, colour) || !b.CheckEmpty(square) {
                       castle = false
                       break
                   }
               }
               if castle {
                   moves = append(moves, Move{sq, sq+2, BLACK_CASTLE_QS, 0})
               }
           }

       }
   }




   return moves
}
