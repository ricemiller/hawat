package main

import (
	"fmt"
	"hawat/board"
)


func test() {
	var b board.Board
        fmt.Printf("PERFT TEST\n")
        fmt.Printf("############################################################\n\n")

    for i:=0; i<10; i++ {
        b.Init()
        nodes := b.Perft(i)
        fmt.Printf("DEPTH: %d, NODES: %d\n", i, nodes)
    }

}

func main() {
    test()
}
