package board

import (
	"strconv"
)

// Pieces representation from http://www.craftychess.com/hyatt/boardrep.html
/***************
Piece & 4 != 0 -> Sliding Piece
Piece & 1 != 0 -> Sliding Diagonal
Piece & 2 != 0 -> Sliding Vertical/Horizontal
****************/
const EMPTY int8 = 0

const WHITE_SIDE = 1
const BLACK_SIDE = -1

const WP int8 = 1
const WN int8 = 2
const WK int8 = 3
const WB int8 = 5
const WR int8 = 6
const WQ int8 = 7
const WHITE_CASTLE_KS int8 = 8
const WHITE_CASTLE_QS int8 = 9

const BP int8 = -1
const BN int8 = -2
const BK int8 = -3
const BB int8 = -5
const BR int8 = -6
const BQ int8 = -7
const BLACK_CASTLE_KS int8 = -8
const BLACK_CASTLE_QS int8 = -9

func CheckValidSquare(sq int8) bool {
	return ((int(sq) & 0x88) == 0)
}

func ParseAlg_0x88(square string) int8 {
	col := rune(square[0])
	column := int(col) - int('a')
	row, _ := strconv.Atoi(string(square[1]))
	row = (row - 1) * 0x10

	return int8(row + column)
}

func Parse0x88_Alg(square int8) string {
	if !CheckValidSquare(square) {
		return ""
	}

	c := square%0x10 + int8('a')
	column := rune(c)
	r := square/0x10 + 1 + int8('0')
	row := rune(r)

	return string([]rune{column, row})
}

func PieceToChar(piece int8) string {
    switch piece {
    case WP:
        return "P"
    case WN:
        return "N"
    case WB:
        return "B"
    case WR:
        return "R"
    case WQ:
        return "Q"
    case WK:
        return "K"
    case BP:
        return "p"
    case BN:
        return "n"
    case BB:
        return "b"
    case BR:
        return "r"
    case BQ:
        return "q"
    case BK:
        return "k"
    }
    return "."
}

