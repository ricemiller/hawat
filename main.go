package main

import (
    "fmt"
    "time"
    "os"
    "bufio"
    "strings"
)

var x int
const ENGINE_NAME = "Hawat"
const ENGINE_AUTHOR = "ricemiller"

func background(quit chan struct{}) {
    for {
        select{
        default:
            x = x+1
            time.Sleep(1000 * time.Millisecond)
        case <-quit:
            return
        }
    }
}

func readUntilCMD(r *bufio.Reader, s string) {
    for {
        cmd,_ := r.ReadString('\n')
        if strings.TrimSpace(cmd) == s {
            return
        }
    }

}

func UCIinit(r *bufio.Reader) {
    readUntilCMD(r, "uci")
    fmt.Println("id name ", ENGINE_NAME)
    fmt.Println("id author ", ENGINE_AUTHOR)
    // Send Options
    fmt.Println("uciok")

    //Receive Options
    //TO DO: Could send quit at this point!

    readUntilCMD(r, "isready")
    fmt.Println("readyok")

    readUntilCMD(r, "ucinewgame")
}







func main() {
    x = 0
    r := bufio.NewReader(os.Stdin)
    quit := make(chan struct{})

    UCIinit(r)
    //Receive position startpos moves ...
    //Receive go infinite

    go background(quit) //start processing

    defer close(quit)
    for {
        cmd,_ := r.ReadString('\n')
        switch strings.TrimSpace(cmd) {
            default:
                fmt.Println("Error, command not found. Only \"sub\", \"print\" and \"quit\"")
            case "sub":
                x = x-10

            case "print":
                fmt.Println(x)

            case "quit":
                os.Exit(0)
        }
    }
}

