package main

import (
    "fmt"
    //"time"
    "os"
    "bufio"
    "strings"
    "./board"
)

var ALLOWED_COMMANDS = [...]string{"uci", "debug", "isready", "setoption", "register", "ucinewgame", "position", "go", "stop", "ponderhit", "quit", "print", "perft"}
const ENGINE_NAME = "Hawat"
const ENGINE_AUTHOR = "ricemiller"


/*

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
*/

func UCIcommand(r *bufio.Reader) []string {
    cmd, _ := r.ReadString('\n')
    cmd = strings.TrimSpace(cmd)
    return strings.Fields(cmd)
}

func IsGoodCMD (cmd string) bool {
    for _, allowedCMD := range ALLOWED_COMMANDS {
        if allowedCMD == cmd {
            return true
        }
    }
    return false
}

func ClearCMDGarbage(cmd []string) []string{
    if len(cmd) == 0 {
        return cmd
    }

    for !IsGoodCMD(cmd[0]) {
        cmd = cmd[1:]
        if len(cmd) == 0 {
            return cmd
        }
    }
    return cmd
}

func main() {
    var b board.Board
    r := bufio.NewReader(os.Stdin)

    //quit := make(chan struct{})


   // go background(quit) //start processing

    // defer close(quit)

    for {
        cmd := UCIcommand(r)
        cmd = ClearCMDGarbage(cmd)
        if len(cmd) > 0 {

            switch cmd[0]{
            case "uci":
                fmt.Printf("id name %s\n", ENGINE_NAME)
                fmt.Printf("id author %s\n", ENGINE_AUTHOR)
                // SEND OPTIONS
                fmt.Printf("uciok\n")

            case "debug":

            case "isready":
                fmt.Printf("readyok\n")

            case "setoption":
                // RECIEVE OPTIONS

            case "register":
                fmt.Printf("This is FOSS, baby\n")


            case "ucinewgame":
                b.Init()



            case "position":
                cmd = cmd[1:]
                if len(cmd) > 0 {
                    switch cmd[0] {
                    case "fen":
                        cmd = cmd[1:]
                        if len(cmd) > 5 {
                            b.SetFEN(cmd[0:6])
                            cmd = cmd[6:]
                        }

                    case "startpos":
                        b.Init()
                        cmd = cmd[1:]
                    }

                    if len(cmd) > 0 {
                        if cmd[0] == "moves" {
                            cmd = cmd[1:]
                            b.Moves(cmd)
                        }
                    }
                }

            case "go":

            case "stop":

            case "ponderhit":

            case "print":
                b.Print()

            case "perft":

            case "quit":
                os.Exit(0)
            default:
            }
        }
    }
}
