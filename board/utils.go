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
