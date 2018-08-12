package board

// Getters and setters for FEN format
/******* SOURCES *******
https://mediocrechess.blogspot.com/2006/12/guide-fen-notation.html
http://www.chessengine.co.uk/2015/03/24/0x88-board/
************************/

func (b *Board) SetFEN(fen string) {
    pos := 0x70
    for _, c := range fen{
        b.Position[pos] = 0

        switch c {
        case ' ':
            return

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
