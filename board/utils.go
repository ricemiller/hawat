package board

import (
    "strconv"
)

func CheckValidPosition(pos int) bool {
    return ((pos & 0x88) == 0)
}


func ParseAlg_0x88(square string) int8{
    col := rune(square[0])
    column := int(col) - int('a')
    row, _ := strconv.Atoi(string(square[1]))
    row = (row-1)*0x10

    return int8(row + column)
}

func Parse0x88_Alg(square int8) string{
    if !CheckValidPosition(int(square)) {
        return ""
    }

    c := square % 0x10 + int8('a')
    column := rune(c)
    r := square / 0x10 + 1 + int8('0')
    row := rune(r)


    return string([]rune{column, row})
}

