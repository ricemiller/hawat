package main

import (
	"fmt"
	"hawat/board"
    "os"
    "strconv"
)

func test(fen []string, depth int) {
	var b board.Board
	fmt.Printf("PERFT TEST\n")
	fmt.Printf("############################################################\n\n")

    b.SetFEN(fen)
    nodes := b.Perft(depth)
    fmt.Printf("DEPTH: %d, NODES: %d\n", depth, nodes)

}

func main() {
    args := os.Args[1:]
    if len(args) != 7 {
        fmt.Printf("Error: Wrong number of parameters\n\n")
        os.Exit(1)
    }

    fmt.Println(args[:6])
    depth, _ := strconv.Atoi(args[6])
    test(args[:6], depth)
}
