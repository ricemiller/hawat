package board

// Pieces representation from http://www.craftychess.com/hyatt/boardrep.html
/***************
Piece & 4 != 0 -> Sliding Piece
Piece & 1 != 0 -> Sliding Diagonal
Piece & 2 != 0 -> Sliding Vertical/Horizontal
****************/
const EMPTY int8 = 0

const WP int8 = 1
const WN int8 = 2
const WK int8 = 3
const WB int8 = 5
const WR int8 = 6
const WQ int8 = 7

const BP int8 = -1
const BN int8 = -2
const BK int8 = -3
const BB int8 = -5
const BR int8 = -6
const BQ int8 = -7
